package handler

import (
	"context"
	"github.com/mauwahid/go-web-api/internal/domain/contact"
)

type ContactService interface {
	Create(c context.Context, entity contact.Contact) error
}
