package main

import (
	runner "codeRunner/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	// "os/signal"

	// db "github.com/sudeep162002/ims-go-backend/db"
	// "github.com/sudeep162002/ims-go-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type Output struct {
	ID         int `gorm:"primary_key"`
	ReqID      int
	Username   string
	CodeOutput string
}

type Message struct {
	ID       string `json:"id"`
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

func handleCode(codeReqJSON CodeRequest) {
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
		// runner.CppRunner()

	}
}

func handleJavaCode(codeReqJSON CodeRequest) {
	fmt.Println("Inside handle java function")

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

	// Publish JSON string to Kafka topic
	err = producer.WriteMessages(context.Background(), kafka.Message{
		Topic: "java-code", // Specify the topic to publish messages to
		Value: codeReqJSONStr,
	})
	if err != nil {
		log.Fatalf("Error publishing message to java-code Kafka: %v", err)
	}

	log.Println("Message published to java-code Kafka")

	if runningCppContainer == 0 {
		runner.JavaRunner()
		runningJavaContainer += 1
		// runner.CppRunner()

	}
}

func handleGoCode(codeReqJSON CodeRequest) {
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

	// Publish JSON string to Kafka topic
	err = producer.WriteMessages(context.Background(), kafka.Message{
		Topic: "go-code", // Specify the topic to publish messages to
		Value: codeReqJSONStr,
	})
	if err != nil {
		log.Fatalf("Error publishing message to Kafka: %v", err)
	}

	log.Println("Message published to Kafka")

	if runningCppContainer == 0 {
		runner.GoRunner()
		runningGolangContainer += 1
		// runner.CppRunner()

	}
}

func handleCCode(codeReqJSON CodeRequest) {
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

	// Publish JSON string to Kafka topic
	err = producer.WriteMessages(context.Background(), kafka.Message{
		Topic: "c-code", // Specify the topic to publish messages to
		Value: codeReqJSONStr,
	})
	if err != nil {
		log.Fatalf("Error publishing message to Kafka: %v", err)
	}

	log.Println("Message published to Kafka")

	if runningCppContainer == 0 {
		runner.CRunner()
		runningCContainer += 1
		// runner.CppRunner()

	}
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

func insertData(db *gorm.DB) {
	// Seed random number generator
	fmt.Println("this is insert data function")
	rand.Seed(time.Now().UnixNano())

	config := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "code-output",
		GroupID:  "kafka1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	// Create Kafka reader
	reader := kafka.NewReader(config)
	defer reader.Close()

	// Create context to cancel consumer on termination signal
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle termination signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-signals:
			fmt.Println("Received termination signal, shutting down consumer...")
			cancel()
		}
	}()

	// Start consuming messages
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, stopping consumer...")
			return
		default:
			message, err := reader.ReadMessage(ctx)
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				continue
			}
			// fmt.Printf("Received message: Key=%s, Value=%s, Partition=%d, Offset=%d\n",
			// 	string(message.Key), string(message.Value), message.Partition, message.Offset)
			var msg Message

			if err := json.Unmarshal(message.Value, &msg); err != nil {
				fmt.Printf("Error decoding JSON: %v\n", err)
				continue
			}
			fmt.Println(msg)

			reqID, err := strconv.Atoi(msg.ID)
			if err != nil {
				fmt.Println(err)
			}
			username := msg.UserName
			codeOutput := msg.Output

			output := Output{
				ReqID:      reqID,
				Username:   username,
				CodeOutput: codeOutput,
			}
			if err := db.Create(&output).Error; err != nil {
				log.Fatal(err)
			}
			fmt.Println("data inserted successfully.")
			// fmt.Println(message.Value)
			// Simulate some processing time
			time.Sleep(1 * time.Second)
		}
	}
}

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

	dsn := "root:Sudeep@16@tcp(localhost:3306)/codeDatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Auto migrate the Output table
	if err := db.AutoMigrate(&Output{}); err != nil {
		log.Fatal(err)
	}

	err = createKafkaTopic(producer, "code-output")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}

	// Create Kafka topic "cpp-code" if not already created
	err = createKafkaTopic(producer, "cpp-code")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}
	err = createKafkaTopic(producer, "java-code")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}
	err = createKafkaTopic(producer, "go-code")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}
	err = createKafkaTopic(producer, "c-code")
	if err != nil {
		log.Fatalf("Error creating Kafka topic: %v", err)
	}

	go insertData(db)

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

	r.GET("/:reqId/:username", func(c *gin.Context) {
		// Extract the req and username parameters from the URL
		reqId := c.Param("reqId")
		username := c.Param("username")

		var output Output
		if err := db.Where("req_id = ? AND username = ?", reqId, username).First(&output).Error; err != nil {
			fmt.Println("still in processing..........")
			response := gin.H{
				"req_id":   reqId,
				"username": username,
				"output":   "still in processing..........",
				"message":  "your code is being processing in backend please wait",
			}

			// Return the JSON response with a custom status code
			c.JSON(http.StatusOK, response)

			return
		}

		// Create a JSON response
		response := gin.H{
			"req_id":   output.ReqID,
			"username": output.Username,
			"output":   output.CodeOutput,
			"message":  "Database query successful",
		}

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
				handleCode(codeReq)
				// runningCppContainer += 1
			case "c":
				handleCCode(codeReq)
				// runningCContainer += 1
			case "golang":
				handleGoCode(codeReq)
				// runningGolangContainer += 1
			case "java":
				handleJavaCode(codeReq)
				// runningJavaContainer += 1

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
