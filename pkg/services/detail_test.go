package services_test

import (
	"hr-dapp/srv/pkg/conf"
	"hr-dapp/srv/pkg/consts"
	models "hr-dapp/srv/pkg/models/db"
	services "hr-dapp/srv/pkg/services"
	"log"
	"testing"
	"time"

	"github.com/golang-module/carbon/v2"
)

var worker *models.Worker
var lastCareerId uint32

func setupTest() {
	log.Print("setup method invoked")
	conf.InitEnv("C:\\projects\\hr-dapp-offchain-srv")
	conf.InitContractConfig("C:\\projects\\hr-dapp-offchain-srv")
	conf.InitDatabase()
	worker = createNewWorker()
	log.Print("test setup completed")
}

func TestDetail(t *testing.T) {
	setupTest()
	testCreateCertificate(t)
	testCreateCareer(t)
	testFinishCareer(t)
	testValidateOnChain(t)
}

func testCreateCertificate(t *testing.T) {
	t.Run("should successfully create the certificate", func(t *testing.T) {
		err := services.CreateCertificate(worker.ID, 1, carbon.CreateFromDate(2021, 4, 11).Carbon2Time())
		if err != nil {
			t.Fatalf(err.Error())
		}
	})
}

func testCreateCareer(t *testing.T) {
	t.Run("should successfully create the career", func(t *testing.T) {
		career := models.WorkerCareer{}
		career.CompanyID = 1
		career.StartAt = carbon.CreateFromDate(2020, 3, 1).Carbon2Time()
		career.EndAt = carbon.CreateFromDate(2022, 4, 1).Carbon2Time()
		career.HasEnded = 0
		career.WorkerID = worker.ID

		err := services.CreateCareer(&career)
		if err != nil {
			t.Fatalf(err.Error())
		}
		lastCareerId = career.ID
	})
}

func testFinishCareer(t *testing.T) {
	t.Run("should finish current unclosed career", func(t *testing.T) {
		time.Sleep(time.Duration(5) * time.Second)
		err := services.FinishCareer(lastCareerId, carbon.CreateFromDate(2022, 8, 22).Carbon2Time())
		if err != nil {
			t.Fatalf(err.Error())
		}
	})
}

func testValidateOnChain(t *testing.T) {
	t.Run("chain data should accord with db", func(t *testing.T) {
		time.Sleep(time.Duration(5) * time.Second)
		result, err := services.ValidateOnChain(worker.ID)
		if err != nil {
			t.Errorf(err.Error())
		}
		if result.WorkerSame != consts.Sync_Result_Same {
			t.Errorf("Sync Result of Worker differs from [same], but %d", result.WorkerSame)
		}
		if result.CertificateSame != consts.Sync_Result_Same {
			t.Errorf("Sync Result of Certificate differs from [same], but %d", result.CertificateSame)
		}
		if result.CareerSame != consts.Sync_Result_Same {
			t.Errorf("Sync Result of Career differs from [same], but %d", result.CareerSame)
		}
	})

}

func createNewWorker() *models.Worker {
	worker := &models.Worker{}
	worker.BirthAt = carbon.CreateFromDate(1996, 9, 1).Carbon2Time()
	worker.GraduatedAt = carbon.CreateFromDate(2018, 7, 1).Carbon2Time()
	worker.CollegeID = 1
	worker.WorkerName = "秦先海"
	worker.SecurityNo = randomString(18)
	_, err := services.CreateWorker(worker)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(5) * time.Second)
	return worker
}
