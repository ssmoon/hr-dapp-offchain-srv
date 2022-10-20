package services

import (
	"hr-dapp/srv/pkg/conf"
	models "hr-dapp/srv/pkg/models/db"
	"hr-dapp/srv/pkg/models/vo"
	"time"
)

func GetWorkerDetail(workerId uint32) vo.WorkerDetail {
	db := conf.Db
	workerDetail := vo.WorkerDetail{}
	db.First(&workerDetail.General, workerId)
	db.Model(&models.WorkerCareer{}).
		Select("tx.id, tc.company_code, tc.company_name, tx.startAt, tx.endAt, tx.hasEnded ").
		Joins("inner join company tc on tc.id = tx.company_id").
		Where("tx.worker_id = ?", workerId).
		Order("tx.start_at").
		Scan(&workerDetail.CareerPeriods)
	db.Model(&models.WorkerCert{}).
		Select("tx.id, tc.cert_code, tc.cert_name, tx.acquired_at ").
		Joins("inner join certificate tc on tc.id = tx.certificate_id").
		Where("tx.worker_id = ?", workerId).
		Order("tx.acquired_at").
		Scan(&workerDetail.Certificates)
	return workerDetail
}

func CreateCertificate(workerId uint32, certificateId uint32, acquiredAt time.Time) {
	db := conf.Db
	workerCert := &models.WorkerCert{}
	workerCert.CertificateID = certificateId
	workerCert.WorkerID = workerId
	workerCert.AcquiredAt = acquiredAt
	db.Create(&workerCert)
}

func CreateCareer(career *models.WorkerCareer) {
	db := conf.Db
	db.Create(&career)
}

func FinishCareer(careerId uint32) {
	db := conf.Db
	db.Model(&models.WorkerCareer{}).Where("id = ?", careerId).Update("has_ended", 1)
}

func ValidateOnChain(workerId uint32) vo.WorkerDetail {
	return vo.WorkerDetail{}
}
