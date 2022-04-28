package service

import (
	"metro/controller/req"
	"metro/dao"
	"metro/model"
)

func AddMetro(metro *model.Metro) (id int64) {
	return dao.InsertMetros(metro)
}

func DelMetros(id int64) {
	dao.DelMetros(id)
}

func ListMetros(param *req.MetroListReq) []model.Metro {
	return dao.ListMetros(param)
}
