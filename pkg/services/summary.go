package services

import (
	"fmt"
	"hr-dapp/srv/pkg/conf"
	models "hr-dapp/srv/pkg/models/db"
)

func GetAllWorkers() []models.Worker {
	db := conf.Db
	var workers []models.Worker
	db.Find(&workers)

	for _, v := range workers {
		v.SecurityNo = fmt.Sprintf("%s********%s", v.SecurityNo[0:6], v.SecurityNo[14:18])
	}

	return workers
}

func CreateWorker(worker *models.Worker) {
	db := conf.Db
	db.Create(&worker)
}
