package routs

import (
	"mhakheancode/api/adapters/controller"
	externalsource "mhakheancode/api/adapters/externalSource"
	repositoryimpl "mhakheancode/api/adapters/repositoryImpl"
	"mhakheancode/api/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewPostsRouter(group fiber.Router, db *gorm.DB) {
	repo := repositoryimpl.NewPostsRepositoryImpl(db)
	externalSource := externalsource.NewConvertMdToHtmlStruct()
	inputSteamManage := externalsource.NewByteArrayToImageStruct()
	postController := &controller.PostsController{
		PostsPort: domain.NewPostsDomain(repo, externalSource, inputSteamManage),
	}

	group.Get("/create", postController.CreatePosts)
	group.Get("/get", postController.GetPostsById)
	group.Get("/getPostsList", postController.GetPostsList)
	group.Get("/getAllPosts", postController.GetPagingPosts)
}
