package http

import (
	"context"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/service"
	http "net/http"
	"sync"
	"time"
)

var (
	ginMode     string
	ginNoLogger bool
	defaultPort = 3000
)

type GinService interface {
	service.Runable
	// block until ready
	Port() int
	isGinService()
}

type Config struct {
	isStarted    bool
	CertFile     string
	KeyFile      string
	IsGinDefault bool

	// used for register to SD
	// disable reg route/config
	NoRegRoute bool
}

type ginService struct {
	service.Runable
	name          string
	Addr 		 string
	Config
	logger        *logger.Logger
	svr           *myHttpServer
	*gin.Engine
	mtx            *sync.Mutex
	handlers      []func(*gin.Engine)
}

func NewGinService(name string) *ginService {
	// prepare router
	return &ginService{
		name:     name,
		logger: logger.New(),
		mtx:       &sync.Mutex{},
		handlers: []func(*gin.Engine){},
	}
}

func (gs *ginService) InitFlags() {
	prefix := "gin"
	flag.StringVar(&gs.Addr, prefix+"addr", ":3000", "gin server bind address")
}

func (gs *ginService) Init() error {
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	gs.logger.Debug("init gin engine...")
	gs.Engine = gin.New()
	gs.Use(logger.HttpLogger(gs.logger), gin.Recovery())
	gs.GET("/ping", func(c *gin.Context) {
		time.Sleep(10 *time.Second)
		c.String(200, "pong")
	})
	gs.svr = &myHttpServer{
		Server: &http.Server{Handler: gs.Engine, Addr: gs.Addr},
	}
	return nil
}

func (gs *ginService) Run() (err error) {
	if !gs.isStarted() {
		return nil
	}
	gs.logger.Info("Init Flags")
	gs.InitFlags()
	if err := gs.Init(); err != nil {
		return err
	}
	gs.logger.Printf("Starting http server: %v", gs.svr.Addr)
	for _, hdl := range gs.handlers {
		hdl(gs.Engine)
	}
	go func() {
		if err = gs.svr.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				gs.logger.Infof("Server closed under request: %v", err)
			} else {
				gs.logger.Fatalf("Server closed unexpect: %v", err)
			}
		}
		// in case of closed normally
		gs.Config.isStarted = true
	}()
	return err
}

func (gs *ginService) isStarted() bool {
	return !gs.Config.isStarted
}

func (gs *ginService) Shutdown(ctx context.Context) (err error) {
	gs.mtx.Lock()
	defer gs.mtx.Unlock()

	if !gs.isStarted() || gs.svr == nil {
		return errors.New("Server is not started")
	}

	stop := make(chan bool)
	go func() {
		// dummy preprocess before interrupted
		//time.Sleep(4 * time.Second)

		// Close immediately closes all active net.Listeners and any
		// connections in state StateNew, StateActive, or StateIdle. For a
		// graceful shutdown, use Shutdown.
		//
		// Close does not attempt to close (and does not even know about)
		// any hijacked connections, such as WebSockets.
		//
		// Close returns any error returned from closing the Server's
		// underlying Listener(s).
		//err = m.server.Close()
		// We can use .Shutdown to gracefully shuts down the server without
		// interrupting any active connection
		err = gs.svr.Shutdown(ctx)
		stop <- true
	}()

	select {
	case <-ctx.Done():
		gs.logger.Errorf("Timeout: %v", ctx.Err())
		break
	case <-stop:
		gs.logger.Infof("Finished")
	}
	return
}

