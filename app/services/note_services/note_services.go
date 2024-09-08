package noteservices

import "github.com/go-prome/app/services/note_services/domain"

type NoteServices interface {
	CreateNote(note domain.Note) (domain.Note, int, error)
	FindAllNotes() ([]domain.Note, int, error)
	FindNoteByID(id int) (domain.Note, int, error)
	UpdateNoteByID(id int, note domain.Note) (domain.Note, int, error)
	DeleteNoteByID(id int) (int, error)
}
