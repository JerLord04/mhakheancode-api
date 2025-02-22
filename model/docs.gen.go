// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDoc = "docs"

// Doc mapped from table <docs>
type Doc struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Image       string    `gorm:"column:image" json:"image"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Link        string    `gorm:"column:link" json:"link"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Doc's table name
func (*Doc) TableName() string {
	return TableNameDoc
}
