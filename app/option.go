package app

import (
	"github.com/tpphu/gobox/service/http"
)

type Option func(a *App)

func Name(name string) Option {
	return func(a *App) {
		a.Name = name
	}
}

func Description(desc string) Option {
	return func(a *App) {
		a.Description = desc
	}
}

func WithHTTPService(address string) Option {
	return func(a *App) {
		a.httpService = http.Default(
			http.Address(address),
			http.Logger(a.Log))
	}
}
