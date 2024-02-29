package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// db "github.com/sudeep162002/ims-go-backend/db"
	// "github.com/sudeep162002/ims-go-backend/middleware"
	runner "codeRunner/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/client"
	// "github.com/docker/docker/pkg/stdcopy"
)

func main() {
	r := gin.New()
	//env

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// make db connection
	// db.Initialize()

	// cors for attack protection

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowCredentials: true,          // Access-Control-Allow-Credentials: true
	}))

	//global  Middleware
	// r.Use(middleware.LoggerMiddleware())

	// Setup routes
	// user_routes.SetupHelloRoutes(r)

	//welcome route
	r.GET("/", func(c *gin.Context) {
		// Create a JSON object
		response := gin.H{
			"message": "welcome to code runner",
		}

		fmt.Println(response)
		// Return the JSON response with a custom status code
		c.JSON(http.StatusOK, response)
	})

	// health check route
	r.GET("/cont-c", runner.CRunner)
	r.GET("/cont-cpp", runner.CppRunner)
	r.GET("/cont-java", runner.JavaRunner)
	r.GET("/cont-golang", runner.GoRunner)

	// Run the server on port 3000
	r.Run(":" + os.Getenv("PORT"))
}
