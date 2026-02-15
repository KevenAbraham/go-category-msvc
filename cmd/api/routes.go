package main

import (
	"github.com/KevenAbraham/go-category-msvc/cmd/api/handler"
	"github.com/KevenAbraham/go-category-msvc/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	cache := repositories.NewCache()

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		handler.CreateCategory(ctx, cache)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		handler.ListCategory(ctx, cache)
	})
}
