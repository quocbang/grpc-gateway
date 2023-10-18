package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
)

// LoginRequest definition.
type LoginRequest struct {
	Username string
}

// LoginReply definition.
type LoginReply struct {
	Username     string
	HashPassword []byte
	Roles        roles.Roles
}

// CreateAccountRequest definition.
type CreateAccountRequest struct {
	Username     string
	Email        string
	HashPassword []byte
}

// GetAccountRequest definition.
type GetAccountRequest struct {
	Username string
}

// GetAccountReply definition.
type GetAccountReply struct {
	models.Account
}

// CreateVerifyAccountRequest definition.
type CreateVerifyAccountRequest struct {
	Username   string
	SecretCode string
}

// GetVerifyAccountRequest definition.
type GetUnVerifyAccountRequest struct {
	Username string
}

// GetVerifyAccountReply definition.
type GetUnVerifyAccountReply struct {
	models.VerifyAccount
}

// UpdateUserRoleRequest definition.
type UpdateUserRoleRequest struct {
	Username string
	ToRole   roles.Roles
}

type UpdateVerifiedAccountRequest struct {
	Username string
}

type CreateSessionsRequest struct {
	SSID         uuid.UUID
	Username     string
	RefreshToken string
	ExpiresAt    time.Duration
	ClientIP     string
	UserAgent    string
	IsBLocked    bool
}

type GetSessionsRequest struct {
	SessionID uuid.UUID
}

type GetSessionsReply struct {
	models.Sessions
}
