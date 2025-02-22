package controller

import (
	"fmt"
	"mhakheancode/api/ports/services"
	"mhakheancode/api/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PostsController struct {
	PostsPort services.PostsPort
}

func (p *PostsController) CreatePosts(c *fiber.Ctx) error {
	userId := c.Query("userId")
	tag := c.Query("tag")
	title := c.Query("title")
	file, err := c.FormFile("file")
	image, err := c.FormFile("image")
	minRead := c.Query("minRead")
	fileName := strings.Split(file.Filename, ".")
	if fileName[1] != "md" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.Response{
				Status:  false,
				Paylaod: "Not md file.",
			})
	}
	if err != nil {
		return err
	}
	userIdToNumber, _ := strconv.Atoi(userId)
	res := p.PostsPort.CreatePosts(userIdToNumber, tag, title, file, image, minRead)
	if res {
		return c.Status(fiber.StatusOK).JSON(
			utils.Response{
				Status:  true,
				Paylaod: "Save files completed.",
			})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.Response{
				Status:  false,
				Paylaod: "Save files incompleted.",
			})
	}

}

func (p *PostsController) GetPostsById(c *fiber.Ctx) error {
	id := c.Query("id")
	idToNumber, _ := strconv.Atoi(id)
	data, err := p.PostsPort.GetById(int32(idToNumber))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.Response{
				Status:  true,
				Paylaod: err,
			})
	}
	fmt.Printf(data.MdHTMLText)
	return c.Status(fiber.StatusOK).JSON(
		utils.Response{
			Status:  true,
			Paylaod: data,
		})
}

func (p *PostsController) GetPostsList(c *fiber.Ctx) error {
	data, err := p.PostsPort.GetPostsList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.Response{
				Status:  true,
				Paylaod: err,
			})
	}
	return c.Status(fiber.StatusOK).JSON(
		utils.Response{
			Status:  true,
			Paylaod: data,
		})
}

func (p *PostsController) GetPagingPosts(c *fiber.Ctx) error {
	data, err := p.PostsPort.GetPagingPosts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.Response{
				Status:  true,
				Paylaod: err,
			})
	}
	return c.Status(fiber.StatusOK).JSON(
		utils.Response{
			Status:  true,
			Paylaod: data,
		})
}
