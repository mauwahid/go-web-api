package http

import (
	"encoding/json"
	"fmt"
	"github.com/mauwahid/go-web-api/internal/domain/post"
	"net/http"
)

func (s *Server) postCreate(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&post.Request{}); err != nil {

	}

	fmt.Println("")
}
