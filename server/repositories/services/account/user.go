package account

import (
	"context"
	"errors"

	"github.com/quocbang/grpc-gateway/server/repositories"
	repositoriesErr "github.com/quocbang/grpc-gateway/server/repositories/errors"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"gorm.io/gorm"
)

type service struct {
	pg *gorm.DB
}

func NewAccount(pg *gorm.DB) repositories.Account {
	return service{
		pg: pg,
	}
}

func (s service) Login(ctx context.Context, req repositories.LoginRequest) (repositories.LoginReply, error) {
	user := models.Account{}
	err := s.pg.Where("username=?", req.Username).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repositories.LoginReply{}, repositoriesErr.Error{
				Code:    repositoriesErr.Code_ACCOUNT_DOES_NOT_EXIST,
				Details: "account does not exist",
			}
		}
		return repositories.LoginReply{}, err
	}
	return repositories.LoginReply{
		Username:     user.Username,
		HashPassword: user.Password,
		Roles:        user.Role,
	}, nil
}

func (s service) CreateAccount(ctx context.Context, req repositories.CreateAccountRequest) error {
	account := models.Account{
		Username: req.Username,
		Email:    req.Email,
		Password: req.HashPassword,
	}
	return s.pg.Create(&account).Error
}

func (s service) GetAccount(ctx context.Context, req repositories.GetAccountRequest) (repositories.GetAccountReply, error) {
	account := models.Account{}
	if err := s.pg.Where("username = ?", req.Username).Take(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repositories.GetAccountReply{}, repositoriesErr.ErrDataNotFound
		}
		return repositories.GetAccountReply{}, err
	}
	return repositories.GetAccountReply{Account: account}, nil
}

func (s service) CreateVerifyAccount(ctx context.Context, req repositories.CreateVerifyAccountRequest) error {
	verifyAccount := models.VerifyAccount{
		Username:   req.Username,
		SecretCode: req.SecretCode,
	}
	return s.pg.Create(&verifyAccount).Error
}

func (s service) GetUnVerifyAccount(ctx context.Context, req repositories.GetUnVerifyAccountRequest) (repositories.GetUnVerifyAccountReply, error) {
	verifyAccount := models.VerifyAccount{}
	if err := s.pg.Where("username = ? and is_used = false", req.Username).Take(&verifyAccount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repositories.GetUnVerifyAccountReply{}, repositoriesErr.ErrDataNotFound
		}
		return repositories.GetUnVerifyAccountReply{}, err
	}
	return repositories.GetUnVerifyAccountReply{VerifyAccount: verifyAccount}, nil
}

// UpdateUserRole is update user role
// *Note: please ensure the user role is always to the higher role.
func (s service) UpdateUserRole(ctx context.Context, req repositories.UpdateUserRoleRequest) (repositories.CommonUpdateReply, error) {
	// TODO: fix to update multiple fields
	updateConditions := models.Account{
		Username:       req.Username,
		IsUserVerified: true,
		Role:           req.ToRole,
	}.BuildUpdateFields()
	reply := s.pg.Model(&models.Account{Username: req.Username, Role: req.ToRole}).Where("username = ?", req.Username).Updates(updateConditions)
	return repositories.CommonUpdateReply{
		AffectedRows: repositories.AffectedRows(reply.RowsAffected),
	}, reply.Error
}

func (s service) UpdateVerifiedAccount(ctx context.Context, req repositories.UpdateVerifiedAccountRequest) (repositories.CommonUpdateReply, error) {
	// TODO: fix to update multiple fields
	updateConditions := models.VerifyAccount{
		Username: req.Username,
		IsUsed:   true,
	}.BuildUpdateFields()
	reply := s.pg.Model(&models.VerifyAccount{}).Where("username = ?", req.Username).Updates(updateConditions)
	return repositories.CommonUpdateReply{
		AffectedRows: repositories.AffectedRows(reply.RowsAffected),
	}, reply.Error
}
