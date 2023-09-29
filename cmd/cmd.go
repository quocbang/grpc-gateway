package cmd

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/quocbang/grpc-gateway/config"
	"github.com/quocbang/grpc-gateway/pkg/grpc"
	"github.com/quocbang/grpc-gateway/pkg/rest"
)

func RunCmd() {
	flag.StringVar(&config.Fs.GatewayHost, "gateway-host", "localhost", "gateway host")
	flag.IntVar(&config.Fs.GatewayPort, "gateway-port", 9091, "gateway port")
	flag.StringVar(&config.Fs.GrpcHost, "grpc-host", "localhost", "grpc host")
	flag.IntVar(&config.Fs.GrpcPort, "grpc-port", 9092, "grpc port")
	flag.StringVar(&config.Fs.ConfigPath, "config-path", "", "config path")
	flag.StringVar(&config.Fs.TLS.Cert, "tls-cert", "", "tls cert")
	flag.StringVar(&config.Fs.TLS.Key, "tls-key", "", "tls key")
	flag.Parse()

	// TODO: register logger

	// TODO: parse config file

	grpcOpts := grpc.GrpcOption{
		GrpcHost: config.Fs.GrpcHost,
		GrpcPort: config.Fs.GrpcPort,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		if err := grpcOpts.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	restOpts := rest.RestOption{
		GrpcServerEndpoint: fmt.Sprintf("%s:%d", config.Fs.GrpcHost, config.Fs.GrpcPort),
		GatewayHost:        config.Fs.GatewayHost,
		GatewayPort:        config.Fs.GatewayPort,
		Tls:                config.Fs.TLS,
	}
	go func() {
		if err := restOpts.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
}
