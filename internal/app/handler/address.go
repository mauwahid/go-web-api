package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mauwahid/go-web-api/internal/domain/contact"
	"net/http"
)

type Address struct {
	contactService ContactService
}

func NewAddressHandler(contactService ContactService) *Address {
	return &Address{contactService: contactService}
}

func (h *Address) Create(c echo.Context) (err error) {

	var contactReq Contact
	if err = c.Bind(&contactReq); err != nil {
		return
	}

	if !contactReq.IsValid() {
		return
	}

	var ctact contact.Contact
	if contactReq.FirstName != nil {
		ctact.FirstName = *contactReq.FirstName
	}
	if contactReq.LastName != nil {
		ctact.LastName = *contactReq.LastName
	}
	if contactReq.Email != nil {
		ctact.Email = *contactReq.Email
	}
	if contactReq.PhoneNumber != nil {
		ctact.PhoneNumber = *contactReq.PhoneNumber
	}

	ctx := context.Background()
	c.Set("echo", c)
	err = h.contactService.Create(ctx, ctact)

	return c.JSON(http.StatusOK, RenderResponse(ctact, err))
}

func (h *Address) Read(c echo.Context) error {
	return nil
}

func (h *Address) Update(c echo.Context) error {
	return nil
}

func (h *Address) Delete(c echo.Context) error {
	return nil
}
