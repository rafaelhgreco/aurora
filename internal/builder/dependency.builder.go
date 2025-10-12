package builder

import (
	"context"

	"github.com/gin-gonic/gin"

	// User imports
	userController "aurora.com/aurora-backend/internal/features/user/controller"
	userFactory "aurora.com/aurora-backend/internal/features/user/factory"
	userGateway "aurora.com/aurora-backend/internal/features/user/gateway/repository"
	userSecurity "aurora.com/aurora-backend/internal/features/user/gateway/security"

	// Events imports
	eventsController "aurora.com/aurora-backend/internal/features/events/controller"
	eventsDomain "aurora.com/aurora-backend/internal/features/events/domain"
	eventsFactory "aurora.com/aurora-backend/internal/features/events/factory"
	eventsGateway "aurora.com/aurora-backend/internal/features/events/gateway"

	// Tickets imports
	ticketsController "aurora.com/aurora-backend/internal/features/tickets/controller"
	ticketsDomain "aurora.com/aurora-backend/internal/features/tickets/domain"
	ticketsFactory "aurora.com/aurora-backend/internal/features/tickets/factory"
	ticketsGateway "aurora.com/aurora-backend/internal/features/tickets/gateway"

	// Orders imports
	orderController "aurora.com/aurora-backend/internal/features/order/controller"
	orderDomain "aurora.com/aurora-backend/internal/features/order/domain"
	orderFactory "aurora.com/aurora-backend/internal/features/order/factory"
	orderGateway "aurora.com/aurora-backend/internal/features/order/gateway"

	"aurora.com/aurora-backend/internal/firebase"
)

type Container struct {
	Router         *gin.Engine
	UserController *userController.UserController
	FirebaseApp    *firebase.FirebaseApp

	// Event repositories
	EventController *eventsController.EventController
	EventRepo       eventsDomain.EventRepository

	// Ticket repositories
	TicketController    *ticketsController.TicketController
	PurchasedTicketRepo ticketsDomain.PurchasedTicketRepository

	// Order repositories]
	OrderController *orderController.OrderController
	OrderRepo       orderDomain.OrderRepository
}

func Build() (*Container, error) {
	ctx := context.Background()
	fbApp, err := firebase.NewFirebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	// User
	authClient, err := fbApp.App.Auth(ctx)
	if err != nil {
		return nil, err
	}

	authGateway := userSecurity.NewFirebaseAuthGateway(authClient)
	userRepoImpl, err := userGateway.NewUserFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	userUseCaseFactory := userFactory.NewUseCaseFactory(userRepoImpl, authGateway)
	userCtrl := userController.NewUserController(userUseCaseFactory.CreateUser, userUseCaseFactory.UpdateUser, userUseCaseFactory.GetUserByID, userUseCaseFactory.DeleteUser, userUseCaseFactory.LoginUser, userUseCaseFactory.ChangePassword)

	// Repositories Events - Tickets - Orders
	eventRepoImpl, err := eventsGateway.NewEventFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	purchasedTicketRepo, err := ticketsGateway.NewPurchasedTicketFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	orderRepo, err := orderGateway.NewOrderFirestoreRepository(fbApp)
	if err != nil {
		return nil, err
	}

	// Event Factory
	eventUseCaseFactory := eventsFactory.NewUseCaseFactory(eventRepoImpl)
	eventCtrl := eventsController.NewEventController(eventUseCaseFactory.CreateEvent, eventUseCaseFactory.FindByIDEvent, eventUseCaseFactory.ListAllEvent, eventUseCaseFactory.SoftDeleteEvent)
	// Ticket Factory
	ticketUseCaseFactory := ticketsFactory.NewUseCaseFactory(
		eventRepoImpl,
		orderRepo,
		purchasedTicketRepo,
	)
	ticketCtrl := ticketsController.NewTicketController(ticketUseCaseFactory.PurchaseTicket)

	// Order Factory
	orderUseCaseFactory := orderFactory.NewUseCaseFactory(orderRepo, eventRepoImpl)
	orderCtrl := orderController.NewOrderController(orderUseCaseFactory.CreateOrder)
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
					eventRoutes.PATCH("/:id/cancel", eventCtrl.SoftDeleteEvent)

				}
				ticketRoutes := protectedRoutes.Group("/tickets")
				{
					ticketRoutes.POST("/", ticketCtrl.CreateTicket)
				}
				orderRoutes := protectedRoutes.Group("/orders")
				{
					orderRoutes.POST("/", orderCtrl.CreateOrder)
				}
			}
		}
	}

	return &Container{
		Router:              router,
		UserController:      userCtrl,
		FirebaseApp:         fbApp,
		EventRepo:           eventRepoImpl,
		PurchasedTicketRepo: purchasedTicketRepo,
		OrderRepo:           orderRepo,
	}, nil
}
