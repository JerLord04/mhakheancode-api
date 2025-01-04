package routs

import (
	"mhakheancode/api/adapters/controller"
	repositoryimpl "mhakheancode/api/adapters/repositoryImpl"
	"mhakheancode/api/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewTagsRouter(group fiber.Router, db *gorm.DB) {
	repo := repositoryimpl.NewTagsRepositoryImpl(db)
	tagController := &controller.TagsController{
		TagsPort: domain.NewTagsDomain(repo),
	}
	group.Get("/getAllTags", tagController.GetAllTags)
}
