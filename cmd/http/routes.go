package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", s.healthCheck)
	router.HandlerFunc(http.MethodPost, "/post/create", s.postCreate)

	router.
	return router
}
