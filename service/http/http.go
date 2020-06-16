package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/service"
	"time"
)

var ErrServerClosed = errors.New("http: Server closed")

type Http struct {
	service.Runable
	*gin.Engine
	addr   string
	logger *logger.Logger
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
	engine := gin.New()
	http := &Http{}
	http.Engine = engine
	http.addr = ":3000"
	for _, opt := range opts {
		opt(http)
	}
	return http
}

func (s *Http) Init() {
	s.Use(logger.HttpLogger(s.logger), gin.Recovery())
	s.GET("/ping", func(c *gin.Context) {
		time.Sleep(10 *time.Second)
		c.String(200, "pong")
	})
}

func (s *Http) Run() error {
	return s.Engine.Run(s.addr)
}

func (s *Http) Shutdown(ctx context.Context) error {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("deadline:", deadline)

	}
	fmt.Println("deadline:", deadline)
	return nil
}


// add [[route, handler],[route, handler]] = group
// add [middleware,middleware,middleware...]
