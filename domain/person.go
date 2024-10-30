package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name" binding:"required"`
	Age     int      `json:"age" binding:"required,min=0"`
	Hobbies []string `json:"hobbies" binding:"required"`
}

type IPersonRepository interface {
	GetAllPersons(ctx context.Context) ([]Person, error)
	GetPersonByID(ctx context.Context, id string) (Person, error)
	CreatePerson(ctx context.Context, person Person) (Person, error)
	UpdatePerson(ctx context.Context, id string, person Person) (Person, error)
	DeletePerson(ctx context.Context, id string) error
}

type IPersonUsecase interface {
	GetAllPersons(ctx context.Context) ([]Person, error)
	GetPersonByID(ctx context.Context, id string) (Person, error)
	CreatePerson(ctx context.Context, person Person) (Person, error)
	UpdatePerson(ctx context.Context, id string, person Person) (Person, error)
	DeletePerson(ctx context.Context, id string) error
}

type IPersonController interface {
	GetAllPersons(c *gin.Context)
	GetPersonByID(c *gin.Context)
	CreatePerson(c *gin.Context)
	UpdatePerson(c *gin.Context)
	DeletePerson(c *gin.Context)
}
