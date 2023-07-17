package database

import "github.com/procode2/accunotes/models"

type Storer interface {
	Init()

	CreateNewUser(user *models.User) (*models.User, error)

	GetUserById(userId string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	GetAllNotesForUser(userId uint32) ([]*models.NoteView, error)
	CreateNewNote(note *models.Note) (*models.Note, error)
	DeleteNoteById(noteId uint32) error
}
