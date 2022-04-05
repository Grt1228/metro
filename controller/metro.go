package controller

import (
	"github.com/gin-gonic/gin"
	"metro/dao"
	"metro/model"
	"time"
)

func Metros(c *gin.Context) {
	metro := model.Metro{
		ChnName:      "chnName",
		EngName:      "engName",
		ShortName:    "shortName",
		FirstLetter:  "S",
		Icon:         "icon",
		OperatePic:   "operatePic",
		PlanPic:      "planPic",
		Website:      "website",
		City:         "city",
		Length:       100,
		LineNum:      10,
		StandNum:     100,
		OpenDate:     "19990101",
		BuildDate:    "19960101",
		BuildCompany: "buildCompany",
		Desc:         "desc",
		TicketDesc:   "ticketDesc",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	db := dao.GetDB()
	db.Create(&metro)
	//db.Find(&metro)
	ResponseSuccess(c, metro)
}
