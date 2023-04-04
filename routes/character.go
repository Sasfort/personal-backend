package routes

import (
	"errors"

	"github.com/Sasfort/personal-backend/database"
	"github.com/Sasfort/personal-backend/models"
	"github.com/gofiber/fiber/v2"
)

type Character struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Appearance string `json:"appearance"`
	Origin     Origin `gorm:"foreignKey:OriginId"`
}

func CreateResponseCharacter(characterModel models.Character, origin Origin) Character {
	return Character{Id: characterModel.Id, Name: characterModel.Name, Gender: characterModel.Gender, Appearance: characterModel.Appearance, Origin: origin}
}

func CreateCharacter(c *fiber.Ctx) error {
	var character models.Character

	if err := c.BodyParser(&character); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var origin models.Origin
	if err := findOrigin(int(character.OriginId), &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&character)

	responseOrigin := CreateResponseOrigin(origin)
	responseCharacter := CreateResponseCharacter(character, responseOrigin)

	return c.Status(200).JSON(responseCharacter)
}

func ReadAllCharacters(c *fiber.Ctx) error {
	characters := []models.Character{}

	database.Database.Db.Find(&characters)
	responseCharacters := []Character{}
	for _, character := range characters {
		var origin models.Origin
		if err := findOrigin(int(character.OriginId), &origin); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		responseOrigin := CreateResponseOrigin(origin)
		responseCharacter := CreateResponseCharacter(character, responseOrigin)
		responseCharacters = append(responseCharacters, responseCharacter)
	}

	return c.Status(200).JSON(responseCharacters)
}

func ReadCharacter(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var character models.Character

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findCharacter(id, &character); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var origin models.Origin
	if err := findOrigin(int(character.OriginId), &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseOrigin := CreateResponseOrigin(origin)
	responseCharacter := CreateResponseCharacter(character, responseOrigin)

	return c.Status(200).JSON(responseCharacter)
}

func UpdateCharacter(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var character models.Character

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findCharacter(id, &character); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateCharacter struct {
		Name       string `json:"name"`
		Gender     string `json:"gender"`
		Appearance string `json:"appearance"`
		OriginId   uint   `json:"origin_id"`
	}

	var updateData UpdateCharacter

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	character.Name = updateData.Name
	character.Gender = updateData.Gender
	character.Appearance = updateData.Appearance
	character.OriginId = updateData.OriginId

	database.Database.Db.Save(&character)

	var origin models.Origin
	if err := findOrigin(int(character.OriginId), &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseOrigin := CreateResponseOrigin(origin)
	responseCharacter := CreateResponseCharacter(character, responseOrigin)

	return c.Status(200).JSON(responseCharacter)
}

func DeleteCharacter(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var character models.Character

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findCharacter(id, &character); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&character).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var origin models.Origin
	if err := findOrigin(int(character.OriginId), &origin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseOrigin := CreateResponseOrigin(origin)
	responseCharacter := CreateResponseCharacter(character, responseOrigin)

	return c.Status(200).JSON(responseCharacter)
}

func findCharacter(id int, character *models.Character) error {
	database.Database.Db.Find(&character, "id = ?", id)

	if character.Id == 0 {
		return errors.New("character with this id does not exist")
	}

	return nil
}
