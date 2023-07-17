package models

import "time"

type Note struct {
	ID   int64  `bun:"id,pk,autoincrement"`
	Note string `bun:"note,notnull"`

	User   *User `bun:"rel:belongs-to,join:user_id=id"`
	UserId int64 `bun:",notnull"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
