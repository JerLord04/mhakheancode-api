package domain

import (
	externalsource "mhakheancode/api/adapters/externalSource"
	"mhakheancode/api/model"
	"mhakheancode/api/ports/repository"
	"mhakheancode/api/ports/services"
	"mime/multipart"
)

type PostsDomain struct {
	PostRepository   repository.PostRepository
	Externalsource   externalsource.ConvertMdToHtml
	InputSteamManage externalsource.InputSteamManage
}

// GetAllList implements services.PostsPort.
func (p *PostsDomain) GetPagingPosts() (*[]model.Post, error) {
	data, err := p.PostRepository.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PostsDomain) GetPostsList() (*[]model.Post, error) {
	data, err := p.PostRepository.GetPostList()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PostsDomain) GetById(id int32) (*model.Post, error) {
	data, err := p.PostRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *PostsDomain) CreatePosts(userId int, tag string, topicName string, file *multipart.FileHeader, image *multipart.FileHeader, minRead string) bool {
	fileByteArray, err := p.InputSteamManage.ConvertFileToByreArray(file)
	if err != nil {
		return false
	}
	imageByteArray, errImageByteArray := p.InputSteamManage.ConvertFileToByreArray(image)

	if errImageByteArray != nil {
		return false
	}
	imageBase64, errImageBase64 := p.InputSteamManage.ConvertByteArrayToBase64(imageByteArray)
	if errImageBase64 != nil {
		return false
	}
	html := p.Externalsource.ConvertMdToHtml(fileByteArray)
	contentText := string(fileByteArray)
	contentHtml := string(html)
	posts := model.Post{
		UserID:      int32(userId),
		Tag:         tag,
		TopicsName:  topicName,
		MdPlainText: contentText,
		MdHTMLText:  contentHtml,
		PostImage:   imageBase64,
		MinRead:     minRead,
	}
	errSave := p.PostRepository.Save(&posts)
	if errSave != nil {
		return false
	}
	return true
}

func NewPostsDomain(postRepository repository.PostRepository, externalsource externalsource.ConvertMdToHtml, inputSteamManage externalsource.InputSteamManage) services.PostsPort {
	return &PostsDomain{
		PostRepository:   postRepository,
		Externalsource:   externalsource,
		InputSteamManage: inputSteamManage,
	}
}
