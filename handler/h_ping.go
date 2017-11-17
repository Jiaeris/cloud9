package handler

import (
	"cloud9/msg"
	"github.com/gin-gonic/gin"
)

func ServerPing(c *gin.Context) {
	c.JSON(200, msg.ResponseMessage{
		Code: 200,
		Data: "pong",
	})
}
