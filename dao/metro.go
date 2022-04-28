package dao

import (
	"fmt"
	"metro/controller/req"
	"metro/model"
)

func InsertMetros(metro *model.Metro) (id int64) {
	insertSql := `insert into metro (chn_name, eng_name, short_name, first_letter, icon, operate_pic, plan_pic, website, city, length, line_num, stand_num, open_date, build_date, build_company, remark, ticket_desc) ` +
		`VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	exec := db.MustExec(insertSql, metro.ChnName, metro.EngName, metro.ShortName, metro.FirstLetter, metro.Icon, metro.OperatePic, metro.PlanPic, metro.Website,
		metro.City, metro.Length, metro.LineNum, metro.StandNum, metro.OpenDate, metro.BuildDate, metro.BuildCompany, metro.Remark, metro.TicketDesc)
	id, _ = exec.LastInsertId()
	return id
}

func DelMetros(id int64) {
	sql := "update metro set delete_status = 1 where id = ?"
	db.MustExec(sql, id)
}

func ListMetros(param *req.MetroListReq) []model.Metro {

	sql := "select * from metro where chn_name like CONCAT('%',?,'%') order by id desc limit ?,?"
	var metros []model.Metro
	if err := db.Select(&metros, sql, param.Name, param.Page, param.Size); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("users:%#v\n", metros)
	return metros
}
