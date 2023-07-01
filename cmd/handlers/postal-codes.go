package handlers

import (
	"fmt"
	"strconv"

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
// @Router /api/plz/all [get]
func GetAllPostalCodes(c *fiber.Ctx, db *db.DB) error {
	var plz []models.PLZ

	db.Conn.Unscoped().Find(&plz)

	return c.JSON(plz)
}

// SearchPostalCodes queries Postal Codes
// @Summary Searches Postal Codes
// @Description Returns a list of postal codes based on the query. Upper limit of "limit" 50 records, Lower limit is 1 Record
// @Tags postal-codes
// @Accept json
// @Produce json
// @Param	q	query	string	false	"plz search query in all fields"
// @Param	limit	query	string	false	"amount of items, lower 1, upper 50"
// @Router /api/plz/search [get]
func SearchPostalCodes(c *fiber.Ctx, db *db.DB) error {
	query := c.Query("q")
	limit := c.Query("limit", "10")

	var plz = make([]models.PLZ, 0)

	limit_number, _ := strconv.Atoi(limit)

	if limit_number > 50 {
		limit_number = 50
	} else if limit_number < 1 {
		limit_number = 1
	}

	if (len(query) > 0) {
		db.Conn.Unscoped().Where(
			"LOWER(place) LIKE LOWER(?)", 
			fmt.Sprintf("%%%s%%", query),
		).Or(
			"LOWER(postal_code::text) LIKE LOWER(?)", 
			fmt.Sprintf("%s%%", query),
		).Limit(limit_number).Find(&plz)
	}

	return c.JSON(plz)
}

// FindPostalCodes finds Postal Codes
// @Summary Finds Postal Codes
// @Description Returns a list of postal codes based on the query.
// @Tags postal-codes
// @Accept json
// @Produce json
// @Param	place	query	string	false	"plz search by place"
// @Param	code	query	string	false	"plz search by code"
// @Router /api/plz/find [get]
func FindPostalCodes(c *fiber.Ctx, db *db.DB) error {
	place := c.Query("place")
	code := c.Query("code")

	var plz = make([]models.PLZ, 0)

	if (len(place) > 0 && len(code) > 0) {
		db.Conn.Unscoped().Where(
			"place LIKE ?", 
			fmt.Sprintf("%%%s%%", place),
		).Where(
			"postal_code::text LIKE ?", 
			fmt.Sprintf("%s%%", code),
		).Find(&plz)
	} else if (len(place) > 0) {
		db.Conn.Unscoped().Where(
			"place LIKE ?", 
			fmt.Sprintf("%%%s%%", place),
		).Find(&plz)
	} else if(len(code) > 0) {
		db.Conn.Unscoped().Where(
			"postal_code::text LIKE ?", 
			fmt.Sprintf("%s%%", code),
		).Find(&plz)
	}

	return c.JSON(plz)
}
