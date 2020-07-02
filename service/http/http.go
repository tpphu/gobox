package http

import (
	"errors"
	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/service"
)

var ErrServerClosed = errors.New("http: Server closed")

type Http struct {
	service.Runable
	*ginService
	addr   string
	logger *logger.Logger
	//*http.Server
}

type Option func(a *Http)

func Address(addr string) Option {
	return func(a *Http) {
		a.addr = addr
	}
}

func Logger(logger *logger.Logger) Option {
	return func(a *Http) {
		a.logger = logger
	}
}

func Default(opts ...Option) *Http {
	myHttp := &Http{}
	myHttp.ginService = NewGinService(":3000")
	for _, opt := range opts {
		opt(myHttp)
	}
	return myHttp
}

func (s *Http) Init() {
	if s.ginService != nil  {
		s.ginService.Init()
	}
}

func (s *Http) Run() error {
	if s.ginService != nil  {
		return s.ginService.Run()
	}
	return nil
}

func (s *Http) Shutdown() error {
	if s.ginService != nil {
		return s.ginService.Shutdown()
	}
	return errors.New("no no ")
}

// add [[route, handler],[route, handler]] = group
// add [middleware,middleware,middleware...]
