package post

import _ "github.com/lib/pq"

type Repo struct {
}

func (r *Repo) FindAllPost() (posts []Entity, err error) {
	return
}
