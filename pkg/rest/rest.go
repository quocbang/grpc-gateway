package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quocbang/grpc-gateway/config"
	"github.com/quocbang/grpc-gateway/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestOption struct {
	GrpcServerEndpoint string
	GatewayHost        string
	GatewayPort        int
	Tls                config.TlsConfig
}

func (s RestOption) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := server.NewRegisterHandler(ctx, mux, s.GrpcServerEndpoint, opts); err != nil {
		return err
	}
	address := fmt.Sprintf("%s:%d", s.GatewayHost, s.GatewayPort)
	log.Printf("starting grpc gateway with host: %s, port: %d \n", s.GatewayHost, s.GatewayPort)

	if s.Tls.IsUseTLS() {
		log.Println("start listen grpc gateway with TLS")
		return http.ListenAndServeTLS(address, s.Tls.Cert, s.Tls.Key, mux)
	}

	log.Println("start listen grpc gateway without TLS")
	return http.ListenAndServe(address, mux)
}
