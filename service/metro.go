package service

import (
	"metro/dao"
	"metro/model"
)

func AddMetro(metro *model.Metro) (id int64) {
	return dao.InsertMetros(metro)
}
