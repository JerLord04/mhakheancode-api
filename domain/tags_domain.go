package domain

import (
	"mhakheancode/api/model"
	"mhakheancode/api/ports/repository"
	"mhakheancode/api/ports/services"
)

type TagsDomain struct {
	TagsRepository repository.TagsRepository
}

// GetAllTags implements services.TagsPort.
func (t *TagsDomain) GetAllTags() (*[]model.Tag, error) {
	data, err := t.TagsRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewTagsDomain(tagsRepository repository.TagsRepository) services.TagsPort {
	return &TagsDomain{
		TagsRepository: tagsRepository,
	}
}
