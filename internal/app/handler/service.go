package handler

import (
	"context"
	"github.com/mauwahid/go-web-api/internal/domain/contact"
)

type AddressService interface {
	Create(c context.Context, entity contact.Contact) error
}
