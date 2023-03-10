package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"wsh/http/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func TermWs(c *gin.Context, timeout time.Duration) *ResponseBody {
	responseBody := ResponseBody{Msg: "succes", Code: 200}
	defer TimeCost(time.Now(), &responseBody)

	sshInfo := c.DefaultQuery("sshInfo", "")
	cols := c.DefaultQuery("cols", "150")
	rows := c.DefaultQuery("rows", "35")
	closeTip := c.DefaultQuery("closeTip", "Connection timed out!")
	col, _ := strconv.Atoi(cols)
	row, _ := strconv.Atoi(rows)
	sshClient, err := model.DecodedMsgToSSHClient(sshInfo)
	if err != nil {
		log.Println(err)
		responseBody.Msg = err.Error()
		responseBody.Code = 500
		return &responseBody
	}

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		responseBody.Msg = err.Error()
		responseBody.Code = 500
		return &responseBody
	}
	err = sshClient.GenerateClient()
	if err != nil {
		wsConn.WriteMessage(1, []byte(err.Error()))
		wsConn.Close()
		log.Println(err)
		responseBody.Msg = err.Error()
		responseBody.Code = 500
		return &responseBody
	}
	sshClient.InitTerminal(wsConn, row, col)
	sshClient.Connect(wsConn, timeout, closeTip)
	return &responseBody
}