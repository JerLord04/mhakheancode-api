package domain

import (
	externalsource "mhakheancode/api/adapters/externalSource"
	"mhakheancode/api/model"
	"mhakheancode/api/ports/repository"
	"mhakheancode/api/ports/services"
)

type PostsDomain struct {
	PostRepository repository.PostRepository
	Externalsource externalsource.ConvertMdToHtml
}

// Get implements services.PostsPort.
func (p *PostsDomain) Get(id int32) (*model.Post, error) {
	data, err := p.PostRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *PostsDomain) CreatePosts(userId int, tag string, topicName string, fileByteArray []byte) bool {
	html := p.Externalsource.ConvertMdToHtml(fileByteArray)
	contentText := string(fileByteArray)
	contentHtml := string(html)
	posts := model.Post{
		UserID:      int32(userId),
		Tag:         tag,
		TopicsName:  topicName,
		MdPlainText: contentText,
		MdHTMLText:  contentHtml,
	}
	err := p.PostRepository.Save(&posts)
	if err != nil {
		return false
	}
	return true
}

func NewPostsDomain(postRepository repository.PostRepository, externalsource externalsource.ConvertMdToHtml) services.PostsPort {
	return &PostsDomain{
		PostRepository: postRepository,
		Externalsource: externalsource,
	}
}
