package users

import (
	"cb/libs"
	"cb/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userData SyncUserDTO
	c.BindJSON(&userData)

	user := User{
		Email:         utils.NormalizeEmail(userData.Email),
		FirstName:     userData.FirstName,
		LastName:      userData.LastName,
		Photo:         userData.Photo,
		Phone:         userData.Phone,
		SocialAccount: userData.SocialAccount,
		Status:        "active",
	}

	// find user
	var userFound User
	if err := libs.DBInit().Where("email = ?", user.Email).First(&userFound).Error; err == nil {
		// update user
		user.Status = "active"
		if err := libs.DBInit().Model(&userFound).Updates(user).Error; err != nil {
			c.JSON(400, utils.Response(
				"error", "Unable to update user",
				nil,
				nil,
			))
			return
		}
		c.JSON(200, utils.Response(
			"success", "User updated successfully",
			nil,
			nil,
		))
		return
	}

	// create user
	if err := libs.DBInit().Create(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to create user",
			nil,
			nil,
		))
		return
	}

	c.JSON(200, utils.Response(
		"success", "User created successfully",
		nil,
		nil,
	))

}

func UpdateUser(c *gin.Context) {
	var userData SyncUserDTO
	c.BindJSON(&userData)

	newUser := User{
		Email:         utils.NormalizeEmail(userData.Email),
		FirstName:     userData.FirstName,
		LastName:      userData.LastName,
		Photo:         userData.Photo,
		Phone:         userData.Phone,
		SocialAccount: userData.SocialAccount,
		Status:        "active",
	}

	// get user
	var user User
	if err := libs.DBInit().Where("social_account = ?", newUser.SocialAccount).First(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to get user",
			nil,
			nil,
		))
		return
	}

	// update user
	if err := libs.DBInit().Model(&user).Updates(newUser).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to update user",
			nil,
			nil,
		))
		return
	}

	c.JSON(200, utils.Response(
		"success", "User updated successfully",
		nil,
		nil,
	))

}

func DeleUser(c *gin.Context) {
	social_account := c.Param("social_account")

	// get user
	var user User
	if err := libs.DBInit().Where("social_account = ?", social_account).First(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to get user",
			nil,
			nil,
		))
		return
	}

	// update status
	if err := libs.DBInit().Model(&user).Updates(User{Status: "deleted"}).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to delete user",
			nil,
			nil,
		))
		return
	}
	c.JSON(200, utils.Response(
		"success", "User deleted successfully",
		nil,
		nil,
	))
}
