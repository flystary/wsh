package model

import (
	"io"
	"log"
	"unicode/utf8"

	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var WcList []*WriteCounter

type WriteCounter struct {
	Total int
	Id    string
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += n
	return n, nil
}

type wsQutput struct {
	ws *websocket.Conn
}

func (w *wsQutput) Write(p []byte) (int, error) {
	if !utf8.Valid(p) {
		bufStr := string(p)
		buf := make([]rune, 0, len(bufStr))

		for _, r := range bufStr {
			if r == utf8.RuneError {
				buf = append(buf, []rune("@")...)
			} else {
				buf = append(buf, r)
			}
		}
		p = []byte(string(buf))
	}
	err := w.ws.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}

// SSHClient 结构体
type SSHClient struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IPAddress string `json:"ipaddress"`
	Port      int    `json:"port"`
	LoginType int    `json:"logintype"`
	Client    *ssh.Client
	Sftp      *sftp.Client
	StdinPipe io.WriteCloser
	Session   *ssh.Session
}

func NewSSHClien() SSHClient {
	client := SSHClient{}
	client.Username = "root"
	client.Port = 22
	return client
}


func(sclient *SSHClient) Close() {
	defer func ()  {
		if err := recover(); err != nil {
			log.Println("SSHClient Close recover from panic: ", err)
		}
	}()

	if sclient.StdinPipe != nil {
		sclient.StdinPipe.Close()
		sclient.StdinPipe = nil
	}
	if sclient.Session != nil {
		sclient.Session.Close()
		sclient.Session = nil
	}
	if sclient.Sftp != nil {
		sclient.Sftp.Close()
		sclient.Sftp = nil
	}
	if sclient.Client != nil {
		sclient.Client.Close()
		sclient.Client = nil
	}
}
