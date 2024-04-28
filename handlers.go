package main

import (
	"github.com/gofiber/fiber/v2"
	"github.io/ThomasCDeveloper/go-postgres-boilerplate/models"
)

func (r *Repository) newItem(c *fiber.Ctx) error {
	data := c.Query("data")

	err := r.DB.Create(&models.Item{
		Data: data,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
