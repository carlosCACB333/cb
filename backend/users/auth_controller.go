package users

import (
	"cb/libs"
	"cb/utils"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
)

func AuthRegister(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	// validate fields
	errors := utils.ValidateFields(user)
	if errors != nil {
		c.JSON(400, utils.Response(
			"error", "Invalid fields",
			errors,
			nil,
		))
		return
	}

	user.Email = utils.NormalizeEmail(user.Email)
	user.Password = libs.HashPassword(user.Password)
	user.Status = "active"

	// create user
	if err := libs.DBInit().Create(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to create user",
			nil,
			nil,
		))
		return
	}

	c.JSON(200, user)

}

func AuthLogin(c *gin.Context) {
	var login LoginDTO
	var user User

	c.BindJSON(&login)

	// validate fields
	errors := utils.ValidateFields(login)
	if errors != nil {
		c.JSON(400, utils.Response(
			"error", "Invalid fields",
			errors,
			nil,
		))
		return
	}

	login.Email = utils.NormalizeEmail(login.Email)

	// check if user exists
	if err := libs.DBInit().Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Invalid email",
			nil,
			nil,
		))
		return
	}

	// check if password is correct
	if !libs.CheckPassword(login.Password, user.Password) {
		c.JSON(400, utils.Response(
			"error", "Invalid password",
			nil,
			nil,
		))
		return
	}

	// generate token
	token, tkerr := libs.GenerateToken(user.ID)
	if tkerr != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to generate token",
			nil,
			nil,
		))

		return
	}

	c.JSON(400, utils.Response(
		"success", "Login successful",
		gin.H{"token": token, "user": user},
		nil,
	))

}

func ChangePassword(c *gin.Context) {
	clearkUser := c.MustGet("user").(*clerk.User)
	var reset ChangePasswordDTO
	c.BindJSON(&reset)

	// validate fields
	if err := utils.ValidateFields(reset); err != nil {
		c.JSON(400, utils.Response(
			"error", "Invalid fields",
			err,
			nil,
		))

		return
	}
	if reset.NewPassword != reset.ConfirmPassword {
		c.JSON(400, utils.Response(
			"error", "Passwords do not match",
			nil,
			nil,
		))
		return
	}
	var user User
	if err := libs.DBInit().Where("id = ?", clearkUser.ID).First(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to fetch user",
			nil,
			nil,
		))
		return
	}

	// check if old password is correct
	if !libs.CheckPassword(reset.OldPassword, user.Password) {
		c.JSON(400, utils.Response(
			"error", "Invalid old password",
			nil,
			nil,
		))

		return
	}

	// update password
	user.Password = libs.HashPassword(reset.NewPassword)
	if err := libs.DBInit().Save(&user).Error; err != nil {
		c.JSON(400, utils.Response(
			"error", "Unable to update password",
			nil,
			nil,
		))
		return
	}

	c.JSON(200, utils.Response(
		"success", "Password updated successfully",
		user,
		nil,
	))

}
