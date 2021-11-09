package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", s.healthCheck)
	router.HandlerFunc(http.MethodPost, "/post/create", s.postCreate)

}
