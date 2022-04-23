package model

import "time"

type Metro struct {
	Id           int       `json:"AUTO_INCREMENT"`
	ChnName      string    `json:"chn_name" binding:"required"`
	EngName      string    `json:"eng_name" binding:"required"`
	ShortName    string    `json:"short_name" binding:"required"`
	FirstLetter  string    `json:"first_letter" binding:"required"`
	Icon         string    `json:"icon" binding:"required"`
	OperatePic   string    `json:"operate_pic" binding:"required"`
	PlanPic      string    `json:"plan_pic" binding:"required"`
	Website      string    `json:"website"`
	City         string    `json:"city" binding:"required"`
	Length       float32   `json:"length" binding:"required"`
	LineNum      int       `json:"line_num" binding:"required"`
	StandNum     int       `json:"stand_num" binding:"required"`
	OpenDate     string    `json:"open_date" binding:"required"`
	BuildDate    string    `json:"build_date" binding:"required"`
	BuildCompany string    `json:"build_company" binding:"required"`
	Remark       string    `json:"remark"`
	TicketDesc   string    `json:"ticket_desc"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

//func (Metro) TableName() string {
//	return "metro"
//}
