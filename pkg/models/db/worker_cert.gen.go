// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameWorkerCert = "worker_cert"

// WorkerCert mapped from table <worker_cert>
type WorkerCert struct {
	ID            uint32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	WorkerID      uint32  `gorm:"column:worker_id;not null" json:"workerId"`
	CertificateID uint32  `gorm:"column:certificate_id;not null" json:"certificateId"`
	AcquiredAt	  time.Time `gorm:"column:acquiredAt;not null" json:"acquiredAt"`
}

// TableName WorkerCert's table name
func (*WorkerCert) TableName() string {
	return TableNameWorkerCert
}
