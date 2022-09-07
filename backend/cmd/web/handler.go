package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/cmd/web/pkg/routes"
)

func setupRoutes(app *fiber.App) {
	// user
	app.Post("/home/user", routes.CreateUser)                                     // create user
	app.Get("/home/user/username=:username&password=:password", routes.LoginUser) // login user
	app.Put("/home/user/:username", routes.UpdateUser)                            // update user
	app.Delete("/home/user/:username", routes.DeleteUser)                         // delete user
	app.Get("/home/user/:username", routes.GetUserData)                           // get user
	//product for user
	app.Post("/home/user/:username/product", routes.CreateProduct)                //Create product
	app.Put("/home/user/:username/product/:productname", routes.UpdateProduct)    //Update product by productName
	app.Delete("/home/user/:username/product/:productname", routes.DeleteProduct) //Delete Product by productName
	app.Get("/home/user/:username/product/:productname", routes.GetProductData)   //Get Product by productName
	// product for user
	app.Put("/home/user/:username/product/:productid", routes.UpdateProductByID)    //Update Product by ID
	app.Delete("/home/user/:username/product/:productid", routes.DeleteProductByID) //Delete Product by ID
	app.Get("/home/user/:username/product/:productid", routes.GetProductDataByID)   //Get Product by ID
	// get all data
	app.Get("/home/product", routes.GetAllProducts) //Get all products
	app.Get("/home/users", routes.GetAllUsers)      //Get all users
}
