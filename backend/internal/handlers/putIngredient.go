package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PutIngredientPayload struct {
	IngredientType string `json:"ingredientType"`
	Quantity float64 `json:"quantity" binding:"required"`
	Units string `json:"units" binding:"required"`
}

func PutIngredient(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload PutIngredientPayload
		expectedExpiry := time.Now() // + i'll implement this properly later
		ingredient := c.Param("ingredient")

		if err:=c.ShouldBindJSON(&payload);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid body request"})
			return
		}

		if _,err:= db.Exec(`INSERT INTO pantry (ingredient,ingredientType,quantity,units,dateUpdated,expectedExpiry) VALUES (?,?,?,?,?,?) ON CONFLICT (ingredient)
							DO UPDATE SET ingredientType=excluded.ingredientType,quantity=excluded.quantity,units=excluded.units,dateUpdated=excluded.dateUpdated,expectedExpiry=excluded.expectedExpiry`,
			ingredient,payload.IngredientType,payload.Quantity,payload.Units,time.Now(),expectedExpiry); err != nil {
				c.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})
				return
			}
		c.JSON(http.StatusOK, gin.H{"message": "ingredient updated"})
	}
}