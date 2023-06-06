package contact

import "context"

type Repository interface {
	SaveAddress(ctx context.Context, contact Contact) (err error)
	GetAddressById(ctx context.Context, id int64) (contact Contact, err error)
	FindAllAddress(ctx context.Context, filter map[string]string, page int, limit int)
}
