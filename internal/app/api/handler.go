package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mauwahid/go-web-api/internal/app"
	"github.com/mauwahid/go-web-api/internal/app/handler"
)

type handlerApi struct {
	Address handlerAddress
}

type handlerAddress interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func newHandler(container *app.Container) *handlerApi {
	address := handler.NewAddressHandler(container.Service.Address)
	return &handlerApi{Address: address}
}