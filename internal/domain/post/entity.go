package post

import "time"

type Entity struct {
	ID            int64
	Title         string
	MetaTitle     string
	Slug          string
	Summary       string
	IsPublished   bool
	Content       string
	CreatedDate   time.Time
	UpdatedDate   time.Time
	CreatedBy     string
	UpdatedBy     string
	PublishedDate time.Time
	PublishedBy   string
	Author        int64
}
