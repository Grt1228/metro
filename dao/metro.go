package dao

import (
	"metro/model"
)

func InsertMetros(metro *model.Metro) (id int64) {
	insertSql := `insert into metro (chn_name, eng_name, short_name, first_letter, icon, operate_pic, plan_pic, website, city, length, line_num, stand_num, open_date, build_date, build_company, remark, ticket_desc) ` +
		`VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	exec := db.MustExec(insertSql, metro.ChnName, metro.EngName, metro.ShortName, metro.FirstLetter, metro.Icon, metro.OperatePic, metro.PlanPic, metro.Website,
		metro.City, metro.Length, metro.LineNum, metro.StandNum, metro.OpenDate, metro.BuildDate, metro.BuildCompany, metro.Remark, metro.TicketDesc)
	println(exec.LastInsertId())
	id, _ = exec.LastInsertId()
	return id
}
