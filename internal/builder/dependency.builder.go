package builder

import (
	"context"

	"github.com/gin-gonic/gin"

	// User imports
	userController "aurora.com/aurora-backend/internal/features/user/controller"
	userFactory "aurora.com/aurora-backend/internal/features/user/factory"
	userGateway "aurora.com/aurora-backend/internal/features/user/gateway/repository"
	userSecurity "aurora.com/aurora-backend/internal/features/user/gateway/security"
	userService "aurora.com/aurora-backend/internal/features/user/service"

	// Events imports
	eventsController "aurora.com/aurora-backend/internal/features/events/controller"
	eventsDomain "aurora.com/aurora-backend/internal/features/events/domain"
	eventsFactory "aurora.com/aurora-backend/internal/features/events/factory"
	eventsGateway "aurora.com/aurora-backend/internal/features/events/gateway"
	eventsService "aurora.com/aurora-backend/internal/features/events/service"

	"aurora.com/aurora-backend/internal/firebase"
)

type Container struct {
	Router         *gin.Engine
	UserController *userController.UserController
	FirebaseApp    *firebase.FirebaseApp

	// Event repositories
	EventController    *eventsController.EventController
	EventRepo           eventsDomain.EventRepository
	TicketLotRepo       eventsDomain.TicketLotRepository
	PurchasedTicketRepo eventsDomain.PurchasedTicketRepository
	OrderRepo           eventsDomain.OrderRepository
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
	userRepoImpl, err := userGateway.NewUserFirestoreRepository(fbApp)
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

	eventRepoImpl, err := eventsGateway.NewEventFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	eventUseCaseFactory := eventsFactory.NewUseCaseFactory(eventRepoImpl)
	eventSvc := eventsService.NewEventService(
		eventUseCaseFactory.CreateEvent,
		eventUseCaseFactory.FindByIDEvent,
		eventUseCaseFactory.ListAllEvent,
	)
	eventCtrl := eventsController.NewEventController(eventSvc)

	ticketLotRepo, err := eventsGateway.NewTicketLotFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	purchasedTicketRepo, err := eventsGateway.NewPurchasedTicketFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	orderRepo, err := eventsGateway.NewOrderFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

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
			protectedRoutes.Use(authMw)
			{
				userRoutes := protectedRoutes.Group("/users")
				{
					userRoutes.GET("/:id", userCtrl.GetUser)
					userRoutes.PUT("/:id", userCtrl.UpdateUser)
					userRoutes.DELETE("/:id", userCtrl.DeleteUser)
					userRoutes.PUT("/me/password", userCtrl.ChangePassword)
				}
				eventRoutes := protectedRoutes.Group("/events")
				{
					eventRoutes.POST("/", eventCtrl.CreateEvent)
					eventRoutes.GET("/:id", eventCtrl.GetEvent)
					eventRoutes.GET("/", eventCtrl.ListEvents)
					
				}
			}
		}
	}

	return &Container{
		Router:              router,
		UserController:      userCtrl,
		FirebaseApp:         fbApp,
		EventRepo:           eventRepoImpl,
		TicketLotRepo:       ticketLotRepo,
		PurchasedTicketRepo: purchasedTicketRepo,
		OrderRepo:           orderRepo,
	}, nil
}