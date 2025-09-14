package gateway

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"aurora.com/aurora-backend/internal/features/events/domain"
	"aurora.com/aurora-backend/internal/firebase"
)


const eventCollection = "events"

type EventFirestoreRepository struct {
	client *firestore.Client
}

// NewEventFirestoreRepository cria uma nova instância do repositório do Firestore.
func NewEventFirestoreRepository(fbApp *firebase.FirebaseApp) (domain.EventRepository, error) {
	client, err := fbApp.Firestore(context.Background())
		if err != nil {
		log.Fatalf("Falha ao criar cliente do Firestore: %v", err)
		return nil, err
	}
	return &EventFirestoreRepository{client: client}, nil
}

func (r *EventFirestoreRepository) Save(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	_, err := r.client.Collection(eventCollection).Doc(event.ID).Set(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventFirestoreRepository) FindByID(ctx context.Context, id string) (*domain.Event, error) {
	docSnap, err := r.client.Collection(eventCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var event domain.Event
	if err := docSnap.DataTo(&event); err != nil {
		return nil, err
	}
	event.ID = docSnap.Ref.ID
	return &event, nil
}

func (r *EventFirestoreRepository) Update(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	_, err := r.client.Collection(eventCollection).Doc(event.ID).Set(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventFirestoreRepository) Delete(ctx context.Context, id string) error {
    _, err := r.client.Collection(eventCollection).Doc(id).Delete(ctx)
    return err
}

func (r *EventFirestoreRepository) ListAll(ctx context.Context) ([]*domain.Event, error) {
	iter := r.client.Collection(eventCollection).Documents(ctx)
	defer iter.Stop()
	var events []*domain.Event
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}
		var event domain.Event
		if err := doc.DataTo(&event); err != nil {
			return nil, err
		}
		event.ID = doc.Ref.ID
		events = append(events, &event)
	}
	return events, nil
}

func (r *EventFirestoreRepository) FindByTitle(ctx context.Context, title string) ([]*domain.Event, error) {
	iter := r.client.Collection(eventCollection).Where("Title", "==", title).Documents(ctx)
	defer iter.Stop()
	var events []*domain.Event
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}	
		var event domain.Event
		if err := doc.DataTo(&event); err != nil {
			return nil, err
		}
		event.ID = doc.Ref.ID
		events = append(events, &event)
	}
	return events, nil
}