package contact

type Contact struct {
	ID          int64
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}

func (c *Contact) Validate() bool {
	if c.FirstName == "" {
		return false
	}
	return true
}
