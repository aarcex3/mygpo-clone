package controllers

import (
	"net/http"

	"github.com/aarcex3/mygpo-clone/internals/schemas"
	"github.com/aarcex3/mygpo-clone/internals/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Registration(ctx *gin.Context) {
	var user schemas.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if err := c.AuthService.Register(&user); err != nil {

	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Login successful"})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Logout successful"})
}
