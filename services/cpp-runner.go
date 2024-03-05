package runner

import (
	// "io"
	// "os"
	"fmt"
	"os/exec"
	// db "github.com/sudeep162002/ims-go-backend/db"
	// "github.com/sudeep162002/ims-go-backend/middleware"
	// user_routes "github.com/sudeep162002/ims-go-backend/routes"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/client"
	// "github.com/docker/docker/pkg/stdcopy"
	// "github.com/gin-gonic/gin"
)

func CppRunner() {
	fmt.Println("inside handel cpp runner function")
	// Command to execute docker-compose up inside the container
	cmd := exec.Command("docker-compose", "-f", "./languge-container/docker-compose.yml", "up")

	// Create a context with cancellation support
	// ctx := context.Background()

	// Attach standard output and error streams
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	fmt.Println(stdout)
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	defer stderr.Close()

	// Start the command
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Copy the command output to the response writer
	// go io.Copy(c.Writer, stdout)
	// go io.Copy(c.Writer, stderr)

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		panic(err)
	}

}
