package builder

import (
	"github.com/gin-gonic/gin"

	// Importações atualizadas com os novos caminhos
	userController "aurora.com/aurora-backend/internal/features/user/controller"
	userPersistence "aurora.com/aurora-backend/internal/features/user/gateway/repository" // Implementação
	userService "aurora.com/aurora-backend/internal/features/user/service"
	userUseCase "aurora.com/aurora-backend/internal/features/user/use-case"
)

type Container struct {
	Router         *gin.Engine
	UserController *userController.UserController
}

func Build() (*Container, error) {
	// Camada de Persistência (implementação do repositório)
	// Chamamos a função do novo pacote 'persistence'
	userRepoImpl := userPersistence.NewUserInMemoryRepository()

	// Camada de Casos de Uso (agora recebe a implementação que satisfaz a nova interface)
	createUserUC := userUseCase.NewCreateUserUseCase(userRepoImpl)
	getUserUC := userUseCase.NewGetUserByIDUseCase(userRepoImpl)

	// Camada de Serviço (recebe os casos de uso)
	userSvc := userService.NewUserService(createUserUC, getUserUC)

	// Camada de Controller (recebe o serviço)
	userCtrl := userController.NewUserController(userSvc)

	// Configuração do Roteador (sem alterações)
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