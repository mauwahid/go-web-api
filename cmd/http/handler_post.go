package http

import (
	"encoding/json"
	"net/http"

	"github.com/mauwahid/go-web-api/internal/domain/post"
)

func (s *Server) postCreate(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&post.Request{}); err != nil {

	}

}
