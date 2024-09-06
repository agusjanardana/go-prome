package noterepository

import "github.com/go-prome/app/drivers/models"

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
	FindAll() ([]models.Note, error)
	FindByID(id int) (models.Note, error)
	UpdateByID(id int, note models.Note) (models.Note, error)
	DeleteByID(id int) (int64, error)
}
