package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/cmd/web/pkg/models"
)

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}

func createResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		UserName:  userModel.UserName,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Password:  userModel.Password,
		Age:       userModel.Age,
		Gender:    userModel.Gender,
		CreatedAt: userModel.CreatedAt,
	}
}
func findUserById(id uint, user *models.User) error {
	database.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}
func findUser(userName string, userModel *models.User) error {
	database.Database.DB.Find(&userModel, "user_name=?", userName)
	if userModel.UserName == userName {
		return nil
	}
	return errors.New("user not found")
}
func LoginUser(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	password := c.Params("password")
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON("passwird or username is not correctly")
	}
	if user.Password != password {
		return c.Status(400).JSON("passwird or username is not correctly")
	}
	responseGetUser := createResponseUser(user)
	return c.Status(200).JSON(responseGetUser)
}
func existsUser(userModel *models.User) error {
	users := []models.User{}
	database.Database.DB.Find(&users)
	for _, user := range users {
		if user.UserName == userModel.UserName || user.Email == userModel.Email || user.Phone == userModel.Phone {
			return errors.New("user dose Exist in database")
		}
	}
	return nil
}
func existsUserByEmail(email string) error {
	users := []models.User{}
	database.Database.DB.Find(&users)
	for _, user := range users {
		if user.Email == email {
			return errors.New("user dose Exist in database")
		}
	}
	return nil
}
func existsUserByUserName(userName string) error {
	users := []models.User{}
	database.Database.DB.Find(&users)
	for _, user := range users {
		if user.UserName == userName {
			return errors.New("user dose Exist in database")
		}
	}
	return nil

}
func existsUserByPhone(phone string) error {
	users := []models.User{}
	database.Database.DB.Find(&users)
	for _, user := range users {
		if user.Phone == phone {
			return errors.New("user dose Exist in database")
		}
	}
	return nil
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON("body is not a valid JSON")
	}
	if err := existsUser(&user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	database.Database.DB.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// username := c.Params("username")
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		ID        uint   `json:"id"`
		UserName  string `json:"user_name"`
		Phone     string `json:"phone"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Age       int    `json:"age"`
		Gender    string `json:"gender"`
	}
	var updateData *UpdateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	if updateData.UserName != "" {
		if err := existsUserByUserName(updateData.UserName); err != nil {
			return c.Status(500).JSON(err.Error())
		}
		user.UserName = updateData.UserName
	}
	if updateData.Phone != "" {
		if err := existsUserByPhone(updateData.Phone); err != nil {
			return c.Status(500).JSON(err.Error())
		}
		user.Phone = updateData.Phone
	}
	if updateData.FirstName != "" {
		user.FirstName = updateData.FirstName
	}
	if updateData.LastName != "" {
		user.LastName = updateData.LastName
	}
	if updateData.Email != "" {
		if err := existsUserByEmail(updateData.Email); err != nil {
			return c.Status(500).JSON(err.Error())
		}
		user.Email = updateData.Email
	}
	if updateData.Age != 0 {
		user.Age = updateData.Age
	}
	if updateData.Gender != "" {
		user.Gender = updateData.Gender
	}
	if updateData.Password != "" {
		user.Password = updateData.Password
	}
	database.Database.DB.Save(&user)
	responseUpdate := createResponseUser(user)
	return c.Status(200).JSON(responseUpdate)
}
func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.DB.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("successfully deleted user")
}
func GetUserData(c *fiber.Ctx) error {
	var user models.User
	username := c.Params("username")
	if err := findUser(username, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseGetUser := createResponseUser(user)
	return c.Status(200).JSON(responseGetUser)
}
func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.DB.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := createResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
