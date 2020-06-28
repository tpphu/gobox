package app

import (
	"github.com/tpphu/gobox/err"
	"github.com/tpphu/gobox/service/http"
)

var HttpServiceNotYetInjectToApp err.Error = err.Message("Http Service Not Yet Injected To App")
var GinServiceNotYetInjectToApp err.Error = err.Message("Gin Service Not Yet Injected To App")

func (a *App) GetHTTPService() *http.Http {
	if a.httpService == nil {
		panic(HttpServiceNotYetInjectToApp)
	}
	return a.httpService
}

func (a *App) initHttpService() {
	if a.httpService != nil {
		a.httpService.Init()
	}
}

func (a *App) runHttpService() error {
	if a.httpService != nil {
		return a.httpService.Run()
	}
	return nil
}

