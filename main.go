package main

import (
	"earthly-elixirs-api/controller"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(CORSConfig()))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "wasssssup"})
	})
	return router
}

func ServeApplication() {
	router := SetupRouter()
	controller.RegisterCheckoutRoutes(router)
	router.Run(":8000")
	host := os.Getenv("HOST")

	hostUrl := fmt.Sprintf("%s/docs/swagger.json", host)

	url := ginSwagger.URL(hostUrl)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Println("Server running on port 8000")
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("unable to load environment")
		return
	}

	ServeApplication()

}
