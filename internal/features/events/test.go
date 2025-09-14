package events

import (
	"context"
	"fmt"
	"log"
	"time"

	"aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/features/events/gateway"
	"aurora.com/aurora-backend/internal/firebase"
)

// TestRepositories - Função para testar todos os repositórios do módulo de eventos
func TestRepositories() {
    ctx := context.Background()

    // Inicializa o Firebase
    fbApp, err := firebase.NewFirebaseApp(ctx)
    if err != nil {
        log.Fatalf("Erro ao inicializar Firebase: %v", err)
    }

    // Cria os repositórios
    eventRepo, err := gateway.NewEventFirestoreRepository(fbApp)
    if err != nil {
        log.Fatalf("Erro ao criar EventRepository: %v", err)
    }

    ticketLotRepo, err := gateway.NewTicketLotFirestoreRepository(fbApp)
    if err != nil {
        log.Fatalf("Erro ao criar TicketLotRepository: %v", err)
    }

    purchasedTicketRepo, err := gateway.NewPurchasedTicketFirestoreRepository(fbApp)
    if err != nil {
        log.Fatalf("Erro ao criar PurchasedTicketRepository: %v", err)
    }

    orderRepo, err := gateway.NewOrderFirestoreRepository(fbApp)
    if err != nil {
        log.Fatalf("Erro ao criar OrderRepository: %v", err)
    }

    fmt.Println("🔄 Iniciando testes dos repositórios...")

    // Testa Event Repository
    testEventRepository(ctx, eventRepo)
    
    // Testa TicketLot Repository
    testTicketLotRepository(ctx, ticketLotRepo)
    
    // Testa Order Repository
    testOrderRepository(ctx, orderRepo)
    
    // Testa PurchasedTicket Repository
    testPurchasedTicketRepository(ctx, purchasedTicketRepo)

    fmt.Println("✅ Todos os testes concluídos!")
}

func testEventRepository(ctx context.Context, repo domain.EventRepository) {
    fmt.Println("\n📅 Testando EventRepository...")

    // Cria um evento de teste
    event := &domain.Event{
        ID:          "test-event-001",
        Title:       "Show de Rock Teste",
        Description: "Um evento incrível para testar nosso sistema",
        Location:    "Arena Teste, São Paulo",
        CreatedAt:   time.Now(),
    }

    // Salva o evento
    savedEvent, err := repo.Save(ctx, event)
    if err != nil {
        log.Printf("❌ Erro ao salvar evento: %v", err)
        return
    }
    fmt.Printf("✅ Evento salvo: %s\n", savedEvent.Title)

    // Busca o evento por ID
    foundEvent, err := repo.FindByID(ctx, event.ID)
    if err != nil {
        log.Printf("❌ Erro ao buscar evento: %v", err)
        return
    }
    fmt.Printf("✅ Evento encontrado: %s\n", foundEvent.Title)

    // Lista todos os eventos
    events, err := repo.ListAll(ctx)
    if err != nil {
        log.Printf("❌ Erro ao listar eventos: %v", err)
        return
    }
    fmt.Printf("✅ Total de eventos: %d\n", len(events))

    // Busca por título
    eventsByTitle, err := repo.FindByTitle(ctx, event.Title)
    if err != nil {
        log.Printf("❌ Erro ao buscar por título: %v", err)
        return
    }
    fmt.Printf("✅ Eventos encontrados por título: %d\n", len(eventsByTitle))
}

func testTicketLotRepository(ctx context.Context, repo domain.TicketLotRepository) {
    fmt.Println("\n🎫 Testando TicketLotRepository...")

    // Cria um lote de ingressos
    ticketLot := &domain.TicketLot{
        ID:                "test-lot-001",
        EventID:           "test-event-001",
        Name:              "Pista Premium",
        Price:             15000, // R$ 150,00 em centavos
        TotalQuantity:     100,
        AvailableQuantity: 100,
    }

    // Salva o lote
    savedLot, err := repo.Save(ctx, ticketLot)
    if err != nil {
        log.Printf("❌ Erro ao salvar lote: %v", err)
        return
    }
    fmt.Printf("✅ Lote salvo: %s - R$ %.2f\n", savedLot.Name, float64(savedLot.Price)/100)

    // Busca o lote por ID
    foundLot, err := repo.FindByID(ctx, ticketLot.ID)
    if err != nil {
        log.Printf("❌ Erro ao buscar lote: %v", err)
        return
    }
    fmt.Printf("✅ Lote encontrado: %s\n", foundLot.Name)

    // Lista lotes por evento
    lots, err := repo.ListByEventID(ctx, "test-event-001")
    if err != nil {
        log.Printf("❌ Erro ao listar lotes: %v", err)
        return
    }
    fmt.Printf("✅ Lotes do evento: %d\n", len(lots))

    // Testa decrementar quantidade
    err = repo.DecrementAvailableQuantity(ctx, ticketLot.ID, 5)
    if err != nil {
        log.Printf("❌ Erro ao decrementar quantidade: %v", err)
        return
    }
    fmt.Printf("✅ Quantidade decrementada com sucesso\n")

    // Verifica se a quantidade foi decrementada
    updatedLot, err := repo.FindByID(ctx, ticketLot.ID)
    if err != nil {
        log.Printf("❌ Erro ao buscar lote atualizado: %v", err)
        return
    }
    fmt.Printf("✅ Quantidade disponível após decremento: %d\n", updatedLot.AvailableQuantity)
}

func testOrderRepository(ctx context.Context, repo domain.OrderRepository) {
    fmt.Println("\n🛒 Testando OrderRepository...")

    // Cria um pedido
    order := &domain.Order{
        ID:        "test-order-001",
        UserID:    "user-123",
        EventID:   "test-event-001",
        Status:    domain.ORDER_PENDING,
        TotalAmount:     30000, // R$ 300,00 em centavos
    }

    // Salva o pedido
    savedOrder, err := repo.Save(ctx, order)
    if err != nil {
        log.Printf("❌ Erro ao salvar pedido: %v", err)
        return
    }
    fmt.Printf("✅ Pedido salvo: %s - R$ %.2f\n", savedOrder.ID, float64(savedOrder.TotalAmount)/100)

    // Busca o pedido por ID
    foundOrder, err := repo.FindByID(ctx, order.ID)
    if err != nil {
        log.Printf("❌ Erro ao buscar pedido: %v", err)
        return
    }
    fmt.Printf("✅ Pedido encontrado: %s\n", foundOrder.ID)

    // Lista pedidos por usuário
    orders, err := repo.ListByUserID(ctx, "user-123")
    if err != nil {
        log.Printf("❌ Erro ao listar pedidos: %v", err)
        return
    }
    fmt.Printf("✅ Pedidos do usuário: %d\n", len(orders))
}

func testPurchasedTicketRepository(ctx context.Context, repo domain.PurchasedTicketRepository) {
    fmt.Println("\n🎟️ Testando PurchasedTicketRepository...")

    // Cria um ingresso comprado
    ticket := &domain.PurchasedTicket{
        ID:         "test-ticket-001",
        UserID:     "user-123",
        OrderID:    "test-order-001",
        TicketLotID: "test-lot-001",
        EventID:    "test-event-001",
        Status:     domain.TICKET_VALID,
    }

    // Salva o ingresso
    savedTicket, err := repo.Save(ctx, ticket)
    if err != nil {
        log.Printf("❌ Erro ao salvar ingresso: %v", err)
        return
    }
    fmt.Printf("✅ Ingresso salvo: %s - R$ %.2f\n", savedTicket.ID, float64(savedTicket.PurchasePrice)/100)

    // Busca o ingresso por ID
    foundTicket, err := repo.FindByID(ctx, ticket.ID)
    if err != nil {
        log.Printf("❌ Erro ao buscar ingresso: %v", err)
        return
    }
    fmt.Printf("✅ Ingresso encontrado: %s\n", foundTicket.ID)

    // Lista ingressos por usuário
    tickets, err := repo.ListByUserID(ctx, "user-123")
    if err != nil {
        log.Printf("❌ Erro ao listar ingressos do usuário: %v", err)
        return
    }
    fmt.Printf("✅ Ingressos do usuário: %d\n", len(tickets))

    // Lista ingressos por pedido
    orderTickets, err := repo.ListByOrderID(ctx, "test-order-001")
    if err != nil {
        log.Printf("❌ Erro ao listar ingressos do pedido: %v", err)
        return
    }
    fmt.Printf("✅ Ingressos do pedido: %d\n", len(orderTickets))

    // Atualiza status do ingresso
    err = repo.UpdateStatus(ctx, ticket.ID, domain.TICKET_USED)
    if err != nil {
        log.Printf("❌ Erro ao atualizar status: %v", err)
        return
    }
    fmt.Printf("✅ Status do ingresso atualizado\n")
}