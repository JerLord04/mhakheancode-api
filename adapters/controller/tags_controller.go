package controller

import (
	"mhakheancode/api/ports/services"
	"mhakheancode/api/utils"

	"github.com/gofiber/fiber/v2"
)

type TagsController struct {
	TagsPort services.TagsPort
}

func (t *TagsController) GetAllTags(c *fiber.Ctx) error {
	data, err := t.TagsPort.GetAllTags()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.Response{
				Status:  false,
				Paylaod: "Save files incompleted.",
			})
	}
	return c.Status(fiber.StatusOK).JSON(
		utils.Response{
			Status:  true,
			Paylaod: data,
		})

}
