package contact

import (
	"context"
	"github.com/mauwahid/go-web-api/internal/platform/errs"
	"log"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(c context.Context, entity Contact) (err error) {

	if !entity.Validate() {
		err = errs.ErrorInvalidReq
		return
	}

	if err = s.repo.SaveAddress(c, entity); err != nil {
		log.Println("error save db ", err.Error())
	}

	return
}
