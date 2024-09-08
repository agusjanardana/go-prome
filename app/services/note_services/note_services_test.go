package noteservices

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-prome/app/drivers/models"
	"github.com/go-prome/app/drivers/repository/note_repository/mocks"
	"github.com/go-prome/app/services/note_services/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	noterepositoryTest mocks.NoteRepository
)

func setup() NoteServices {

	newNoteService := NewNoteServices(&noterepositoryTest)

	return newNoteService
}

func TestCreateNote(t *testing.T) {
	noteService := setup()

	t.Run("Create Note Success", func(t *testing.T) {
		note := domain.Note{
			Title: "Test",
			Body:  "Test",
		}

		noterepositoryTest.On("Create", mock.Anything).Return(models.Note{
			ID:    1,
			Title: "Test",
			Body:  "Test",
		}, nil).Once()

		result, _, err := noteService.CreateNote(note)
		assert.Nil(t, err)
		assert.Equal(t, "Test", result.Title)
		assert.Equal(t, "Test", result.Body)
	})

	t.Run("Create Note Bad Request", func(t *testing.T) {

		// Call the method
		_, statusCode, _ := noteService.CreateNote(domain.Note{})

		// Assert that the status code is http.StatusBadRequest
		assert.Equal(t, http.StatusBadRequest, statusCode)
	})

	t.Run("Create Note Internal Server Error", func(t *testing.T) {
		note := domain.Note{
			Title: "Test",
		}

		noterepositoryTest.On("Create", mock.Anything).Return(models.Note{}, errors.New("DB connection failed")).Once()
		_, statusCode, err := noteService.CreateNote(note)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.NotNil(t, err)
	})

}
