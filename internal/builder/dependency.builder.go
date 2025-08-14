package builder

import (
	"github.com/gin-gonic/gin"

	// Importações que já estavam corretas
	userController "aurora.com/aurora-backend/internal/features/user/controller"
	userFactory "aurora.com/aurora-backend/internal/features/user/factory"
	userPersistence "aurora.com/aurora-backend/internal/features/user/gateway/repository"
	userSecurity "aurora.com/aurora-backend/internal/features/user/gateway/security"
	userService "aurora.com/aurora-backend/internal/features/user/service"
)

type Container struct {
	Router         *gin.Engine
	UserController *userController.UserController
}

func Build() (*Container, error) {

	userRepoImpl := userPersistence.NewUserInMemoryRepository()
	passwordHasher := userSecurity.NewBcryptHasher()
	
	userUseCaseFactory := userFactory.NewUseCaseFactory(userRepoImpl, passwordHasher)

	userSvc := userService.NewUserService(
		userUseCaseFactory.CreateUser,
		userUseCaseFactory.GetUser,
	)

	userCtrl := userController.NewUserController(userSvc)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		userRoutes := v1.Group("/users")
		{
			userRoutes.POST("", userCtrl.CreateUser)
			userRoutes.GET("/:id", userCtrl.GetUser)
		}
	}

	return &Container{
		Router:         router,
		UserController: userCtrl,
	}, nil
}