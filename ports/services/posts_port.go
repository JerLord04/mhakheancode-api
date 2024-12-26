package services

type PostsPort interface {
	CreatePosts(userId int, tag string, topicName string, fileByteArray []byte) bool
}
