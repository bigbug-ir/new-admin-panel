package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/cmd/web/pkg/models"
)

type Product struct {
	ID          uint    `json:"id"`
	User        User    `json:"user"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func createResponseProduct(productModel models.Product, user User) Product {
	return Product{
		ID:          productModel.ID,
		User:        user,
		Name:        productModel.Name,
		Price:       productModel.Price,
		Description: productModel.Description,
	}
}
func findProduct(productName string, productModel *models.Product) error {
	database.Database.DB.Find(&productModel, "name = ?", productName)
	if productModel.Name == productName {
		return errors.New("product name not found")
	}
	return nil
}
func findProductByID(id int, product *models.Product) error {
	database.Database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}
func CreateProduct(c *fiber.Ctx) error {
	userName := c.Params("username")
	var user models.User
	if err := findUser(userName, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product.UserRefer = user.ID
	database.Database.DB.Create(&product)
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product, responseUser)
	return c.Status(200).JSON(responseProduct)
}
func UpdateProduct(c *fiber.Ctx) error {
	userName := c.Params("username")
	var user models.User
	if err := findUser(userName, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	productName := c.Params("productname")
	var product models.Product
	if err := findProduct(productName, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
	}
	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if updateData.Name != "" {
		product.Name = updateData.Name
	}
	if updateData.Description != "" {
		product.Description = updateData.Description
	}
	if updateData.Price != 0 {
		product.Price = updateData.Price
	}
	database.Database.DB.Save(&product)
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product, responseUser)
	return c.Status(200).JSON(responseProduct)
}
func DeleteProduct(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	var product models.Product
	productName := c.Params("productName")
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findProduct(productName, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.DB.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON("successfully deleted product")
}
func GetProductData(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	var product models.Product
	productName := c.Params("productName")
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findProduct(productName, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product, responseUser)
	return c.Status(200).JSON(responseProduct)
}

func UpdateProductByID(c *fiber.Ctx) error {
	userName := c.Params("username")
	var user models.User
	if err := findUser(userName, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	productId, err := c.ParamsInt("productid")
	if err != nil {
		return c.Status(400).JSON("please ensur correct product id")
	}
	var product models.Product
	if err := findProductByID(productId, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
	}
	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if updateData.Name != "" {
		product.Name = updateData.Name
	}
	if updateData.Description != "" {
		product.Description = updateData.Description
	}
	if updateData.Price != 0 {
		product.Price = updateData.Price
	}
	database.Database.DB.Save(&product)
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product, responseUser)
	return c.Status(200).JSON(responseProduct)
}
func DeleteProductByID(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	var product models.Product
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	productId, err := c.ParamsInt("productid")
	if err != nil {
		return c.Status(400).JSON("please ensur correct product id")
	}
	if err := findProductByID(productId, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.DB.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON("successfully deleted product")
}
func GetProductDataByID(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	var product models.Product

	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	productId, err := c.ParamsInt("productid")
	if err != nil {
		return c.Status(400).JSON("please ensur correct product id")
	}

	if err := findProductByID(productId, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product, responseUser)
	return c.Status(200).JSON(responseProduct)
}
func GetAllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	responseProducts := []Product{}
	for _, product := range products {
		var user models.User
		database.Database.DB.Find(&user, "id = ?", product.UserRefer)
		responseUser := createResponseUser(user)
		responseProduct := createResponseProduct(product, responseUser)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}
func SearchProduct(c *fiber.Ctx) error {
	productName := c.Params("productName")
	products := []models.Product{}
	responseProducts := []Product{}
	for _, product := range products {
		var user models.User
		database.Database.DB.Find(&product, "name LIKE ?", productName)
		database.Database.DB.Find(&user, "id = ?", product.UserRefer)
		responseUser := createResponseUser(user)
		responseProduct := createResponseProduct(product, responseUser)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}
