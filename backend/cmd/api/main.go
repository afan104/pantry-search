package main

import (
	"database/sql"
	"log"

	"github.com/afan104/pantry-search/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)



func main() {
    // http router and address
	router := gin.Default()
	address := ":3000"

	// check that sql is reachable
	db,e := sql.Open("sqlite3", "./db/pantry.db")
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
	defer db.Close()

	// // check that you can ping the database and log error
	if e := db.Ping(); e != nil {
		log.Fatalf("Error: %v", e)
	}


	// controllers
	router.GET("/getIngredients", handlers.GetIngredients(db))
	router.GET("/getIngredient/:ingredient", handlers.GetIngredient(db))
	// router.POST("/postIngredient/:ingredient", postIngredient)
	// router.DELETE("/deleteIngredient/:ingredient", deleteIngredient)
	// router.PUT("/putIngredient/:ingredient", putIngredient)

	// attaach router to server; log success/error
	if e := router.Run(address); e != nil {
		log.Fatal(e)
	}

}