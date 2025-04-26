package api

import (
	"net/http"

	"github.com/ChayanDass/beneficiary-manager/pkg/db"
	"github.com/ChayanDass/beneficiary-manager/pkg/models"
	"github.com/gin-gonic/gin"
)

// SubmitCredential handles the submission of user credentials
func SubmitCredential(c *gin.Context) {
	var credential models.User

	// Bind the incoming JSON to the User struct (credential)
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Error:   err.Error(),
		})
		return
	}

	// Check if the user exists by email (email is unique)
	var user models.User
	err := db.DB.Where("email = ?", credential.Email).First(&user).Error

	if err != nil {
		// If user does not exist, create a new user
		if err.Error() == "record not found" {
			// Create a new user if it doesn't exist
			if createErr := db.DB.Create(&credential).Error; createErr != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Failed to create user",
					Error:   createErr.Error(),
				})
				return
			}
			// Return success if new user is created
			c.JSON(http.StatusOK, models.SuccessResponse{
				Code:    http.StatusOK,
				Message: "User created and credential submitted successfully",
				Data:    credential,
			})
			return
		}
		// Return any other error encountered while querying the user
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Error:   err.Error(),
		})
		return
	}

	// If user exists, we can submit the credential (optional modification)
	// Here, I'm assuming the "credential" is the user model (as per your logic)
	if createErr := db.DB.Create(&credential).Error; createErr != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to submit credential",
			Error:   createErr.Error(),
		})
		return
	}

	// Return success if credential is submitted for an existing user
	c.JSON(http.StatusOK, models.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Credential submitted successfully",
		Data:    credential,
	})
}
