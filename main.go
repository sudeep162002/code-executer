package main

import (
	runner "codeRunner/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// "os/signal"

	// db "github.com/sudeep162002/ims-go-backend/db"
	// "github.com/sudeep162002/ims-go-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	// "github.com/Shopify/sarama"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/client"
	// "github.com/docker/docker/pkg/stdcopy"
)

type CodeRequest struct {
	ID       string `json:"id"`
	Languge  string `json:"languge"`
	UserName string `json:"userName"`
	Output   string `json:"output"`
}

var producer *kafka.Writer

func initKafkaProducer(brokers []string) {
	// Initialize Kafka producer
	producer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        "",
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 1 * time.Millisecond,
	})
}

func createKafkaTopic(producer *kafka.Writer, topic string) error {
	// Define Kafka config
	config := kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	// Create Kafka connection
	conn, err := config.Dial("tcp", "localhost:9092")
	if err != nil {
		return fmt.Errorf("error connecting to Kafka broker: %v", err)
	}
	defer conn.Close()

	// Create topic
	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	if err != nil {
		return fmt.Errorf("error creating Kafka topic: %v", err)
	}

	return nil
}

var runningCppContainer int = 0
var runningJavaContainer int = 0
var runningGolangContainer int = 0
var runningCContainer int = 0

func handelCpp(codeReqJSON CodeRequest) {
	fmt.Println("Inside handle cpp function")

	// Convert CodeRequest struct to JSON string
	codeReqJSONStr, err := json.Marshal(codeReqJSON)
	if err != nil {
		log.Fatalf("Error marshalling CodeRequest to JSON: %v", err)
	}

	fmt.Println(string(codeReqJSONStr))

	// Initialize Kafka producer if not already initialized
	if producer == nil {
		brokers := []string{"localhost:9092"} // Define Kafka brokers
		initKafkaProducer(brokers)            // Initialize Kafka producer without specifying topic
	}

	// Create Kafka topic "cpp-code" if not already created
	err = createKafkaTopic(producer, "cpp-code")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}
	err = createKafkaTopic(producer, "code-output")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}

	// Publish JSON string to Kafka topic
	err = producer.WriteMessages(context.Background(), kafka.Message{
		Topic: "cpp-code", // Specify the topic to publish messages to
		Value: codeReqJSONStr,
	})
	if err != nil {
		log.Fatalf("Error publishing message to Kafka: %v", err)
	}

	log.Println("Message published to Kafka")

	if runningCppContainer == 0 {
		runner.CppRunner()
		runningCppContainer += 1
	}

	// Your code to handle the C++ code request goes here
}

// func handelCpp(codeReq CodeRequest) {

// 	codeReqJSON, err := json.Marshal(codeReq)
// 	if err != nil {
// 		log.Fatalf("Error marshalling CodeRequest to JSON: %v", err)
// 	}

// 	// Convert JSON byte slice to string
// 	codeReqString := string(codeReqJSON)

// 	// Your code to handle the C++ code request goes here
// }

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
	r.POST("/run-code", func(c *gin.Context) {
		var codeReq CodeRequest

		// Bind the JSON body to the CodeRequest struct
		if err := c.BindJSON(&codeReq); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		// Extract the output string from the request
		lang := codeReq.Languge

		// Process the output string concurrently using goroutines
		go func() {
			switch lang {
			case "cpp":
				handelCpp(codeReq)
			case "c":
				handelCpp(codeReq)
			case "golang":
				handelCpp(codeReq)
			case "java":
				handelCpp(codeReq)

			default:
				// Handle unknown language
				fmt.Println("Unknown language:", lang)
			}

		}()

		// Return the response immediately
		c.JSON(200, gin.H{"message": "Output received"})
	})
	// r.GET("/cont-cpp", runner.CppRunner)
	// r.GET("/cont-java", runner.JavaRunner)
	// r.GET("/cont-golang", runner.GoRunner)

	// Run the server on port 3000
	r.Run(":" + os.Getenv("PORT"))
}
