package controller

import (
	"github.com/gin-gonic/gin"
	"metro/controller/req"
	"metro/model"
	"metro/service"
	"strconv"
)

func AddMetros(c *gin.Context) {
	var metro model.Metro

	if err := c.ShouldBind(&metro); err != nil {
		ResponseError(c, "参数异常")
		return
	}
	id := service.AddMetro(&metro)
	ResponseSuccess(c, id)
	return
}

func DelMetros(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	service.DelMetros(id)
	ResponseSuccess(c, 1)
	return
}

func ListMetros(c *gin.Context) {
	var param req.MetroListReq
	c.ShouldBind(&param)
	metros := service.ListMetros(&param)
	ResponseSuccess(c, metros)
	return
}
