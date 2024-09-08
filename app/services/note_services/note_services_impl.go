package noteservices

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-prome/app/drivers/models"
	noterepository "github.com/go-prome/app/drivers/repository/note_repository"
	"github.com/go-prome/app/services/note_services/domain"
	"github.com/jinzhu/copier"
)

type NoteServicesImpl struct {
	repository noterepository.NoteRepository
}

// CreateNote implements NoteServices.
func (n *NoteServicesImpl) CreateNote(note domain.Note) (domain.Note, int, error) {
	var modelNote models.Note
	err := copier.Copy(&modelNote, &note)
	if err != nil || reflect.ValueOf(note).IsZero() {
		return domain.Note{}, http.StatusBadRequest, err
	}

	data, err := n.repository.Create(modelNote)
	fmt.Println("data", data, "err", err)
	if err != nil {
		return domain.Note{}, http.StatusInternalServerError, err
	}

	var domainNote domain.Note
	err = copier.Copy(&domainNote, &data)
	if err != nil || reflect.ValueOf(domainNote).IsZero() {
		return domain.Note{}, http.StatusBadRequest, err
	}

	return domainNote, http.StatusCreated, nil
}

// DeleteNoteByID implements NoteServices.
func (n *NoteServicesImpl) DeleteNoteByID(id int) (int, error) {
	_, err := n.repository.DeleteByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// FindAllNotes implements NoteServices.
func (n *NoteServicesImpl) FindAllNotes() ([]domain.Note, int, error) {
	data, err := n.repository.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var domainNotes []domain.Note
	err = copier.Copy(&domainNotes, &data)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return domainNotes, http.StatusOK, nil
}

// FindNoteByID implements NoteServices.
func (n *NoteServicesImpl) FindNoteByID(id int) (domain.Note, int, error) {
	data, err := n.repository.FindByID(id)
	if err != nil {
		return domain.Note{}, http.StatusInternalServerError, err
	}

	var domainNote domain.Note
	err = copier.Copy(&domainNote, &data)
	if err != nil {
		return domain.Note{}, http.StatusBadRequest, err
	}

	return domainNote, http.StatusOK, nil
}

// UpdateNoteByID implements NoteServices.
func (n *NoteServicesImpl) UpdateNoteByID(id int, note domain.Note) (domain.Note, int, error) {
	var modelNote models.Note
	err := copier.Copy(&modelNote, &note)
	if err != nil {
		return domain.Note{}, http.StatusBadRequest, err
	}

	data, err := n.repository.UpdateByID(id, modelNote)
	if err != nil {
		return domain.Note{}, http.StatusInternalServerError, err
	}

	var domainNote domain.Note
	err = copier.Copy(&domainNote, &data)
	if err != nil {
		return domain.Note{}, http.StatusBadRequest, err
	}

	return domainNote, http.StatusOK, nil
}

func NewNoteServices(repository noterepository.NoteRepository) NoteServices {
	return &NoteServicesImpl{
		repository: repository,
	}
}
