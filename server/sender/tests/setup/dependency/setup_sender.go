package dependency

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"

	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/sender/connection"
)

var senderConfig struct {
	SmtpServer  string `long:"smtp-server" description:"the smtp server" env:"SMTP_SERVER_TEST"`
	SmtpPort    int    `long:"smtp-port" description:"the smtp port" env:"SMTP_PORT_TEST"`
	SenderEmail string `long:"smtp-sender" description:"sender email" env:"SMTP_SENDER_TEST"`
	Password    string `long:"smtp-sender-password" description:"sender's password" env:"SMTP_PASSWORD_TEST"`
}

func parseFlags() error {
	configuration := []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "smtp configuration",
			LongDescription:  "smtp configuration",
			Options:          &senderConfig,
		},
	}

	parse := flags.NewParser(nil, flags.IgnoreUnknown)
	for _, opt := range configuration {
		if _, err := parse.AddGroup(opt.LongDescription, opt.LongDescription, opt.Options); err != nil {
			return err
		}
	}

	if _, err := parse.Parse(); err != nil {
		return fmt.Errorf("failed to parse postgres flags")
	}

	return nil
}

func InitSenderTest() (sender.Sender, error) {
	err := parseFlags()
	if err != nil {
		return nil, err
	}

	sender, err := connection.NewEmailSender(connection.EmailSenderConfig{
		SmtpServer:  senderConfig.SmtpServer,
		SmtpPort:    senderConfig.SmtpPort,
		SenderEmail: senderConfig.SenderEmail,
		Password:    senderConfig.Password,
	})
	if err != nil {
		return nil, err
	}

	return sender, nil
}
