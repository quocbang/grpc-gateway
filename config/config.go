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

type PostgresConfig struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type DatabaseGroup struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Redis    RedisConfig    `yaml:"redis"`
}

type SMTPConfig struct {
	SmtpServer  string `yaml:"smtp_server"`
	SmtpPort    int    `yaml:"smtp_port"`
	SenderEmail string `yaml:"sender_email"`
	Password    string `yaml:"password"`
}

type SenderGroup struct {
	SMTP SMTPConfig
}

type AuthConfig struct {
	SecretKeyPath        string `yaml:"secret_key_path"`
	AccessTokenLifeTime  string `yaml:"access_token_life_time"`
	RefreshTokenLifeTime string `yaml:"refresh_token_life_time"`
}

type ServerConfig struct {
	Auth   AuthConfig  `yaml:"auth"`
	Sender SenderGroup `yaml:"sender"`
}

type Config struct {
	DevMode  bool          `yaml:"dev_mode"`
	Database DatabaseGroup `yaml:"database"`
	Server   ServerConfig  `yaml:"server"`
}
