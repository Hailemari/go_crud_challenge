package usecases

import (
	"context"
	"go-crud-challenge/domain"

	"github.com/google/uuid"
)

type personUseCase struct {
	repo domain.IPersonRepository
}

func NewPersonUseCase(repo domain.IPersonRepository) domain.IPersonUsecase {
	return &personUseCase{repo: repo}
}

func (uc *personUseCase) GetAllPersons(ctx context.Context) ([]domain.Person, error) {
	return uc.repo.GetAllPersons(ctx)
}

func (uc *personUseCase) GetPersonByID(ctx context.Context, id string) (domain.Person, error) {
	return uc.repo.GetPersonByID(ctx, id)
}

func (uc *personUseCase) CreatePerson(ctx context.Context, person domain.Person) (domain.Person, error) {
	// Generate a new UUID for the person
	person.ID = uuid.NewString()

	// Simulate storing the person in an in-memory database
	// Assume uc.repo is the repository that handles storage
	createdPerson, err := uc.repo.CreatePerson(ctx, person)
	if err != nil {
		return domain.Person{}, err // Return an empty Person and the error if creation fails
	}

	return createdPerson, nil // Return the created Person object
}

func (uc *personUseCase) UpdatePerson(ctx context.Context, id string, person domain.Person) (domain.Person, error) {
	return uc.repo.UpdatePerson(ctx, id, person)
}

func (uc *personUseCase) DeletePerson(ctx context.Context, id string) error {
	return uc.repo.DeletePerson(ctx, id)
}
