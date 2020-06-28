package http

import (
	"github.com/tpphu/gobox/logger"
	"net/http"
	"sync"
)

type myHttpServer struct {
	addr   string
	Port int
	*http.Server
	logger *logger.Logger
	mtx       *sync.Mutex
}

func NewHttpSrv(port int) *myHttpServer {
	return &myHttpServer{
		Port:      port,
		mtx:       &sync.Mutex{},
	}
}

