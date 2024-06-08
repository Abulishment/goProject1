package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid)
	if err != nil {
		ReturnError(c, 4004, "无相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
}