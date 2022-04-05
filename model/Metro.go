package model

import "time"

type Metro struct {
	Id           int    `grom:"AUTO_INCREMENT"`
	ChnName      string `gorm:"not null;"`
	EngName      string
	ShortName    string
	FirstLetter  string
	Icon         string
	OperatePic   string
	PlanPic      string
	Website      string
	City         string
	Length       float32
	LineNum      int
	StandNum     int
	OpenDate     string
	BuildDate    string
	BuildCompany string
	Desc         string
	TicketDesc   string
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (Metro) TableName() string {
	return "metro"
}
