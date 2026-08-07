package handlers

import (
	"database/sql"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

// define item struct
type Item struct {
	Id int `json:"id"`
	Ingredient string `json:"ingredient"`
	IngredientType string `json:"ingredientType"`
	Quantity float64 `json:"quantity"`
	Units string `json:"units"`
	DateUpdated time.Time `json:"dateUpdated"`
	ExpectedExpiry time.Time `json:"expectedExpiry"`
}

// getIngredient handler returns handler function
func GetIngredients(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows,err := db.Query(`
		SELECT * FROM pantry`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	// loop through rows to get what you need
	items := []Item{}
	for rows.Next() {
		var itm Item
		if err := rows.Scan(
			&itm.Id,
			&itm.Ingredient,
			&itm.IngredientType,
			&itm.Quantity,
			&itm.Units,
			&itm.DateUpdated,
			&itm.ExpectedExpiry,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
			return
		}
		items=append(items,itm)
	}
	c.JSON(http.StatusOK, items)
	}
	}
