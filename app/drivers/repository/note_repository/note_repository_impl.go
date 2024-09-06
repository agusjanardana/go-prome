package noterepository

import "github.com/go-prome/app/drivers/models"

type NoteRepositoryImpl struct {
}

// Create implements NoteRepository.
func (n *NoteRepositoryImpl) Create(note models.Note) (models.Note, error) {
	panic("unimplemented")
}

// DeleteByID implements NoteRepository.
func (n *NoteRepositoryImpl) DeleteByID(id int) (int64, error) {
	panic("unimplemented")
}

// FindAll implements NoteRepository.
func (n *NoteRepositoryImpl) FindAll() ([]models.Note, error) {
	panic("unimplemented")
}

// FindByID implements NoteRepository.
func (n *NoteRepositoryImpl) FindByID(id int) (models.Note, error) {
	panic("unimplemented")
}

// UpdateByID implements NoteRepository.
func (n *NoteRepositoryImpl) UpdateByID(id int, note models.Note) (models.Note, error) {
	panic("unimplemented")
}

func NewNoteRepositoryImpl() NoteRepository {
	return &NoteRepositoryImpl{}
}
