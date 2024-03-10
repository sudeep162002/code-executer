## **Executioner: A Scalable Code Execution Platform for Everyone**


![diagram-export-07-03-2024-21_56_12](https://github.com/sudeep162002/code-executer/assets/80563363/96fcf291-3abc-4be9-bdc2-6b66c1c42ef8)



**Introduction**

Welcome to Executioner, your one-stop shop for seamless and efficient code execution! Inspired by platforms like LeetCode, we offer a **highly scalable and user-friendly** environment to execute your code with ease.

**Unmatched Scalability:**

- **Built for Performance:** Engineered to handle **millions of code execution requests concurrently**, ensuring smooth operation even under massive workloads.
- **Containerized with Docker:** Leverages containerization for **flexible deployment**, efficient resource utilization, and effortless scaling.
- **Master-Worker Architecture:** Employs a **Kafka-powered master-worker architecture** where the master distributes tasks to worker containers for parallel execution.

**Tech Stack and Tools:**

- **Golang:** Enables efficient container orchestration using [https://pkg.go.dev/runtime](https://pkg.go.dev/runtime).
- **TypeScript:** Drives worker containers, tailored for code execution with specific languages.
- **Node.js and Express:** Forms the foundation for the API gateway, facilitating interaction with the platform.
- **Go Gin:** Powers the API gateway, providing a high-performance framework for building robust APIs in Golang.
- **Kafka:** Acts as the central hub for message queuing, handling task distribution and ensuring efficient communication between components.
- **MySQL Database:** Stores code execution results and user data securely.
- **Server-Sent Events (SSE):** Enables real-time result delivery to the frontend, enhancing user experience.

**Multithreading for Maximum Efficiency:**

Maximizes CPU utilization by employing multithreading within worker containers, ensuring optimal performance for code execution.

**Why Golang?**

- **Efficient Docker Management:** Golang excels at handling the Docker engine daemon effectively, contributing to the platform's smooth operation.
- **Concurrency and Performance:** Well-suited for building concurrent and high-performance applications.

## Local Development Setup

**Prerequisites:**

- **Golang:** Download and install Golang from [https://go.dev/](https://go.dev/). Ensure you have set the `GOPATH` and `GOBIN` environment variables correctly.
- **Docker:** Download and install Docker from [https://www.docker.com/](https://www.docker.com/).

**Steps:**

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/sudeep162002/code-executer
   ```

2. **Define Port:**
   Create a file named `.env` in the base directory of your project. Add a line specifying the port you want to use for the API gateway:

   ```
   API_PORT=8080
   ```

   Replace `8080` with your desired port number.

3. **Install Dependencies:**
   Open a terminal in the project directory and run:

   ```bash
   go get
   ```

   This command downloads all the necessary Go dependencies listed in the `go.mod` file.

4. **Run the Platform:**
   Start all components using Docker Compose:

   ```bash
   docker-compose up
   ```

   This command builds and runs the Docker images for your platform, including Kafka, MySQL, and other dependencies.

**Troubleshooting:**

If you encounter any issues during setup, you can use the following commands to check the status of your Kafka queues:

1. **Open a Kafka container shell:**
   ```bash
   docker exec -it kafka bash
   ```

2. **Use `kafkacat` to check queue contents:**
   ```bash
   kafkacat -b kafka:29092 -t topic-you-want-to-check -C
   ```
   Replace `topic-you-want-to-check` with the actual name of the topic you want to inspect.

## Sample and Usage

**Sample Input Body:**

```json
{
  "id": "69896",
  "userName": "sudeep162002",
  "Language": "java",
  "output": "..."  // Paste your code here
}
```

**Code Deparser and Parser:**

This platform utilizes code deparser and parser functionalities to prepare code for execution and display the results.

 **Code Deparser and Parser (TypeScript):**

```typescript
// Code Deparser: Transforms code for execution
const deparser = (inputString: string): string => {
  return inputString
    .replace(/\\n/g, '\n')   // Restore actual line breaks
    .replace(/^   /gm, '')   // Remove leading indentation
    .replace(/\\"/g, '"');   // Replace escaped double quotes
};

// Code Parser: Formats code for display
const parser = (normalCode: string): string => {
  return normalCode
    .replace(/"/g, '\\"')    // Escape double quotes
    .split('\n')           // Split into lines
    .map(line => `   ${line}`) // Indent each line
    .join('\\n\\n');       // Join lines with "\n\n"
};
```

**Usage:**

**Deparsing Code for Execution:**

```typescript
const preparedCode = deparser(inputCode);
// Send preparedCode to worker for execution
```

**Parsing Results for Display:**

```typescript
const formattedOutput = parser(executionResult);
// Display formattedOutput to the user
```


**Execution:**

1. Send a POST request to the API endpoint (`http://localhost:<API_PORT>/execute`) with the sample input body containing your code.
2. The platform uses Server-Sent Events (SSE) to push execution results to the frontend in real-time.

**Remember:**

- Replace placeholders like `<API_PORT>` with appropriate values.
- Ensure your code is properly formatted and adheres to the supported language's syntax.


## Contributing to the Project

We believe in fostering a collaborative environment! Your contributions are valued! Feel free to explore the project, share your ideas, and submit pull requests to shape the future of Executioner.

**Roadmap (To-Do):**

- **Frontend Development:** Create a user-friendly interface for interacting with the platform. This includes features like code submission, execution monitoring, and result display.
- **Prometheus Monitoring:** Integrate with Prometheus for detailed infrastructure and application monitoring, providing valuable insights into platform performance and resource usage.
- **CI/CD Pipeline:** Establish a continuous integration and continuous delivery (CI/CD) pipeline utilizing GitHub Actions. This will automate tasks like building and pushing Docker images for supported languages to a Docker hub repository, streamlining the deployment process and ensuring consistent updates.
- **Language Support Expansion:** Gradually expand support for additional programming languages, catering to a wider developer audience.

## APIs

**Execution API:**

```
Method: POST
URL: http://localhost:5000/run-code
Body:
```

```json
{
  "id": "69896",
  "userName": "sudeep162002",
  "Language": "java",
  "output": "..."  // Paste your code here
}
```

**Execution Status and Output API:**

```
Method: GET
URL: http://localhost:5000/{reqId}/{username}
```

Replace `{reqId}` with the ID of the execution request and `{username}` with the username for whom the request was made.

**Response:**

```json
{
  "status": "pending" | "running" | "completed" | "failed",
  "output": "..." (only present if status is "completed")
}
```

**Note:**

- Replace placeholders like `<API_PORT>` with appropriate values.
- Ensure your code is properly formatted and adheres to the supported language's syntax.
