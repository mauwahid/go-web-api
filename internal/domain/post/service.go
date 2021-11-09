package post

import (
	"context"

	"github.com/mauwahid/go-web-api/internal/domain/user"
)

type Contract interface {
	Create(c context.Context, request Request) (err error, response Response)
}

type Service struct {
	User user.Contract
	repo Repo
}

func NewService(user user.Service, sql datab)

func (s Service) Create(c context.Context, request Request) (err error) {
	if err = request.CreateValidation(); err != nil {
		return err
	}

	return nil
}
