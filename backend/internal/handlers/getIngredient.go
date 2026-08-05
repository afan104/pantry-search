package handlers

import {

}

// getIngredient handler returns handler function
func getIngredient(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// initialize rows and error
		// check that rows exist in db
	}
	// keep rows open (why stop closing here instead of in previous brackets?)
	// initialize products
	// loop through rows to create products struct
	// // make sure no empty values for required rows
	// make sure products is not empty
	// result
}
// 