package http

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/module/apmlogrus/v2"
)

const (
	TIMEOUT_SHUTDOWN time.Duration = 5 * time.Second
)

type HTTP struct {
	Router   *gin.Engine
	Listener *net.Listener
	Server   *http.Server
	Log      *logrus.Entry
}

func (c *HTTP) Run() {
	c.Log.Trace("Listen on ", os.Getenv("HTTP_ADDR"))

	if err := c.Server.Serve(*c.Listener); err != nil && err != http.ErrServerClosed {
		c.Log.Fatal("Server closed unexpect")
	}
}

func (c *HTTP) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SHUTDOWN)
	defer cancel()

	if err := c.Server.Shutdown(ctx); err != nil {
		c.Log.Error("Forced to shutdown: ", err)
	}
}

func New() *HTTP {
	log := logrus.WithFields(logrus.Fields{"module": "http"})
	log.Logger.AddHook(&apmlogrus.Hook{})

	listener, err := net.Listen("tcp", os.Getenv("HTTP_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(apmgin.Middleware(router))

	return &HTTP{
		Router:   router,
		Listener: &listener,
		Server: &http.Server{
			Handler:           router,
			ReadHeaderTimeout: 5 * time.Second,
		},
		Log: log,
	}
}