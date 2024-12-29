package services

import (
	"mhakheancode/api/model"
	"mime/multipart"
)

type PostsPort interface {
	CreatePosts(userId int, tag string, topicName string, file *multipart.FileHeader, image *multipart.FileHeader, minRead string) bool
	GetById(id int32) (*model.Post, error)
	GetPostsList() (*[]model.Post, error)
	GetPagingPosts() (*[]model.Post, error)
}
