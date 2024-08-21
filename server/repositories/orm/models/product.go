package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/quocbang/grpc-gateway/pkg/pb"
)

type Capacity struct {
	Size int64       `json:"size"`
	Unit pb.SizeUnit `json:"unit"`
}

func (c *Capacity) Scan(src any) error {
	return ScanJSON(src, c)
}

func (c Capacity) Value() (driver.Value, error) {
	return json.Marshal(c)
}

type Advance struct {
	Rom *Capacity `json:"rom"`
	Ram *Capacity `json:"ram"`
	CPU float32   `json:"cpu"`
}

func (a *Advance) Scan(src any) error {
	return ScanJSON(src, a)
}

func (a Advance) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type Product struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ProductCode string    `gorm:"type:text"`
	Color       string    `gorm:"type:text"`
	Producer    string    `gorm:"type:text"`
	Series      string    `gorm:"type:text"`
	AdvanceInfo *Advance  `gorm:"type:jsonb;default:'{}'"`
}

func (Product) TableName() string {
	return "product"
}
