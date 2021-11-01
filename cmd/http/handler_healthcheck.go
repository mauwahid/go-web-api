package http

import (
	"fmt"
	"net/http"
)

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "service up")
}
