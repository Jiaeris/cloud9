package handler

import (
	"github.com/gin-gonic/gin"
	"cloud9/msg"
)

func ServerPing(c *gin.Context) {
	c.JSON(200, msg.ResponseMessage{
		Code: 200,
		Data: "pong",
	})
}
