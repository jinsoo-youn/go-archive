package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinsoo-youn/go-pkg/logger"
	"log"
	"net/http"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/jinsoo-youn/go-archive/banking/docs"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8000
// @BasePath    /
func Start() {

	//sanityCheck()

	// HTTP Server
	handler := gin.New()
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	//Swagger

	handler.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER"))
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	handler.GET("/history", history)

	// starting server
	//address := os.Getenv("SERVER_ADDRESS")
	//port := os.Getenv("SERVER_PORT")
	address := "localhost"
	port := "8000"
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), handler))

}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    history
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500
// @Router      /history [get]
func history(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "history is OK",
	})
}
