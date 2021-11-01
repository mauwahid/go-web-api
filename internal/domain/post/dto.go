package post

import "github.com/mauwahid/go-web-api/internal/platform/errors"

type Request struct {
	ID          int64
	Title       string
	MetaTitle   string
	Slug        string
	Summary     string
	IsPublished bool
	Content     string
	UserID      int64
}

func (r *Request) CreateValidation() error {
	if r.Title == "" {
		return errors.InvalidRequest
	}

	return nil
}

type Response struct {
}
