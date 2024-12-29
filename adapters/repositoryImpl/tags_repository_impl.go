package repositoryimpl

import (
	"mhakheancode/api/model"
	"mhakheancode/api/ports/repository"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	DB *gorm.DB
}

// GetAll implements repository.TagsRepository.
func (t *TagsRepositoryImpl) GetAll() (*[]model.Tag, error) {
	var tags *[]model.Tag
	err := t.DB.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func NewTagsRepositoryImpl(db *gorm.DB) repository.TagsRepository {
	return &TagsRepositoryImpl{
		DB: db,
	}
}
