package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"hr-dapp/srv/pkg/conf"
	"hr-dapp/srv/pkg/consts"
	"hr-dapp/srv/pkg/contracts"
	models "hr-dapp/srv/pkg/models/db"
	"hr-dapp/srv/pkg/models/vo"
	"log"
	"time"

	"gorm.io/gorm"
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

	var cert *models.Certificate
	result := db.First(&cert, certificateId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprint("cert not found with id ", certificateId))
	}

	_createOnChainCertficate(getSecurityNoByWorkerId(workerId), cert.CertCode, uint16(acquiredAt.Year()))
}

func CreateCareer(career *models.WorkerCareer) {
	db := conf.Db
	db.Create(&career)

	var company *models.Company
	result := db.First(&company, career.CompanyID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprint("Company not found with id ", career.CompanyID))
	}

	_createOnChainCareer(
		getSecurityNoByWorkerId(uint32(career.WorkerID)),
		company.CompanyCode,
		uint16(career.StartAt.Year()),
		uint16(career.EndAt.Year()),
		career.HasEnded == 1)
}

func FinishCareer(careerId uint32, endAt time.Time) {
	db := conf.Db
	var career *models.WorkerCareer
	result := db.First(&career, careerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprintf("career not found with id %d", careerId))
	}
	if career.HasEnded == 0 {
		panic(fmt.Sprintf("career with id: %d has ended already", careerId))
	}
	db.Model(&models.WorkerCareer{}).Where("id = ?", careerId).Updates(map[string]any{"has_ended": 1, "end_at": endAt})

	_finishCareerOnChain(getSecurityNoByWorkerId(uint32(career.WorkerID)), uint16(endAt.Year()))
}

func getSecurityNoByWorkerId(workerId uint32) string {
	db := conf.Db
	worker := models.Worker{}
	result := db.First(&worker, workerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprintf("worker not found with id %d", workerId))
	}
	return worker.SecurityNo
}

func ValidateOnChain(workerId uint32) *vo.ValidateResult {
	vaidateResult := &vo.ValidateResult{}

	facade := conf.GetFacadeContract()

	securityNo := conf.Sha3StringToByte32(getSecurityNoByWorkerId(workerId))

	callOpts := conf.GetCallOpts(conf.GetCurrentUserAddr())
	onChainWorker, e1 := facade.GetWorkerBySecurityNo(callOpts, securityNo)
	if e1 != nil {
		panic(e1)
	}
	onChainCerts, e2 := facade.GetCertificateBySecurityNo(callOpts, securityNo)
	if e2 != nil {
		panic(e2)
	}
	onChainCareers, e3 := facade.GetWorkExperienceBySecurityNo(callOpts, securityNo)
	if e3 != nil {
		panic(e3)
	}

	db := conf.Db
	// validate worker data
	var offChainWorker *models.Worker
	result := db.First(&offChainWorker, workerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprintf("worker not found with id %d", workerId))
	}

	var college *models.College
	result = db.First(&college, offChainWorker.CollegeID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(fmt.Sprint("college not found with id ", string(offChainWorker.CollegeID)))
	}

	onChainWorkerHash := conf.ComputeWorkerHash(onChainWorker.BirthAt, onChainWorker.GraduatedAt, string(onChainWorker.CollegeCode[:]))
	offChainWorkerHash := conf.ComputeWorkerHash(uint16(offChainWorker.BirthAt.Year()), uint16(offChainWorker.GraduatedAt.Year()), college.CollegeCode)
	if onChainWorkerHash == offChainWorkerHash {
		vaidateResult.WorkerSame = consts.Sync_Result_Same
	} else {
		vaidateResult.WorkerSame = consts.Sync_Result_NotSame
	}

	//validate career data
	type careerItem struct {
		startAt     uint16
		endAt       uint16
		hasEnded    bool
		companyCode string
	}

	var careerResults []careerItem
	db.Raw("Select start_at, end_at, has_ended, company_code from worker_career ta inner join company tb on ta.company_id=tb.id where worker_id=? order by start_at", workerId).Scan(&careerResults)

	offChainCareers := make([]contracts.WorkExperienceDefineWorkExperience, len(careerResults))
	for i, x := range careerResults {
		item := contracts.WorkExperienceDefineWorkExperience{}
		item.StartAt = x.startAt
		item.EndAt = x.endAt
		item.HasEnded = x.hasEnded
		item.CompanyCode = conf.ConvertStringToByte32(x.companyCode)
		offChainCareers[i] = item
	}
	offChainCareerHash := conf.ComputeCareerHash(offChainCareers)
	onChainCareerHash := conf.ComputeCareerHash(onChainCareers)
	if offChainCareerHash == onChainCareerHash {
		vaidateResult.CareerSame = consts.Sync_Result_Same
	} else {
		vaidateResult.CareerSame = consts.Sync_Result_NotSame
	}

	//validate certs
	type certItem struct {
		acquiredAt uint16
		certCode   string
	}

	var certResults []certItem
	db.Raw("Select acquired_at,  cert_code from worker_cert ta inner join certificate tb on ta.certificate_id=tb.id where worker_id=1 order by cert_code", workerId).Scan(&certResults)

	offChainCerts := make([]contracts.CertificateDefineCertificate, len(certResults))
	for i, x := range certResults {
		item := contracts.CertificateDefineCertificate{}
		item.AcquiredAt = x.acquiredAt
		item.CertCode = conf.ConvertStringToByte32(x.certCode)
		offChainCerts[i] = item
	}
	offChainCertHash := conf.ComputeCertHash(offChainCerts)
	onChainCertHash := conf.ComputeCertHash(onChainCerts)
	if offChainCertHash == onChainCertHash {
		vaidateResult.CertificateSame = consts.Sync_Result_Same
	} else {
		vaidateResult.CertificateSame = consts.Sync_Result_NotSame
	}

	return vaidateResult
}

func _finishCareerOnChain(securityNo string, endAt uint16) {
	facade := conf.GetFacadeContract()
	tx, err := facade.FinishLastCareer(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(securityNo), endAt)
	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	if err != nil {
		log.Fatal(err)
	}
}

func _createOnChainCertficate(securityNo string, certCode string, acquiredAt uint16) {
	facade := conf.GetFacadeContract()

	cert := contracts.CertificateDefineCertificate{}
	cert.AcquiredAt = acquiredAt
	cert.CertCode = conf.ConvertStringToByte32(certCode)

	tx, err := facade.CreateCertificate(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(securityNo), cert)

	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	if err != nil {
		log.Fatal(err)
	}
}

func _createOnChainCareer(securityNo string, companyCode string, startAt uint16, endAt uint16, hasEnded bool) {
	facade := conf.GetFacadeContract()

	career := contracts.WorkExperienceDefineWorkExperience{}
	career.StartAt = startAt
	career.EndAt = endAt
	career.HasEnded = hasEnded
	career.CompanyCode = conf.ConvertStringToByte32(companyCode)

	tx, err := facade.AddWorkExperience(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(securityNo), career)

	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	if err != nil {
		log.Fatal(err)
	}
}
