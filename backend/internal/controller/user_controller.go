package controller

import (
	"net/http"

	"github.com/user-authentication-go/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct {
    userRepo *repository.UserRepository
}

func NewUserController(userRepo *repository.UserRepository) *UserController {
    return &UserController{
        userRepo: userRepo,
    }
}

func (c *UserController) GetProfile(ctx *gin.Context) {
    userID := ctx.GetString("user_id")
    // userIDを使用した実装
    ctx.JSON(http.StatusOK, gin.H{
        "user_id": userID,
        "message": "get profile",
    })
}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
    userID := ctx.GetString("user_id")
    // userIDを使用した実装
    ctx.JSON(http.StatusOK, gin.H{
        "user_id": userID,
        "message": "update profile",
    })
}
