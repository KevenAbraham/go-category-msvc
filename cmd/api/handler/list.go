package handler

import (
	"net/http"

	"github.com/KevenAbraham/go-category-msvc/internal/repositories"
	"github.com/KevenAbraham/go-category-msvc/internal/usecase"
	"github.com/gin-gonic/gin"
)

func ListCategory(ctx *gin.Context, respository repositories.ICategoryRepository) {
	usecase := usecase.NewListUsecase(respository)

	categories, err := usecase.List()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
