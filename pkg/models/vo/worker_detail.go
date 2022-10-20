package vo

import models "hr-dapp/srv/pkg/models/db"

type WorkerDetail struct {
	General       models.Worker         `json:"general"`
	CareerPeriods []models.WorkerCareer `json:"careerPeriods"`
	Certificates  []models.Certificate  `json:"certificates"`
}
