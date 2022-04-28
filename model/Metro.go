package model

import "time"

type Metro struct {
	Id           int       `json:"id"`
	ChnName      string    `json:"chnName" binding:"required" db:"chn_name"`
	EngName      string    `json:"engName" binding:"required" db:"eng_name"`
	ShortName    string    `json:"shortName" binding:"required" db:"short_name"`
	FirstLetter  string    `json:"FirstLetter" binding:"required" db:"first_letter"`
	Icon         string    `json:"icon" binding:"required" db:"icon"`
	OperatePic   string    `json:"operatePic" binding:"required" db:"operate_pic"`
	PlanPic      string    `json:"planPic" binding:"required" db:"plan_pic"`
	Website      string    `json:"website" db:"website"`
	City         string    `json:"city" binding:"required" db:"city"`
	Length       float32   `json:"length" binding:"required" db:"length"`
	LineNum      int       `json:"lineNum" binding:"required" db:"line_num"`
	StandNum     int       `json:"standNum" binding:"required" db:"stand_num"`
	OpenDate     string    `json:"openDate" binding:"required" db:"open_date"`
	BuildDate    string    `json:"buildDate" binding:"required" db:"build_date"`
	BuildCompany string    `json:"buildCompany" binding:"required" db:"build_company"`
	Remark       string    `json:"remark" db:"remark"`
	TicketDesc   string    `json:"ticketDesc" db:"ticket_desc"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
	UpdateTime   time.Time `json:"updateTime" db:"update_time"`
	DeleteStatus int       `json:"deleteStatus" db:"delete_status"`
}

//func (Metro) TableName() string {
//	return "metro"
//}
