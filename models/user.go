package models

import "time"

type User struct {
	ID       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name,notnull"`
	Email    string `bun:"email,notnull"`
	Password string `bun:"password,notnull"`

	Notes []*Note `bun:"rel:has-many,join:id=user_id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
