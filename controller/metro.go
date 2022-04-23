package controller

import (
	"github.com/gin-gonic/gin"
	"metro/model"
	"metro/service"
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
