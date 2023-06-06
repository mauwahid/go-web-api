package app

import (
	"github.com/mauwahid/go-web-api/internal/app/handler"
	"github.com/mauwahid/go-web-api/internal/domain/contact"
	"github.com/mauwahid/go-web-api/internal/platform/config"
)

type Container struct {
	Config  *config.Config
	Service Service
}

type Service struct {
	Address handler.AddressService
}

type Repository struct {
	Address contact.Repository
}

func GetContainer() *Container {
	config := config.GetConfig()

	return &Container{
		Config: config,
	}
}
