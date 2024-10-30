package routers

import (
	"go-crud-challenge/delivery/controllers"
	"go-crud-challenge/repositories"
	"go-crud-challenge/usecases"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup repository and use case
	repo := repositories.NewInMemoryPersonRepository()
	personUseCase := usecases.NewPersonUseCase(repo)
	personHandler := controllers.NewPersonHandler(personUseCase)

	router.GET("/person", personHandler.GetAllPersons)
	router.GET("/person/:id", personHandler.GetPersonByID)
	router.POST("/person", personHandler.CreatePerson)
	router.PUT("/person/:id", personHandler.UpdatePerson)
	router.DELETE("/person/:id", personHandler.DeletePerson)

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Resource not found",
			"path":  c.Request.URL.Path,
		})
	})

	return router
}
