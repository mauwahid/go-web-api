package contact

import "context"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(c context.Context, entity Contact) (err error) {
	return nil
}
