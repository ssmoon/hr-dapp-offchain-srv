package services_test

import (
	models "hr-dapp/srv/pkg/models/db"
	services "hr-dapp/srv/pkg/services"
	"log"
	"testing"

	"github.com/golang-module/carbon/v2"
)

var worker *models.Worker

func setupTest() {
	worker = createNewWorker()
	log.Print("test setup completed")
}

func TestCreateCertificate(t *testing.T) {
	t.Run("should successfully create the certificate", func(t *testing.T) {
		err := services.CreateCertificate(worker.ID, 1, carbon.CreateFromDate(2020, 9, 1).Carbon2Time())
		if err != nil {
			t.Fatalf(err.Error())
		}
	})
}

func TestCreateCareer(t *testing.T) {
	t.Run("should successfully create the career", func(t *testing.T) {
		career := models.WorkerCareer{}
		career.CompanyID = 1
		career.StartAt = carbon.CreateFromDate(2020, 3, 1).Carbon2Time()
		career.EndAt = carbon.CreateFromDate(2022, 4, 1).Carbon2Time()
		career.HasEnded = 1
		career.WorkerID = worker.ID

		err := services.CreateCareer(&career)
		if err != nil {
			t.Fatalf(err.Error())
		}
	})
}

func TestGetWorkerDetail(t *testing.T) {

	t.Run("create worker")
	worker := &models.Worker{}
	worker.BirthAt = carbon.CreateFromDate(1996, 9, 1).Carbon2Time()
	worker.GraduatedAt = carbon.CreateFromDate(2018, 7, 1).Carbon2Time()
	worker.CollegeID = 1
	worker.WorkerName = "张奉先"
	worker.SecurityNo = "320103199609013241"
	services.CreateWorker(worker)

}

func createNewWorker() *models.Worker {
	worker := &models.Worker{}
	worker.BirthAt = carbon.CreateFromDate(1996, 9, 1).Carbon2Time()
	worker.GraduatedAt = carbon.CreateFromDate(2018, 7, 1).Carbon2Time()
	worker.CollegeID = 1
	worker.WorkerName = "秦先海"
	worker.SecurityNo = randomString(18)
	err := services.CreateWorker(worker)
	if err != nil {
		panic(err)
	}
	return worker
}
