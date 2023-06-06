package handler

import "github.com/labstack/echo/v4"

type Address struct {
	addressService AddressService
}

func NewAddressHandler(service AddressService) *Address{
	return &Address{addressService:service}
}

func (h *Address) Create(c echo.Context) error {
	return nil
}

func (h *Address)  Read(c echo.Context) error {
	return nil
}

func (h *Address)  Update(c echo.Context) error {
	return nil
}

func (h *Address)  Delete(c echo.Context) error {
	return nil
}

