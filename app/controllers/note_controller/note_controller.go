package notecontroller

import "github.com/labstack/echo/v4"

type NoteController interface {
	CreateNote(c echo.Context) error
	FindAllNotes(c echo.Context) error
	FindNoteByID(c echo.Context) error
	UpdateNoteByID(c echo.Context) error
	DeleteNoteByID(c echo.Context) error
}
