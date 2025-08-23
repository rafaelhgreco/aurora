package builder

import (
	"context"

	"github.com/gin-gonic/gin"

	userController "aurora.com/aurora-backend/internal/features/user/controller"
	userFactory "aurora.com/aurora-backend/internal/features/user/factory"
	userPersistence "aurora.com/aurora-backend/internal/features/user/gateway/repository"
	userSecurity "aurora.com/aurora-backend/internal/features/user/gateway/security"
	userService "aurora.com/aurora-backend/internal/features/user/service"
	"aurora.com/aurora-backend/internal/firebase"
)

type Container struct {
	Router         *gin.Engine
	UserController *userController.UserController
	FirebaseApp    *firebase.FirebaseApp
}

func Build() (*Container, error) {
	ctx := context.Background()
	fbApp, err := firebase.NewFirebaseApp(ctx)
    if err != nil {
        return nil, err
    }


	authClient, err := fbApp.App.Auth(ctx)
	if err != nil {
		return nil, err
	}

	authGateway := userSecurity.NewFirebaseAuthGateway(authClient)

	userRepoImpl, err := userPersistence.NewUserFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}
	passwordHasher := userSecurity.NewBcryptHasher()
	
	userUseCaseFactory := userFactory.NewUseCaseFactory(userRepoImpl, passwordHasher, authGateway)

	userSvc := userService.NewUserService(
		userUseCaseFactory.CreateUser,
		userUseCaseFactory.GetUser,
		userUseCaseFactory.UpdateUser,
		userUseCaseFactory.DeleteUser,
		userUseCaseFactory.LoginUser,
		userUseCaseFactory.ChangePassword,
	)

	userCtrl := userController.NewUserController(userSvc)

	router := gin.Default()
	authMw := userSecurity.AuthMiddleware(authClient)

	{
    v1 := router.Group("/v1")
    {
        // Rotas públicas
        v1.POST("/auth/login", userCtrl.Login)
        v1.POST("/users", userCtrl.CreateUser)

        // Rotas protegidas
        protectedRoutes := v1.Group("/")
        protectedRoutes.Use(authMw) // APLICA O MIDDLEWARE AQUI
        {
            userRoutes := protectedRoutes.Group("/users")
            {
                userRoutes.GET("/:id", userCtrl.GetUser) // Agora protegido
                userRoutes.PUT("/:id", userCtrl.UpdateUser) // Agora protegido
                userRoutes.DELETE("/:id", userCtrl.DeleteUser) // Agora protegido
                
                // Rota mais segura para alterar a senha
                userRoutes.PUT("/me/password", userCtrl.ChangePassword) 
            }
        }
    }
	}

	return &Container{
		Router:         router,
		UserController: userCtrl,
		FirebaseApp:    fbApp,
	}, nil
}