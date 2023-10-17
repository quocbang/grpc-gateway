package distributor

type VerifyEmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}
