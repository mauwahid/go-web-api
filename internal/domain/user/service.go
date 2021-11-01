package user

import "context"

type Contract interface {
	ValidateUser(c context.Context, userID int64) (isValid bool, err error)
}

type Service struct {
}
