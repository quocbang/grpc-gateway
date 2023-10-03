package models

import "github.com/quocbang/grpc-gateway/server/utils/roles"

type User struct {
	Username string      `gorm:"type:text;primaryKey"`
	Password []byte      `gorm:"type:bytea;NOT NULL"`
	Role     roles.Roles `gorm:"type:smallint;NOT NULL"`
}

func (User) TableName() string {
	return "user"
}
