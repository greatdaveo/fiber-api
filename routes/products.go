package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greatdaveo/fiber-api/database"
	"github.com/greatdaveo/fiber-api/models"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(201).JSON(responseProduct)
}

