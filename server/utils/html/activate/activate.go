package activate

import (
	"fmt"

	"github.com/matcornic/hermes/v2"

	"github.com/quocbang/grpc-gateway/server/utils/html"
)

type emailActivate struct {
	id         string
	secretCode string
}

func NewHTMLActivateService(id string, secretCode string) html.HTML {
	return emailActivate{
		id:         id,
		secretCode: secretCode,
	}
}

func (ea emailActivate) GenerateHTML() (string, error) {
	// check validate
	if err := ea.Validate(); err != nil {
		return "", err
	}

	h := html.NewHTMLHermes()
	email := hermes.Email{
		Body: hermes.Body{
			Name: ea.id,
			Intros: []string{
				"Welcome to grpc gateway designed by quocbang",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with grpc gateway, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  fmt.Sprintf("https://quocbangdev.com.vn/user/activate?id=%s?secret_code=%s", ea.id, ea.secretCode),
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return "", fmt.Errorf("failed to generate html, error: %v", err)
	}

	return emailBody, nil
}

func (ea emailActivate) Validate() error {
	if ea.id == "" {
		return fmt.Errorf("missing id")
	}
	if ea.secretCode == "" {
		return fmt.Errorf("missing secret code")
	}
	return nil
}
