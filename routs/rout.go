package routs

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	NewHandShakeRouter(api)
	post := app.Group("/posts")
	NewPostsRouter(post, db)
	tag := app.Group("/tags")
	NewTagsRouter(tag, db)
}
