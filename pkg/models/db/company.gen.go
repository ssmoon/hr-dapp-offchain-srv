// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameCompany = "company"

// Company mapped from table <company>
type Company struct {
	ID          uint32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CompanyCode string  `gorm:"column:company_code;not null" json:"companyCode"`
	CompanyName string `gorm:"column:company_name;not null" json:"companyName"`
}

// TableName Company's table name
func (*Company) TableName() string {
	return TableNameCompany
}
