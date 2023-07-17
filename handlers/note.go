package handlers

import (
	"fmt"
	"net/http"

	"github.com/procode2/accunotes/database"
	"github.com/procode2/accunotes/models"
	"github.com/procode2/accunotes/utils"
)

type GetNoteRequest struct {
	Sid string `json:"sid"`
}

type PostNoteRequest struct {
	Sid  string `json:"sid"`
	Note string `json:"note"`
}

type NoteRequest struct {
	Sid  string `json:"sid"`
	Note string `json:"note"`
	Id   uint32 `json:"id"`
}

func HandleGetNotes(w http.ResponseWriter, r *http.Request) error {
	userId, _ := r.Context().Value("userId").(uint32)
	fmt.Printf("%+v\n", r.Body)

	notes, err := database.Db.GetAllNotesForUser(userId)
	if err != nil {
		utils.WriteJson(w, []*models.Note{})
		return nil
	}

	return utils.WriteJson(w, notes)
}

func HandlePostNote(w http.ResponseWriter, r *http.Request) error {
	userId, _ := r.Context().Value("userId").(uint32)
	body := &PostNoteRequest{}
	if err := utils.BodyAsJson(r.Body, body); err != nil {
		http.Error(w, "bad Request", http.StatusBadRequest)
		return err
	}

	note := &models.Note{
		Note:   body.Note,
		UserId: userId,
	}
	note, err := database.Db.CreateNewNote(note)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}

	return utils.WriteJson(w, map[string]uint32{"id": note.ID})
}

func HandleDeleteNote(w http.ResponseWriter, r *http.Request) error {
	body := &NoteRequest{}
	if err := utils.BodyAsJson(r.Body, body); err != nil {
		http.Error(w, "bad Request", http.StatusBadRequest)
		return err
	}

	err := database.Db.DeleteNoteById(body.Id)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
