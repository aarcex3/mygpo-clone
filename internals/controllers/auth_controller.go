package controllers

import (
	"net/http"

	"github.com/aarcex3/mygpo-clone/internals/schemas"
	"github.com/aarcex3/mygpo-clone/internals/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Registration(ctx *gin.Context) {
	var form schemas.Registration

	// Validate the incoming form data
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Attempt to register the user
	if err := c.authService.Register(ctx, &form); err != nil {
		if err.Error() == "user already exists" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var form schemas.Login
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	token, err := c.authService.Authenticate(ctx, &form)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Login error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
	ctx.Header("Authorization", "Bearer "+token)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Logout successful"})
}
