package handler

type Contact struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	PhoneNumber *string `json:"phoneNumber"`
	Email       *string `json:"email"`
}

func (c Contact) IsValid() bool {
	if c.FirstName == nil || *c.FirstName == "" {
		return false
	}
	return true
}
