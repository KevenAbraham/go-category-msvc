package handler

import (
	"net/http"

	"github.com/KevenAbraham/go-category-msvc/internal/repositories"
	"github.com/KevenAbraham/go-category-msvc/internal/usecase"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, respository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil { 
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	usecase := usecase.NewCreateUsecase(respository)

	err := usecase.Create(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
