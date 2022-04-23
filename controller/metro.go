package controller

import (
	"github.com/gin-gonic/gin"
	"metro/model"
	"metro/service"
)

func AddMetros(c *gin.Context) {
	var metro model.Metro
	if err := c.ShouldBind(&metro); err != nil {
		ResponseSuccess(c, "参数异常")
	}
	id := service.AddMetro(&metro)
	ResponseSuccess(c, id)
}
