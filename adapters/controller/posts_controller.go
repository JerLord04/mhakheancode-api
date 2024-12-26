package controller

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("====> %s", file.Filename)
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
	fileByteArray, err := file.Open()
	if err != nil {
		return err
	}
	fileByteArray.Close()
	openFile, err := ioutil.ReadAll(fileByteArray)
	if err != nil {
		return err
	}
	userIdToNumber, _ := strconv.Atoi(userId)
	res := p.PostsPort.CreatePosts(userIdToNumber, tag, title, openFile)
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
