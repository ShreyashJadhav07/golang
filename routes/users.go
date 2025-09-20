package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	// Signup logic will be implemented here
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not create user", "error": err.Error()})
		return
	}

	context.JSON(201, gin.H{"message": "User created successfully", "user_id": user.ID})

}

