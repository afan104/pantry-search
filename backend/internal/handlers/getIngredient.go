package handlers

import (
	"database/sql"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func GetIngredient(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itm Item
		var ingredientType sql.NullString
		ingredient := c.Param("ingredient")
		err := db.QueryRow(`SELECT * FROM pantry WHERE ingredient=?`,ingredient).Scan(
			&itm.Id,
			&itm.Ingredient,
			&ingredientType,
			&itm.Quantity,
			&itm.Units,
			&itm.DateUpdated,
			&itm.ExpectedExpiry,
		)
		if err != nil {
			// check if empty
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "no ingredient with that name"})
			// or if other error
			} else {
				log.Printf("Error grabbing ingredient from db: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}
	if ingredientType.Valid {
		itm.IngredientType = ingredientType.String
	}
	c.JSON(http.StatusOK, itm)

	}
}