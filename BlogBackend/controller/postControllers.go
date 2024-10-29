package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tegveer-singh123/blog/database"
	"github.com/tegveer-singh123/blog/models"
	"github.com/tegveer-singh123/blog/util"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost models.Blog

	err := c.BodyParser(&blogPost)
	if err != nil {
		fmt.Println("Unable to parse body")
	}

	err = database.DB.Create(&blogPost).Error
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid PayLoad",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Congratulations!, Your Post is Live...",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offSet := (page - 1) * limit

	var total int64
	var getBlog []models.Blog

	database.DB.Preload("User").Offset(offSet).Limit(limit).Find(&getBlog)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": getBlog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var blogPost models.Blog

	database.DB.Where("id=?", id).Preload("User").First(&blogPost)
	return c.JSON(fiber.Map{
		"data": blogPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	blog := models.Blog{
		ID: uint(id),
	}

	err := c.BodyParser(&blog)
	if err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)

	return c.JSON(fiber.Map{
		"message": "Congratulations!, Your Post is Updated...",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id = ?", id).Preload("User").Find(&blog)

	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	blog := models.Blog{
		ID: uint(id),
	}

	deleteQuery := database.DB.Delete(&blog)

	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	return c.JSON(fiber.Map{
		"message":"Post Deleted Successfully",
	})
}
