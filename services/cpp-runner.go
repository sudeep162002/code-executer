package runner

import (
	// "io"
	// "os"
	"context"
	"fmt"
	"io"

	// "io"

	// "fmt"
	"os"
	// "os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	// db "github.com/sudeep162002/ims-go-backend/db"
	// "github.com/sudeep162002/ims-go-backend/middleware"
	// user_routes "github.com/sudeep162002/ims-go-backend/routes"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/client"
	// "github.com/docker/docker/pkg/stdcopy"
	// "github.com/gin-gonic/gin"
)

// func CppRunner() {
// 	fmt.Println("inside handel cpp runner function")
// 	// Command to execute docker-compose up inside the container
// 	cmd := exec.Command("docker-compose", "-f", "./languge-container/docker-compose.yml", "up")

// 	// Create a context with cancellation support
// 	// ctx := context.Background()

// 	// Attach standard output and error streams
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(stdout)
// 	defer stdout.Close()

// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stderr.Close()

// 	// Start the command
// 	if err := cmd.Start(); err != nil {
// 		panic(err)
// 	}

// 	// Copy the command output to the response writer
// 	// go io.Copy(c.Writer, stdout)
// 	// go io.Copy(c.Writer, stderr)

// 	// Wait for the command to finish
// 	if err := cmd.Wait(); err != nil {
// 		panic(err)
// 	}

// }

func fetchDockerImage() error {
	ctx := context.Background()

	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	// Pull the sudeep162002/cpp-code-runner:latest image
	reader, err := cli.ImagePull(ctx, "sudeep162002/cpp-code-runner:latest", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()

	// Copy the output of the pull operation to stdout
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return err
	}

	fmt.Println("Image fetched successfully.")
	return nil
}

func CppRunner() {

	fetchDockerImage()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Failed to create Docker client: %v\n", err)
		return
	}
	defer cli.Close()

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "sudeep162002/cpp-code-runner:latest",
		Cmd:   []string{"npm", "start"},
	}, &container.HostConfig{
		NetworkMode: "notification_default",
	}, nil, nil, "")
	if err != nil {
		fmt.Printf("Failed to create container: %v\n", err)
		return
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("Failed to start container: %v\n", err)
		return
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			fmt.Printf("Error waiting for container: %v\n", err)
			return
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		fmt.Printf("Failed to get container logs: %v\n", err)
		return
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	fmt.Println("Container executed successfully")
}
