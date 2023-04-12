package routers

import (
	"project-1-chapter-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetAllBook)
	router.GET("/books/:bookID", controllers.GetBookById)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
