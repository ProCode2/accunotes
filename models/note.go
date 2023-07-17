package models

import "time"

type Note struct {
	ID   uint32 `bun:"id,pk,autoincrement"`
	Note string `bun:"note,notnull"`

	User   *User  `bun:"rel:belongs-to,join:user_id=id"`
	UserId uint32 `bun:",notnull"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type NoteView struct {
	Id   uint32 `json:"id"`
	Note string `json:"note"`
}
