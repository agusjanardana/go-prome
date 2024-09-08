package notecontroller

import (
	noteservices "github.com/go-prome/app/services/note_services"
	"github.com/labstack/echo/v4"
)

type NoteControllerImpl struct {
	noteservices noteservices.NoteServices
}

// CreateNote implements NoteController.
func (n NoteControllerImpl) CreateNote(c echo.Context) error {
	panic("unimplemented")
}

// DeleteNoteByID implements NoteController.
func (n NoteControllerImpl) DeleteNoteByID(c echo.Context) error {
	panic("unimplemented")
}

// FindAllNotes implements NoteController.
func (n NoteControllerImpl) FindAllNotes(c echo.Context) error {
	data, httpCode, err := n.noteservices.FindAllNotes()
	if err != nil {
		return c.JSON(httpCode, err)
	}

	return c.JSON(httpCode, data)
}

// FindNoteByID implements NoteController.
func (n NoteControllerImpl) FindNoteByID(c echo.Context) error {
	panic("unimplemented")
}

// UpdateNoteByID implements NoteController.
func (n NoteControllerImpl) UpdateNoteByID(c echo.Context) error {
	panic("unimplemented")
}

func NewNoteControllerImpl(noteservices noteservices.NoteServices) NoteController {
	return NoteControllerImpl{noteservices: noteservices}
}
