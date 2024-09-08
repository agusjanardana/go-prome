package v1

import (
	notecontroller "github.com/go-prome/app/controllers/note_controller"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	NoteController notecontroller.NoteController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {

	apiV1 := e.Group("/api/v1")

	apiV1.POST("/notes", c1.NoteController.CreateNote)
	apiV1.GET("/notes", c1.NoteController.FindAllNotes)
	apiV1.GET("/notes/:id", c1.NoteController.FindNoteByID)
	apiV1.POST("/notes/:id", c1.NoteController.UpdateNoteByID)
	apiV1.DELETE("/notes/:id", c1.NoteController.DeleteNoteByID)
}
