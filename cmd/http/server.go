package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mauwahid/go-web-api/internal/domain/post"
	"github.com/mauwahid/go-web-api/internal/platform/config"
)

type envelope map[string]interface{}

type Server struct {
	config config.Config
	logger logger.Logger
	post   post.Contract
}

func Initialize(config config.Config, logger logger.Logger) *http.Server {
	srv := &Server{}
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%d", config.Http.Port),
		Handler:      srv.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  config.Http.ReadTimeout,
		WriteTimeout: config.Http.WriteTimeout,
	}
	return httpServer
}
