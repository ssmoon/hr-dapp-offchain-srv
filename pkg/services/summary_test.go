package services_test

import (
	"fmt"
	"hr-dapp/srv/pkg/consts"
	models "hr-dapp/srv/pkg/models/db"
	services "hr-dapp/srv/pkg/services"
	"math/rand"
	"testing"
	"time"

	"github.com/golang-module/carbon/v2"
)

func TestCreateWorker(t *testing.T) {
	var newWorkerId uint32
	newWorkerSecurityNo := randomString(18)

	t.Run("create worker", func(t *testing.T) {
		worker := &models.Worker{}
		worker.BirthAt = carbon.CreateFromDate(1996, 9, 1).Carbon2Time()
		worker.GraduatedAt = carbon.CreateFromDate(2018, 7, 1).Carbon2Time()
		worker.CollegeID = 1
		worker.WorkerName = "秦先海"
		worker.SecurityNo = newWorkerSecurityNo
		err := services.CreateWorker(worker)
		if err != nil {
			t.Fatalf(err.Error())
		}
		newWorkerId = worker.ID
	})
	t.Run("worker should be both on db & chain", func(t *testing.T) {
		syncResult := services.SyncOnChain(newWorkerId)
		if syncResult != consts.Sync_Result_Same {
			t.Error("Sync Result differs from [same], but %d", syncResult)
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
			t.Error("newly appended worker does not exist, worker id: %d", newWorkerId)
		}
	})
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}