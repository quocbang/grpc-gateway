package models

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/quocbang/grpc-gateway/server/repositories/errors"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"gorm.io/gorm"
)

type Account struct {
	Username       string        `gorm:"type:text;primaryKey"`
	Email          string        `gorm:"type:text;uniqueIndex:idx_email;NOT NULL"`
	IsUserVerified bool          `gorm:"NOT NULL;default:false"`
	Password       []byte        `gorm:"type:bytea;NOT NULL"`
	Role           roles.Roles   `gorm:"type:smallint;NOT NULL;default:0"`
	AccountVerify  VerifyAccount `gorm:"foreignKey:Username;references:Username"`
}

// BeforeUpdate check before update roles
func (a *Account) BeforeUpdate(tx *gorm.DB) error {
	if a.Role != 0 {
		account := Account{}
		if err := tx.Where("username = ?", a.Username).Take(&account).Error; err != nil {
			// TODO: should check err
			return nil
		}

		if account.Role >= a.Role {
			return errors.Error{
				Code:    errors.Code_COULD_NOT_DOWN_ROLE,
				Details: "could not update to lower role",
			}
		}
	}

	return nil
}

func (Account) TableName() string {
	return "account"
}

type VerifyAccount struct {
	Username   string `gorm:"type:text;NOT NULL"`
	SecretCode string `gorm:"NOT NULL"`
	IsUsed     bool   `gorm:"NOT NULL,default:false"`
	CreatedAt  int64  `gorm:"NOT NULL;autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
}

func (VerifyAccount) TableName() string {
	return "verify_account"
}

func (va VerifyAccount) BuildUpdateFields() map[string]interface{} {
	values := reflect.ValueOf(va)
	result := make(map[string]interface{}, values.NumField())

	for i := 0; i < values.NumField(); i++ {
		filed := values.Field(i)
		fieldName := values.Type().Field(i).Name

		if !filed.IsZero() {
			result[fieldName] = filed.Interface()
		}
	}

	return result
}

func (a Account) BuildUpdateFields() map[string]interface{} {
	values := reflect.ValueOf(a)
	result := make(map[string]interface{}, values.NumField())

	for i := 0; i < values.NumField(); i++ {
		filed := values.Field(i)
		fieldName := values.Type().Field(i).Name

		if !filed.IsZero() {
			result[fieldName] = filed.Interface()
		}
	}

	return result
}

type Sessions struct {
	ID           uuid.UUID `gorm:"NOT NULL;primaryKey"`
	Username     string    `gorm:"type:text;NOT NULL"`
	RefreshToken string    `gorm:"type:text;NOT NULL"`
	ExpiresAt    int64     `gorm:"NOT NULL"`
	ClientIP     string    `gorm:"type:text;NOT NULL"`
	UserAgent    string    `gorm:"type:text;NOT NULL"`
	IsBLocked    bool      `gorm:"NOT NULL;default:false"`
	CreatedAt    int64     `gorm:"NOT NULL;autoCreateTime"`
	UpdatedAt    int64     `gorm:"autoUpdateTime"`
}

func (Sessions) TableName() string {
	return "sessions"
}
