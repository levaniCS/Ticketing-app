package routes

import (
	"net/http"

	"example.com/ticketing/models"
	"example.com/ticketing/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logged in successfully", "token": token})
}
