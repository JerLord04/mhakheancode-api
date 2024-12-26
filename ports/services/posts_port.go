package services

import "mhakheancode/api/model"

type PostsPort interface {
	CreatePosts(userId int, tag string, topicName string, fileByteArray []byte) bool
	Get(id int32) (*model.Post, error)
}
