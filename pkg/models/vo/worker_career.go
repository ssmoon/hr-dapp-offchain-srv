package vo

import "time"

type WorkerCareer struct {
	ID          uint32    `json:"id"`
	CompanyCode string    `json:"company_code"`
	CompanyName string    `json:"companyId"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	HasEnded    int32     `json:"hasEnded"`
}
