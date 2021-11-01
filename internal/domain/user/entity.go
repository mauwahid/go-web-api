package user

import "time"

type Entity struct {
	ID          int64
	FistName    string
	LastName    string
	Username    string
	Password    string
	Type        int64
	CreatedDate time.Time
	UpdatedDate time.Time
	CreatedBy   string
	UpdatedBy   string
}
