package services_test

import (
	"fmt"
	"hr-dapp/srv/pkg/conf"
	"hr-dapp/srv/pkg/consts"
	models "hr-dapp/srv/pkg/models/db"
	services "hr-dapp/srv/pkg/services"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/golang-module/carbon/v2"
)

func setupSummary() {
	log.Print("setup method invoked")
	conf.InitEnv("C:\\projects\\hr-dapp-offchain-srv")
	conf.InitContractConfig("C:\\projects\\hr-dapp-offchain-srv")
	conf.InitDatabase()
	services.GrantUserPrivilege(conf.GetCurrentUserAddr())
	log.Print("test setup completed")
}

func TestSummary(t *testing.T) {
	setupSummary()
	var newWorkerId uint32
	newWorkerSecurityNo := randomString(18)

	t.Run("create worker", func(t *testing.T) {
		worker := &models.Worker{}
		worker.BirthAt = carbon.CreateFromDate(1996, 9, 1).Carbon2Time()
		worker.GraduatedAt = carbon.CreateFromDate(2018, 7, 1).Carbon2Time()
		worker.CollegeID = 1
		worker.WorkerName = "秦先海"
		worker.SecurityNo = newWorkerSecurityNo
		newWorker, err := services.CreateWorker(worker)
		if err != nil {
			t.Fatalf(err.Error())
		}
		newWorkerId = newWorker.ID
		// wait for transaction complete, or the following func can't get the updated data
		time.Sleep(time.Duration(5) * time.Second)
	})
	t.Run("worker should be both on db & chain", func(t *testing.T) {
		syncResult, err := services.SyncOnChain(newWorkerId)
		if err != nil {
			t.Errorf(err.Error())
		}
		if syncResult != consts.Sync_Result_Same {
			t.Errorf("Sync Result differs from [same], but %d", syncResult)
		}
	})
	t.Run("should get worker list from db", func(t *testing.T) {
		workers := services.GetAllWorkers()
		exist := false
		for _, v := range workers {
			if v.ID == newWorkerId {
				exist = true
			}
		}
		if !exist {
			t.Errorf("newly appended worker does not exist, worker id: %d", newWorkerId)
		}
	})
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
