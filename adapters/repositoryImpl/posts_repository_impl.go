package repositoryimpl

import (
	"mhakheancode/api/model"
	"mhakheancode/api/ports/repository"

	"gorm.io/gorm"
)

type PostsRepositoryImpl struct {
	db *gorm.DB
}

func (p *PostsRepositoryImpl) Get(id int32) (*model.Post, error) {
	var post model.Post
	err := p.db.First(&post, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (p *PostsRepositoryImpl) Save(model *model.Post) error {
	err := p.db.Create(&model)
	if err != nil {
		return err.Error
	}
	return err.Error
}

func NewPostsRepositoryImpl(db *gorm.DB) repository.PostRepository {
	return &PostsRepositoryImpl{
		db: db,
	}
}
