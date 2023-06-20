package handlers

import (
	"github.com/3n3a/plz-api/cmd/db"
	"github.com/3n3a/plz-api/cmd/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllPostalCodes lists all Postal Codes
// @Summary Get all Postal Codes
// @Description Get all Postal Codes
// @Tags postal-codes
// @Accept json
// @Produce json
// @Router /api/postal-codes [get]
func GetAllPostalCodes(c *fiber.Ctx, db *db.DB) error {
	var plz []models.PLZ

	db.Conn.Unscoped().Find(&plz)
	
	return c.JSON(plz)
}