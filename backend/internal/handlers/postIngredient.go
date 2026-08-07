package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostQuantityPayload struct {
	Quantity float64 `json:"quantity" binding:"required"`
}

func PostIngredient(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload PostQuantityPayload
		ingredient := c.Param("ingredient")
		// if payload doesn't have quantity
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error": "invalid body request"})
			return
		}
		res,err := db.Exec(`UPDATE pantry SET quantity=quantity+? WHERE ingredient=?`, payload.Quantity, ingredient)
		// if update fails
		if err != nil {
				c.JSON(http.StatusInternalServerError,gin.H{"error": "internal server error"})
				return
		}
		// if updating an ingredient that doesn't exist
		if rowsAffected,_:=res.RowsAffected(); rowsAffected==0 {
			c.JSON(http.StatusNotFound,gin.H{"error": "ingredient not found"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"message":"ingredient quantity updated"})
	}
}