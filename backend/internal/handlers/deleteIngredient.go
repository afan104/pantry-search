package handlers

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteIngredient(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ingredient:=c.Param("ingredient")
		res,err:=db.Exec(`DELETE FROM pantry WHERE ingredient=?`,ingredient)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
			return
		}
		// if deleting an ingredient that doesn't exist
		if rowsAffected,_:=res.RowsAffected(); rowsAffected==0 {
			c.JSON(http.StatusNotFound,gin.H{"error": "ingredient not found"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"message": "ingredient deleted"})
	}
}