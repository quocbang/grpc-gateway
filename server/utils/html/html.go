package html

import (
	"github.com/matcornic/hermes/v2"

	"github.com/quocbang/grpc-gateway/server/utils/html/internal"
)

type HTML interface {
	GenerateHTML() (string, error)
}

func NewHTMLHermes() hermes.Hermes {
	product := internal.NewHermesProduct()
	return hermes.Hermes{
		Product: product,
	}
}
