package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/procode2/accunotes/models"
	"github.com/uptrace/bun"
)

func (p *PostgresStore) GetAllNotesForUser(search string) ([]*models.Note, error) {
	notes := make([]*models.Note, 0)
	ctx := context.Background()
	q := p.db.NewSelect().Model(&notes)
	if search != "" {
		q = q.Where("? ILIKE ?", bun.Ident("title"), "%"+search+"%")
	}
	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (p *PostgresStore) CreateNewNote(note *models.Note) (*models.Note, error) {
	ctx := context.Background()
	err := p.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

		_, err := tx.NewInsert().Model(note).Exec(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return err

	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return note, nil
}

func (p *PostgresStore) UpdateNote(note *models.Note) error {
	// TODO
	return nil
}

func (p *PostgresStore) GetNoteById(noteId int64) (*models.Note, error) {
	note := &models.Note{}

	ctx := context.Background()
	err := p.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

		err := tx.NewSelect().Model(note).Where("? = ?", bun.Ident("id"), noteId).Scan(ctx)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return note, nil
}

func (p *PostgresStore) DeleteNoteById(noteId int64) error {
	// TODO
	return nil
}
