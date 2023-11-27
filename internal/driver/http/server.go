package http

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

const (
	TIMEOUT_SHUTDOWN time.Duration = 5 * time.Second
)

type HTTP struct {
	Router   *echo.Echo
	Listener net.Listener
	Server   *http.Server
	Log      *logrus.Entry
}

func (c *HTTP) Run() {
	c.Log.Trace("Listen on ", os.Getenv("HTTP_ADDR"))

	if err := c.Server.Serve(c.Listener); err != nil && err != http.ErrServerClosed {
		c.Log.Fatal("Server closed unexpectedly: ", err)
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

	listener, err := net.Listen("tcp", os.Getenv("HTTP_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(RateLimiter(rate.Limit(1), 5))
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &HTTP{
		Router:   e,
		Listener: listener,
		Server: &http.Server{
			Handler:           e,
			ReadHeaderTimeout: 5 * time.Second,
		},
		Log: log,
	}
}

func RateLimiter(r rate.Limit, b int) echo.MiddlewareFunc {
	limiter := rate.NewLimiter(r, b)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				return echo.NewHTTPError(http.StatusTooManyRequests, "Rate limit exceeded")
			}
			return next(c)
		}
	}
}
