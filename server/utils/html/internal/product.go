package internal

import "github.com/matcornic/hermes/v2"

// NewHermesProduct appears in header & footer of e-mails
func NewHermesProduct() hermes.Product {
	return hermes.Product{
		Name:      "Quoc Bang",
		Link:      "https://quocbangdev.com.vn/",
		Copyright: "Copyright © 2023 quocbang Initiative. All rights reserved.",
	}
}
