package noterepository

import (
	"github.com/go-prome/app/drivers/models"
	"github.com/go-prome/config/db"
)

type NoteRepositoryImpl struct {
	dbCon db.Client
}

// Create implements NoteRepository.
func (n *NoteRepositoryImpl) Create(note models.Note) (models.Note, error) {
	err := n.dbCon.Conn().Create(&note).Error
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

// DeleteByID implements NoteRepository.
func (n *NoteRepositoryImpl) DeleteByID(id int) (int64, error) {
	err := n.dbCon.Conn().Delete(&models.Note{}, id).Error
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

// FindAll implements NoteRepository.
func (n *NoteRepositoryImpl) FindAll() ([]models.Note, error) {
	var notes []models.Note
	err := n.dbCon.Conn().Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

// FindByID implements NoteRepository.
func (n *NoteRepositoryImpl) FindByID(id int) (models.Note, error) {
	var note models.Note
	err := n.dbCon.Conn().First(&note, id).Error
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

// UpdateByID implements NoteRepository.
func (n *NoteRepositoryImpl) UpdateByID(id int, note models.Note) (models.Note, error) {
	err := n.dbCon.Conn().Model(&models.Note{}).Where("id = ?", id).Updates(&note).Error
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func NewNoteRepositoryImpl(dbClient db.Client) NoteRepository {
	return &NoteRepositoryImpl{
		dbCon: dbClient,
	}
}
