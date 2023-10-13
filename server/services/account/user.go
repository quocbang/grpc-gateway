package account

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/quocbang/grpc-gateway/pkg/grpc/interceptors"
	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/errors"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/utils/hashing"
	"github.com/quocbang/grpc-gateway/server/utils/html/activate"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"github.com/quocbang/grpc-gateway/server/utils/token"
	"github.com/quocbang/grpc-gateway/server/utils/validator"
)

type server struct {
	repo                 repositories.Repositories
	sender               sender.Sender
	accessTokenLifeTime  time.Duration
	refreshTokenLifeTime time.Duration
	secretKey            string
	hasPermission        func(string, roles.Roles) bool
}

func NewAccount(
	repo repositories.Repositories,
	sender sender.Sender,
	accessTokenLifeTime time.Duration,
	refreshTokenLifeTime time.Duration,
	secretKey string,
	hasPermission func(string, roles.Roles) bool) pb.AccountServiceServer {
	return &server{
		repo:                 repo,
		sender:               sender,
		accessTokenLifeTime:  accessTokenLifeTime,
		refreshTokenLifeTime: refreshTokenLifeTime,
		secretKey:            secretKey,
		hasPermission:        hasPermission,
	}
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// check input data
	validateRules := map[string]string{
		"Username": "required",
		"Password": "required",
	}
	if err := validator.ValidateStructWithoutTag[pb.LoginRequest](req, validateRules); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.repo.Account().GetAccount(ctx, repositories.GetAccountRequest{
		Username: req.GetUsername(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// check password
	if ok, err := hashing.IsMatchedPassword(req.GetPassword(), user.Password); !ok && err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	// generate access token
	jwt := token.JWT{
		SecretKey: s.secretKey,
		User: token.UserInfo{
			Username: user.Username,
			Role:     user.Role,
		},
		TokenLifeTime: s.accessTokenLifeTime,
	}
	accessToken, err := jwt.GenerateToken()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// generate refresh token
	jwt.TokenLifeTime = s.refreshTokenLifeTime
	refreshToken, err := jwt.GenerateToken()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.LoginResponse{
		SessionId:             uuid.NewString(),
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  timestamppb.Now(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: timestamppb.Now(),
	}, nil
}

func (s *server) RenewAccess(ctx context.Context, req *pb.RenewAccessRequest) (*pb.RenewAccessResponse, error) {
	claims := interceptors.GetJWTClaimsFromContext(ctx)
	fmt.Println(claims)
	return &pb.RenewAccessResponse{}, nil
}

func (s *server) SignUp(ctx context.Context, req *pb.SignUpRequest) (resp *pb.SignUpResponse, err error) {
	// check input values
	validateRules := map[string]string{
		"Username": "required",
		"Email":    "required,email",
		"Password": "required,min=8",
	}
	if err := validator.ValidateStructWithoutTag[pb.SignUpRequest](req, validateRules); err != nil {
		return nil, errors.Error{
			Code:    errors.Code_MISSING_REQUEST,
			Details: err.Error(),
		}
	}

	// start transaction
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		if e, ok := errors.As(err); ok {
			if e.Code == errors.Code_ALREADY_IN_TRANSACTION {
				interceptors.GetLoggerFormContext(ctx).Warn("duplicated transaction")
			}
		}
		return nil, errors.Error{
			Details: fmt.Sprintf("failed to start transaction, error: %v", err),
		}
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// create user
	hashPassword, err := hashing.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Error{
			Details: fmt.Sprintf("failed to hash password, error: %v \n", err),
		}
	}
	if err := tx.Account().CreateAccount(ctx, repositories.CreateAccountRequest{
		Username:     req.Username,
		Email:        req.Email,
		HashPassword: hashPassword,
	}); err != nil {
		return nil, err
	}

	// get user to ensure the use was created and generate token.
	user, err := tx.Account().GetAccount(ctx, repositories.GetAccountRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	// generate JWT for access token.
	jwt := token.JWT{
		SecretKey: s.secretKey,
		User: token.UserInfo{
			Username: user.Username,
			Role:     user.Role,
		},
		TokenLifeTime: s.accessTokenLifeTime,
	}
	accessToken, err := jwt.GenerateToken()
	if err != nil {
		return nil, errors.Error{
			Details: err.Error(),
		}
	}

	// generate JWT for refresh token.
	jwt.TokenLifeTime = s.refreshTokenLifeTime
	refreshToken, err := jwt.GenerateToken()
	if err != nil {
		return nil, errors.Error{
			Details: err.Error(),
		}
	}

	// create verify account
	if err = tx.Account().CreateVerifyAccount(ctx, repositories.CreateVerifyAccountRequest{Username: req.Username, SecretCode: uuid.NewString()}); err != nil {
		return nil, fmt.Errorf("failed to create new account, error: %v", err)
	}

	// get verify account
	verifyAccountInfo, err := tx.Account().GetVerifyAccount(ctx, repositories.GetVerifyAccountRequest{
		Username: req.Username,
	})

	// all ok
	err = tx.Commit()
	if err != nil {
		return nil, errors.Error{
			Details: fmt.Sprintf("failed to commit, error: %v", err),
		}
	}

	// send verify mail
	htmlActivateService := activate.NewHTMLActivateService(req.Username, verifyAccountInfo.SecretCode)
	content, err := htmlActivateService.GenerateHTML()
	if err != nil {
		return nil, fmt.Errorf("failed to generate activate html, error: %v", err)
	}
	subject := "WellCome to the grpc gateway project design by quocbang"
	err = s.sender.Email().SendVerifyEmail(ctx, req.Email, subject, content)
	if err != nil {
		return nil, fmt.Errorf("failed to send mail, error: %v", err)
	}

	return &pb.SignUpResponse{
		SessionId:             uuid.NewString(),
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  timestamppb.Now(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: timestamppb.Now(),
	}, nil
}
