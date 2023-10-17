package models

import (
	"github.com/quocbang/grpc-gateway/server/utils/roles"
)

type Account struct {
	Username       string        `gorm:"type:text;primaryKey"`
	Email          string        `gorm:"type:text;uniqueIndex:idx_email;NOT NULL"`
	IsUserVerified bool          `gorm:"NOT NULL;default:false"`
	Password       []byte        `gorm:"type:bytea;NOT NULL"`
	Role           roles.Roles   `gorm:"type:smallint;NOT NULL;default:0"`
	AccountVerify  AccountVerify `gorm:"foreignKey:Username;references:Username"`
}

func (Account) TableName() string {
	return "account"
}

type AccountVerify struct {
	Username   string `gorm:"type:text;NOT NULL"`
	SecretCode string `gorm:"NOT NULL"`
	CreatedAt  int64  `gorm:"NOT NULL;autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime:nano"`
}

func (AccountVerify) TableName() string {
	return "verify_account"
}
