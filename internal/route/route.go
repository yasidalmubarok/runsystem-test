package route

import (
	"runsystem-test/internal/handler"
	"runsystem-test/internal/repository"
	"runsystem-test/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/healthy", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "healthy"})
	})

	v1 := r.Group("api/v1")
	{
		users := v1.Group("users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUserByID)
		}
	}

	return r
}
