package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	as AuthService
}

func Controller(authService AuthService) *controller {
	return &controller{as: authService}
}

func (c *controller) Register(ctx *gin.Context) {
	var form RegistrationForm

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := c.as.Register(ctx, &form); err != nil {
		if err.Error() == "user already exists" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (c *controller) Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	token, err := c.as.Authenticate(ctx, &form)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Login error"})
		return
	}

	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func (c *controller) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Logout successful"})
}
