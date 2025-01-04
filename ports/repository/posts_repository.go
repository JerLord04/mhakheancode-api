package repository

import "mhakheancode/api/model"

type PostRepository interface {
	Save(model *model.Post) error
	Get(id int32) (model *model.Post, err error)
	GetPostList() (*[]model.Post, error)
	GetAllPosts() (*[]model.Post, error)
}
