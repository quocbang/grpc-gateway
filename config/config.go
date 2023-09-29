package config

var Fs Flags

type Flags struct {
	GatewayHost string
	GatewayPort int
	GrpcHost    string
	GrpcPort    int
	ConfigPath  string

	TLS TlsConfig
}

type TlsConfig struct {
	Cert string
	Key  string
}

func (t TlsConfig) IsUseTLS() bool {
	return t.Cert != "" && t.Key != ""
}
