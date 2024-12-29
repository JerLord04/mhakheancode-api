package repositoryimpl

import (
	"mhakheancode/api/model"
	"mhakheancode/api/ports/repository"

	"gorm.io/gorm"
)

type PostsRepositoryImpl struct {
	db *gorm.DB
}

func (p *PostsRepositoryImpl) GetAllPosts() (*[]model.Post, error) {
	var posts []model.Post
	err := p.db.Raw(`select id,user_id,tag,topics_name,md_plain_text,md_html_text,created_at,updated_at
	from posts p 
	order by created_at desc`).Scan(&posts).Error
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func (p *PostsRepositoryImpl) GetPostList() (*[]model.Post, error) {
	var posts []model.Post
	err := p.db.Raw(`select id,user_id,tag,topics_name,md_plain_text,md_html_text,created_at,updated_at,post_image,min_read
		from posts p 
		order by created_at desc
		limit 5`).Scan(&posts).Error

	if err != nil {
		return nil, err
	}

	return &posts, nil

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
