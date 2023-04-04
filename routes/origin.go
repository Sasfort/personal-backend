package routes

import (
	"errors"

	"github.com/Sasfort/personal-backend/database"
	"github.com/Sasfort/personal-backend/models"
	"github.com/gofiber/fiber/v2"
)

type Origin struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

func CreateResponseOrigin(originModel models.Origin) Origin {
	return Origin{Id: originModel.Id, Title: originModel.Title}
}

func CreateOrigin(c *fiber.Ctx) error {
	var origin models.Origin

	if err := c.BodyParser(&origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&origin)
	responseOrigin := CreateResponseOrigin(origin)

	return c.Status(200).JSON(responseOrigin)
}

func ReadAllOrigins(c *fiber.Ctx) error {
	origins := []models.Origin{}

	database.Database.Db.Find(&origins)
	responseOrigins := []Origin{}
	for _, origin := range origins {
		responseOrigin := CreateResponseOrigin(origin)
		responseOrigins = append(responseOrigins, responseOrigin)
	}

	return c.Status(200).JSON(responseOrigins)
}

func ReadOrigin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var origin models.Origin

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findOrigin(id, &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseOrigin := CreateResponseOrigin(origin)

	return c.Status(200).JSON(responseOrigin)
}

func UpdateOrigin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var origin models.Origin

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findOrigin(id, &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateOrigin struct {
		Title string `json:"title"`
	}

	var updateData UpdateOrigin

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	origin.Title = updateData.Title

	database.Database.Db.Save(&origin)

	responseOrigin := CreateResponseOrigin(origin)

	return c.Status(200).JSON(responseOrigin)
}

func DeleteOrigin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var origin models.Origin

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findOrigin(id, &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&origin).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseOrigin := CreateResponseOrigin(origin)

	return c.Status(200).JSON(responseOrigin)
}

func findOrigin(id int, origin *models.Origin) error {
	database.Database.Db.Find(&origin, "id = ?", id)

	if origin.Id == 0 {
		return errors.New("origin with this id does not exist")
	}

	return nil
}
