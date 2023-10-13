package interceptors

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	"github.com/quocbang/grpc-gateway/server/utils/roles"
	t "github.com/quocbang/grpc-gateway/server/utils/token"
	"google.golang.org/grpc"
)

type Auth struct {
	SecretKey string
}

func (a Auth) Authorization(ctx context.Context) (context.Context, error) {
	fullMethod, ok := grpc.Method(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract method in authorization step")
	}

	if !roles.IsNeedToCheckMethodAccess(fullMethod) {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	jwt := t.JWT{
		SecretKey: a.SecretKey,
	}
	claims, err := jwt.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	// set claims to context
	ctx = setContextWithJWTClaims(ctx, claims)

	return ctx, nil
}
