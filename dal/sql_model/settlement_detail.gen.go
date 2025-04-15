// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql_model

import (
	"time"
)

const TableNameSettlementDetail = "settlement_detail"

// SettlementDetail mapped from table <settlement_detail>
type SettlementDetail struct {
	ID             int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	SettlementID   int64     `gorm:"column:settlement_id;type:bigint;not null;index:idx_settlement,priority:1;comment:ID" json:"settlement_id"` // ID
	VerifyID       int64     `gorm:"column:verify_id;type:bigint;not null;index:idx_verify,priority:1;comment:ID" json:"verify_id"`             // ID
	CouponID       int64     `gorm:"column:coupon_id;type:bigint;not null;index:idx_coupon,priority:1;comment:ID" json:"coupon_id"`             // ID
	CouponCode     string    `gorm:"column:coupon_code;type:varchar(64);not null" json:"coupon_code"`
	VerifyAmount   float64   `gorm:"column:verify_amount;type:decimal(10,2);not null" json:"verify_amount"`
	PlatformFee    float64   `gorm:"column:platform_fee;type:decimal(10,2);not null" json:"platform_fee"`
	SupplierAmount float64   `gorm:"column:supplier_amount;type:decimal(10,2);not null" json:"supplier_amount"`
	VerifyTime     time.Time `gorm:"column:verify_time;type:datetime;not null" json:"verify_time"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName SettlementDetail's table name
func (*SettlementDetail) TableName() string {
	return TableNameSettlementDetail
}
