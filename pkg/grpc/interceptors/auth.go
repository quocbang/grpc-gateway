package interceptors

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"google.golang.org/grpc"
)

func Authorization(ctx context.Context) (context.Context, error) {
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

	// TODO: validate token here
	log.Println(token)

	// WARNING: In production define your own type to avoid context collisions.
	return ctx, nil
}
