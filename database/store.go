package database

import "github.com/procode2/accunotes/models"

type Storer interface {
	Init()

	CreateNewUser(user *models.User) (*models.User, error)

	GetUserById(userId string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	GetAllNotesForUser(search string) ([]*models.Note, error)
	CreateNewNote(note *models.Note) (*models.Note, error)
	UpdateNote(note *models.Note) error
	GetNoteById(noteId int64) (*models.Note, error)
	DeleteNoteById(noteId int64) error
}
