package api

import (
	"app/api/response"
	"app/util"
	"app/websocket"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Websocket(c *gin.Context) {
	userId := c.Query("user_id")
	group := c.Query("group")
	if !util.InArray(group, []string{"api", "admin", "mini", "test"}) {
		response.Error(c, "group参数不正确")
		return
	}
	if strings.ToLower(c.GetHeader("Connection")) != "upgrade" {
		response.Error(c, "header Connection not Upgrade")
		return
	}
	if strings.ToLower(c.GetHeader("Upgrade")) != "websocket" {
		response.Error(c, "header Upgrade not websocket")
		return
	}
	err := websocket.NewServer().RegisterClient(c, userId, group)
	if err != nil {
		util.NewLogger().Error("websocket", util.Map{
			"error": err.Error(),
		})
		return
	}
}

func SendMsg(c *gin.Context) {
	response.Error(c, "该接口暂未开放")
	return
	id := c.Query("id")
	group := c.Query("group")
	msgType := c.Query("type") //group user client
	if !util.InArray(group, []string{"api", "admin", "mini", "test"}) {
		response.Error(c, "group参数不正确")
		return
	}
	msg := websocket.NewClientMessage("test1", group+"_"+msgType+"_"+util.Date())
	if msgType == "group" {
		websocket.NewServer().SendGroupMessage(msg, group)
	} else if msgType == "user" {
		fmt.Println(msg, group, id)
		websocket.NewServer().SendUserMessage(msg, group, id)
	}
	response.Success(c, "success", nil)
	return
}
