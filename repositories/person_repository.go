package repositories

import (
	"context"
	"errors"
	"sync"

	"go-crud-challenge/domain"
)

type InMemoryPersonRepository struct {
	persons map[string]domain.Person
	mutex   sync.RWMutex
}

func NewInMemoryPersonRepository() *InMemoryPersonRepository {
	return &InMemoryPersonRepository{
		persons: make(map[string]domain.Person),
	}
}

func (r *InMemoryPersonRepository) GetAllPersons(ctx context.Context) ([]domain.Person, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	var result []domain.Person
	for _, person := range r.persons {
		result = append(result, person)
	}
	return result, nil
}

func (r *InMemoryPersonRepository) GetPersonByID(ctx context.Context, id string) (domain.Person, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	person, exists := r.persons[id]
	if !exists {
		return domain.Person{}, errors.New("person not found")
	}
	return person, nil
}

func (r *InMemoryPersonRepository) CreatePerson(ctx context.Context, person domain.Person) (domain.Person, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.persons[person.ID] = person
	return person, nil
}

func (r *InMemoryPersonRepository) UpdatePerson(ctx context.Context, id string, person domain.Person) (domain.Person, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.persons[id]; !exists {
		return domain.Person{}, errors.New("person not found")
	}
	r.persons[id] = person
	return person, nil
}

func (r *InMemoryPersonRepository) DeletePerson(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.persons[id]; !exists {
		return errors.New("person not found")
	}
	delete(r.persons, id)
	return nil
}
