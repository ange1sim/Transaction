package utils

import (
	"log"
	"transaction/db"
	"transaction/model"
	"transaction/request"
)

func CheckDataVendors() {
	var (
		trnx []model.ShowTrn
		err  error
		status int
	)

	status = 400

	err = db.DB.Model(&model.Trnx{}).Where("status = ? and created_at >= CURRENT_TIMESTAMP - interval '24 hours'", status).Select("vendor, count(vendor) as total").Group("vendor").Find(&trnx).Error
	if err != nil {
		log.Println("ошибка при получении данных из ДБ!")
	}

	request.ReqServer(trnx)
}
