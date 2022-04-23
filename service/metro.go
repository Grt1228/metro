package service

import (
	"metro/dao"
	"metro/model"
)

func AddMetro(metro *model.Metro) (id int) {
	return dao.InsertMetros(metro)
}
