## **Executioner: A Scalable Code Execution Platform for Everyone**

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

**Contributing to the Project:**

We believe in fostering a collaborative environment. Your contributions are valued! Feel free to explore the project, share your ideas, and submit pull requests to shape the future of Executioner.

**Join us in creating the ultimate code execution platform!**