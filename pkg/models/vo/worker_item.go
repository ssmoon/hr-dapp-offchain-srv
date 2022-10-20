package vo

import "time"

type WorkerItem struct {
	WorkerId         int       `json:"workerId"`
	WorkerName       string    `json:"workerName"`
	GraduatedAt      time.Time `json:"graduatedAt"`
	BirthAt          time.Time `json:"birthAt"`
	CertificateCount int       `json:"certificateCount"`
}
