code deparser

const deparser = inputString
      .replace(/\\n/g, '\n')       // Replace "\\n" with actual line breaks
      .replace(/^    /gm, '')      // Remove leading indentation
      .replace(/\\"/g, '"');       // Replace escaped double quotes with double quotes



code parser

const code = normalCode
    .replace(/"/g, '\\"') // Replace all double quotes with escaped double quotes
    .split('\n')          // Split the code into lines
    .map(line => `    ${line}`) // Indent each line
    .join('\\n\\n');      // Join the lines with "\n\n"

console.log(`const code = "${code}";`);



command for running kafka container

docker compose up



command for input data to a topic from cmd 

kafka-console-producer.sh --broker-list localhost:9092 --topic code


command for checking what is inside a tpic

kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic code-output --from-beginning





sampel input
{"id":"69888","userName":"sudeep162002","output":"#include <iostream>\\n\\n    int main() {\\n\\n        std::cout << \"hello, World!\" << std::endl;\\n\\n        return 0;\\n    }"}



kafka internal ip: kafka:29092
kafka external ip:  localhost:9092


use kafkacat container command to check status of queue    `kafkacat -b kafka:29092 -t code-output -C`


code to be used when calling docker image from docker hub 

`
package runner

import (
	"context"
	"io"
	"os"

	// db "github.com/sudeep162002/ims-go-backend/db"
	// "github.com/sudeep162002/ims-go-backend/middleware"
	// user_routes "github.com/sudeep162002/ims-go-backend/routes"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/gin-gonic/gin"
)

func GoRunner(c *gin.Context) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

}

`



always wrap your java code inside class named `HelloWorld`