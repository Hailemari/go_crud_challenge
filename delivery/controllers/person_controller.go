package controllers

import (
	"net/http"

	"go-crud-challenge/domain"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	useCase domain.IPersonUsecase
}

func NewPersonHandler(useCase domain.IPersonUsecase) *PersonHandler {
	return &PersonHandler{useCase: useCase}
}

func (h *PersonHandler) GetAllPersons(c *gin.Context) {
	persons, err := h.useCase.GetAllPersons(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve persons"})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetPersonByID(c *gin.Context) {
	id := c.Param("id")
	person, err := h.useCase.GetPersonByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var newPerson domain.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Validate required fields
	if newPerson.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": "Name is required",
		})
		return
	}

	if newPerson.Age < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": "Age must be non-negative",
		})
		return
	}

	if newPerson.Hobbies == nil {
		newPerson.Hobbies = []string{} // Initialize empty array if nil
	}

	person, err := h.useCase.CreatePerson(c, newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create person",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var updatedPerson domain.Person
	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ID must be preserved in updatedPerson for the update operation
	updatedPerson.ID = id

	person, err := h.useCase.UpdatePerson(c, id, updatedPerson)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := h.useCase.DeletePerson(c, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
