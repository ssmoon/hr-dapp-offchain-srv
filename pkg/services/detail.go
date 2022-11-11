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

func GetWorker(workerId uint32) (*vo.WorkerItem, error) {
	db := conf.Db
	var worker vo.WorkerItem
	result := db.Model(&models.Worker{}).
		Select("tx.id, worker_name, birth_at, graduated_at, college_id, ty.college_name ").
		Joins("inner join college ty on ty.id = tx.college_id").
		Where("tx.id = ?", workerId).
		Scan(&worker)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("worker not found with id %d", workerId)
	}
	return &worker, nil
}

func GetCareerByWorkerId(workerId uint32) *[]vo.WorkerCareer {
	db := conf.Db
	var careers []vo.WorkerCareer
	db.Model(&models.WorkerCareer{}).
		Select("tx.id, tc.company_code, tc.company_name, tx.startAt, tx.endAt, tx.hasEnded ").
		Joins("inner join company tc on tc.id = tx.company_id").
		Where("tx.worker_id = ?", workerId).
		Order("tx.start_at").
		Scan(&careers)
	return &careers
}

func GetCertificateByWorkerId(workerId uint32) *[]vo.WorkerCertificate {
	db := conf.Db
	var certs []vo.WorkerCertificate
	db.Model(&models.WorkerCert{}).
		Select("tx.id, tc.cert_code, tc.cert_name, tx.acquired_at ").
		Joins("inner join certificate tc on tc.id = tx.certificate_id").
		Where("tx.worker_id = ?", workerId).
		Order("tx.acquired_at").
		Scan(&certs)
	return &certs
}

func CreateCertificate(workerId uint32, certificateId uint32, acquiredAt time.Time) error {
	db := conf.Db
	workerCert := &models.WorkerCert{}
	workerCert.CertificateID = certificateId
	workerCert.WorkerID = workerId
	workerCert.AcquiredAt = acquiredAt
	db.Create(&workerCert)

	var cert *models.Certificate
	result := db.First(&cert, certificateId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("cert not found with id %d", certificateId)
	}

	return _createOnChainCertficate(getSecurityNoByWorkerId(workerId), cert.CertCode, uint16(acquiredAt.Year()))
}

func CreateCareer(career *models.WorkerCareer) error {
	db := conf.Db
	db.Create(&career)

	var company *models.Company
	result := db.First(&company, career.CompanyID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("company not found with id %d", career.CompanyID)
	}

	return _createOnChainCareer(
		getSecurityNoByWorkerId(uint32(career.WorkerID)),
		company.CompanyCode,
		uint16(career.StartAt.Year()),
		uint16(career.EndAt.Year()),
		career.HasEnded == 1)
}

func FinishCareer(careerId uint32, endAt time.Time) error {
	db := conf.Db
	var career *models.WorkerCareer
	result := db.First(&career, careerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("career not found with id %d", careerId)
	}
	if career.HasEnded == 1 {
		return fmt.Errorf("career with id: %d has ended already", careerId)
	}
	db.Model(&models.WorkerCareer{}).Where("id = ?", careerId).Updates(map[string]any{"has_ended": 1, "end_at": endAt})

	return _finishCareerOnChain(getSecurityNoByWorkerId(uint32(career.WorkerID)), uint16(endAt.Year()))
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

func ValidateOnChain(workerId uint32) (*vo.ValidateResult, error) {
	vaidateResult := &vo.ValidateResult{}

	facade := conf.GetFacadeContract()

	securityNo := conf.Sha3StringToByte32(getSecurityNoByWorkerId(workerId))

	callOpts := conf.GetCallOpts(conf.GetCurrentUserAddr())
	onChainWorker, e1 := facade.GetWorkerBySecurityNo(callOpts, securityNo)
	if e1 != nil {
		return nil, e1
	}
	onChainCerts, e2 := facade.GetCertificateBySecurityNo(callOpts, securityNo)
	if e2 != nil {
		return nil, e2
	}
	onChainCareers, e3 := facade.GetWorkExperienceBySecurityNo(callOpts, securityNo)
	if e3 != nil {
		return nil, e3
	}

	db := conf.Db
	// validate worker data
	var offChainWorker *models.Worker
	result := db.First(&offChainWorker, workerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("worker not found with id %d", workerId)
	}

	var college *models.College
	result = db.First(&college, offChainWorker.CollegeID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("college not found with id %d", offChainWorker.CollegeID)
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
		StartAt     uint16
		EndAt       uint16
		HasEnded    bool
		CompanyCode string
	}

	var careerResults []careerItem
	db.Table("worker_career ta").Select("year(start_at) as start_at, year(end_at) as end_at, has_ended, company_code").Joins("inner join company tb on ta.company_id=tb.id").Where("worker_id = ?", workerId).Order("start_at").Scan(&careerResults)

	offChainCareers := make([]contracts.WorkExperienceDefineWorkExperience, len(careerResults))
	for i, x := range careerResults {
		item := contracts.WorkExperienceDefineWorkExperience{}
		item.StartAt = x.StartAt
		item.EndAt = x.EndAt
		item.HasEnded = x.HasEnded
		item.CompanyCode = conf.ConvertStringToByte32(x.CompanyCode)
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
		AcquiredAt uint16
		CertCode   string
	}

	var certResults []certItem
	db.Raw("Select year(acquired_at) as acquired_at,  cert_code from worker_cert ta inner join certificate tb on ta.certificate_id=tb.id where worker_id=? order by cert_code", workerId).Scan(&certResults)

	offChainCerts := make([]contracts.CertificateDefineCertificate, len(certResults))
	for i, x := range certResults {
		item := contracts.CertificateDefineCertificate{}
		item.AcquiredAt = x.AcquiredAt
		item.CertCode = conf.ConvertStringToByte32(x.CertCode)
		offChainCerts[i] = item
	}
	offChainCertHash := conf.ComputeCertHash(offChainCerts)
	onChainCertHash := conf.ComputeCertHash(onChainCerts)
	if offChainCertHash == onChainCertHash {
		vaidateResult.CertificateSame = consts.Sync_Result_Same
	} else {
		vaidateResult.CertificateSame = consts.Sync_Result_NotSame
	}

	return vaidateResult, nil
}

func _finishCareerOnChain(securityNo string, endAt uint16) error {
	facade := conf.GetFacadeContract()
	tx, err := facade.FinishLastCareer(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(securityNo), endAt)
	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	return err
}

func _createOnChainCertficate(securityNo string, certCode string, acquiredAt uint16) error {
	facade := conf.GetFacadeContract()

	cert := contracts.CertificateDefineCertificate{}
	cert.AcquiredAt = acquiredAt
	cert.CertCode = conf.ConvertStringToByte32(certCode)

	tx, err := facade.CreateCertificate(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(securityNo), cert)

	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	return err
}

func _createOnChainCareer(securityNo string, companyCode string, startAt uint16, endAt uint16, hasEnded bool) error {
	facade := conf.GetFacadeContract()

	career := contracts.WorkExperienceDefineWorkExperience{}
	career.StartAt = startAt
	career.EndAt = endAt
	career.HasEnded = hasEnded
	career.CompanyCode = conf.ConvertStringToByte32(companyCode)

	tx, err := facade.AddWorkExperience(conf.GetTransactionOpts(conf.GetCurrentUserKey(), conf.GetCurrentUserAddr()), conf.Sha3StringToByte32(securityNo), career)

	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	return err
}
