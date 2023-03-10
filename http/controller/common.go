package controller

import (
	"log"
	"time"
	"wsh/http/model"

	"github.com/gin-gonic/gin"
)


type ResponseBody struct {
	Duration string
	Data	 interface{}
	Msg		 string
	Code	 int
}

func TimeCost(start time.Time, body *ResponseBody) {
	body.Duration = time.Since(start).String()
}

func CheckSSH(c *gin.Context) *ResponseBody {
	responseBody :=  ResponseBody{Msg: "success", Code: 200}
	defer TimeCost(time.Now(), &responseBody)

	sshInfo := c.DefaultQuery("sshInfo", "")
	sshClient, err := model.DecodedMsgToSSHClient(sshInfo)
	if err != nil {
		log.Println(err)
		responseBody.Msg = err.Error()
		responseBody.Code = 500
		return &responseBody
	}

	err = sshClient.GenerateClient()
	defer sshClient.Close()

	if err != nil {
		log.Println(err)
		responseBody.Msg = err.Error()
		responseBody.Code = 500
	}
	return &responseBody
}