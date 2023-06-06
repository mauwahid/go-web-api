package api

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mauwahid/go-web-api/internal/app"
	"github.com/mauwahid/go-web-api/internal/platform/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	cfg        *config.Config
	echo       *echo.Echo
	apiHandler *handlerApi
}

func StartHttp() {
	fmt.Println("==== Start Run HTPP Server ====")

	httpServer := newServer()
	port := fmt.Sprintf(":%v", httpServer.cfg.Http.Port)

	// start server
	go func() {
		if err := httpServer.echo.Start(port); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.echo.Shutdown(ctx); err != nil {
		httpServer.echo.Logger.Fatal(err)
	}
}

func newServer() *Server {
	container := app.GetContainer()
	cfg := config.GetConfig()
	e := setupEcho(cfg)
	r := &Server{
		echo: e,
		cfg:  container.Config,
	}
	r.setupHandler(container)
	r.setupApiRoute()
	return r
}

func (r *Server) setupHandler(container *app.Container) {
	handler := newHandler(container)
	r.apiHandler = handler
}

func (r *Server) setupApiRoute() {

	{
		r.echo.GET("/health", echoPing)
	}

	groupUser := r.echo.Group("/v1/address")
	{
		groupUser.POST("/", r.apiHandler.Address.Create)
	}

}

func echoPing(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func setupEcho(cfg *config.Config) *echo.Echo {
	e := echo.New()

	e.Static("/assets", "web/assets")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("data table"))))

	/*jwtConfig := middleware.JWTConfig{
		Skipper:    skipJWT,
		SigningKey: cfg.Security.SigningKey,
		ContextKey: security.JWTEchoContext,
		Claims:     security.AppJWTClaims{},
	}
	e.Use(middleware.JWTWithConfig(jwtConfig))*/
	e.HideBanner = true
	return e
}
