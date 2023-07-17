package postgres

import (
	"context"
	"fmt"

	"github.com/procode2/accunotes/models"
	"github.com/uptrace/bun"
)

func (p *PostgresStore) GetAllNotesForUser(userId uint32) ([]*models.NoteView, error) {
	notes := make([]*models.NoteView, 0)
	ctx := context.Background()
	err := p.db.NewSelect().Model((*models.Note)(nil)).Column("id", "note").Where("? = ?", bun.Ident("user_id"), userId).Scan(ctx, &notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (p *PostgresStore) CreateNewNote(note *models.Note) (*models.Note, error) {
	ctx := context.Background()

	_, err := p.db.NewInsert().Model(note).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return note, nil
}

func (p *PostgresStore) DeleteNoteById(noteId uint32) error {
	ctx := context.Background()

	_, err := p.db.NewDelete().Model((*models.Note)(nil)).Where("id = ?", noteId).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
