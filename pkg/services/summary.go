package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"hr-dapp/srv/pkg/conf"
	"hr-dapp/srv/pkg/consts"
	"hr-dapp/srv/pkg/contracts"
	models "hr-dapp/srv/pkg/models/db"
	"log"

	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

func GetAllWorkers() []models.Worker {
	db := conf.Db
	var workers []models.Worker
	db.Find(&workers)

	for _, v := range workers {
		v.SecurityNo = fmt.Sprint(v.SecurityNo[0:6], "********", v.SecurityNo[14:18])
	}

	return workers
}

func CreateWorker(worker *models.Worker) error {
	db := conf.Db
	db.Create(&worker)

	var collegeByCode *models.College
	result := db.First(&collegeByCode, worker.CollegeID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf(fmt.Sprint("college not found with id ", string(worker.CollegeID)))
	}

	err := _createOnChainWorker(worker, collegeByCode.CollegeCode)
	return err
}

func SyncOnChain(workerId uint32) uint8 {
	facade := conf.GetFacadeContract()

	var offChainWorker *models.Worker
	db := conf.Db
	result := db.First(&offChainWorker, workerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprintf("worker not found with id %d", workerId))
	}

	var college *models.College
	result = db.First(&college, offChainWorker.CollegeID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprint("college not found with id ", string(offChainWorker.CollegeID)))
	}

	onChainWorker, err := facade.GetWorkerBySecurityNo(conf.GetCallOpts(conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(offChainWorker.SecurityNo))

	if err != nil {
		panic(err)
	}

	if onChainWorker.IsValue {
		onChainHash := conf.ComputeWorkerHash(onChainWorker.BirthAt, onChainWorker.GraduatedAt, string(onChainWorker.CollegeCode[:]))
		offChainHash := conf.ComputeWorkerHash(uint16(offChainWorker.BirthAt.Year()), uint16(offChainWorker.GraduatedAt.Year()), college.CollegeCode)
		if onChainHash == offChainHash {
			return consts.Sync_Result_Same
		} else {
			var collegeByCode *models.College
			result = db.First(&collegeByCode, offChainWorker.CollegeID)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				panic(fmt.Sprint("college not found with id ", string(offChainWorker.CollegeID)))
			}
			offChainWorker.BirthAt = carbon.Time2Carbon(offChainWorker.BirthAt).SetYear(int(onChainWorker.BirthAt)).Carbon2Time()
			offChainWorker.GraduatedAt = carbon.Time2Carbon(offChainWorker.GraduatedAt).SetYear(int(onChainWorker.GraduatedAt)).Carbon2Time()
			offChainWorker.CollegeID = int32(collegeByCode.ID)
			db.Save(offChainWorker)
			return consts.Sync_Result_OverrideOffChain
		}
	} else {
		_createOnChainWorker(offChainWorker, college.CollegeCode)
		return consts.Sync_Result_UploadToChain
	}
}

func _createOnChainWorker(offChainWorker *models.Worker, collegeCode string) error {
	facade := conf.GetFacadeContract()
	worker := contracts.WorkerDefineWorker{}

	worker.SecurityNo = conf.Sha3StringToByte32(offChainWorker.SecurityNo)
	worker.BirthAt = uint16(offChainWorker.BirthAt.Year())
	worker.GraduatedAt = uint16(offChainWorker.GraduatedAt.Year())
	worker.IsValue = true
	worker.CollegeCode = conf.ConvertStringToByte32(collegeCode)
	tx, err := facade.CreateWorker(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), worker)

	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	return err
}
