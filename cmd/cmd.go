package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/quocbang/grpc-gateway/config"
	"github.com/quocbang/grpc-gateway/pkg/grpc"
	"github.com/quocbang/grpc-gateway/pkg/rest"
	"github.com/quocbang/grpc-gateway/server/utils/times"
)

// parseConfig is parse config data in server/config.yaml file.
func parseConfig(configPath string) (*config.Config, error) {
	var cfs config.Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfs); err != nil {
		return nil, err
	}

	return &cfs, nil
}

func parseSecretKey(secretKeyPath string) (string, error) {
	data, err := os.ReadFile(secretKeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read secret key file path, error: %v", err)
	}

	return string(data), nil
}

func RunCmd() {
	flag.StringVar(&config.Fs.GatewayHost, "gateway-host", "localhost", "gateway host")
	flag.IntVar(&config.Fs.GatewayPort, "gateway-port", 9091, "gateway port")
	flag.StringVar(&config.Fs.GrpcHost, "grpc-host", "localhost", "grpc host")
	flag.IntVar(&config.Fs.GrpcPort, "grpc-port", 9092, "grpc port")
	flag.StringVar(&config.Fs.ConfigPath, "config-path", "", "config path")
	flag.StringVar(&config.Fs.TLS.Cert, "tls-cert", "", "tls cert")
	flag.StringVar(&config.Fs.TLS.Key, "tls-key", "", "tls key")
	flag.Parse()

	// init logger
	if err := registerLogger(false); err != nil {
		log.Fatalf("failed to register logger, error: %v", err)
	}

	// parse config
	configurations, err := parseConfig(config.Fs.ConfigPath)
	if err != nil {
		log.Fatalf("failed to parse config file, error: %v", err)
	}

	// parse secret key
	secretKey, err := parseSecretKey(configurations.Server.Auth.SecretKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	accessTokenLifeTime, err := times.StringToDuration(configurations.Server.Auth.AccessTokenLifeTime)
	if err != nil {
		log.Fatalf("failed to parse access token life time, error: %v", err)
	}

	refreshTokenLifeTime, err := times.StringToDuration(configurations.Server.Auth.RefreshTokenLifeTime)
	if err != nil {
		log.Fatalf("failed to parse refresh token life time, error: %v", err)
	}

	grpcOpts := grpc.GrpcOption{
		GrpcHost: config.Fs.GrpcHost,
		GrpcPort: config.Fs.GrpcPort,
		Database: config.DatabaseGroup{
			Postgres: config.PostgresConfig{
				Address:  configurations.Database.Postgres.Address,
				Port:     configurations.Database.Postgres.Port,
				Name:     configurations.Database.Postgres.Name,
				UserName: configurations.Database.Postgres.UserName,
				Password: configurations.Database.Postgres.Password,
				Schema:   configurations.Database.Postgres.Schema,
			},
			Redis: config.RedisConfig{
				Address:  configurations.Database.Redis.Address,
				Password: configurations.Database.Redis.Password,
			},
		},
		SecretKey:            secretKey,
		AccessTokenLifeTime:  accessTokenLifeTime,
		RefreshTokenLifeTime: refreshTokenLifeTime,
		SenderConfig:         configurations.Server.Sender.SMTP,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		if err := grpcOpts.Run(); err != nil {
			log.Fatal(err)
		}
		zap.L().Info("gRPC server stopped")
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
		zap.L().Info("gRPC gateway stopped")
	}()
	wg.Wait()
}

func registerLogger(isDevMod bool) error {
	logger := &zap.Logger{}
	var err error
	if isDevMod {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			return err
		}
	}

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
	return nil
}
