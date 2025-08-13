package builder

import (
	"github.com/gin-gonic/gin"

	// Importações atualizadas com os novos caminhos
	userController "aurora.com/aurora-backend/internal/features/user/controller"
	userPersistence "aurora.com/aurora-backend/internal/features/user/gateway/repository" // Implementação
	userSecurity "aurora.com/aurora-backend/internal/features/user/gateway/security"
	userService "aurora.com/aurora-backend/internal/features/user/service"
	userUseCase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type Container struct {
	Router         *gin.Engine
	UserController *userController.UserController
}

func Build() (*Container, error) {

	userRepoImpl := userPersistence.NewUserInMemoryRepository()
	passwordHasher := userSecurity.NewBcryptHasher()

	createUserUC := userUseCase.NewCreateUserUseCase(userRepoImpl, passwordHasher)
	getUserUC := userUseCase.NewGetUserByIDUseCase(userRepoImpl)

	userSvc := userService.NewUserService(createUserUC, getUserUC)

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