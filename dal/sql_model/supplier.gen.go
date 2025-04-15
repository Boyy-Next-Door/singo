// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql_model

import (
	"time"
)

const TableNameSupplier = "supplier"

// Supplier mapped from table <supplier>
type Supplier struct {
	ID              int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;type:varchar(128);not null" json:"name"`
	Code            string    `gorm:"column:code;type:varchar(64);not null;uniqueIndex:idx_code,priority:1" json:"code"`
	ContactName     string    `gorm:"column:contact_name;type:varchar(64);not null" json:"contact_name"`
	ContactPhone    string    `gorm:"column:contact_phone;type:varchar(20);not null" json:"contact_phone"`
	Address         string    `gorm:"column:address;type:varchar(255);not null" json:"address"`
	Longitude       float64   `gorm:"column:longitude;type:decimal(10,7);index:idx_location,priority:1" json:"longitude"`
	Latitude        float64   `gorm:"column:latitude;type:decimal(10,7);index:idx_location,priority:2" json:"latitude"`
	BusinessScope   string    `gorm:"column:business_scope;type:varchar(255);not null" json:"business_scope"`
	BusinessLicense string    `gorm:"column:business_license;type:varchar(128)" json:"business_license"`
	BankAccount     string    `gorm:"column:bank_account;type:varchar(64)" json:"bank_account"`
	BankName        string    `gorm:"column:bank_name;type:varchar(128)" json:"bank_name"`
	AccountHolder   string    `gorm:"column:account_holder;type:varchar(64)" json:"account_holder"`
	Status          int32     `gorm:"column:status;type:tinyint;not null;default:1;comment:0: 1:" json:"status"` // 0: 1:
	CreateTime      time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName Supplier's table name
func (*Supplier) TableName() string {
	return TableNameSupplier
}
