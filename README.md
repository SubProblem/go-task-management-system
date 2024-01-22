# Task Management System

Task Management System is a microservices-based application designed for efficient task management, user authentication/authorization, and timely task notifications. The project is built with **Go**, **Gorilla Mux**, **Kafka**, **Nginx**, and utilizes **Docker** for **PostgreSQL** and **Apache Kafka**.

## Microservices Overview

### 1. Security Service

The **Security Service** is responsible for user authentication and authorization, ensuring secure access to the system, and managing user credentials.

**Integration with Kafka**

This service communicates with **Apache Kafka** to handle authentication events and maintain a secure user environment.

### 2. Task Management Service

The **Task Management Service** facilitates task-related operations such as task creation, retrieval, updating, and deletion. Additionally, a scheduler runs once every day to check task deadlines.

**Integration with Kafka**

Task Management Service interacts with **Apache Kafka** to exchange messages related to task events and updates.

### 3. Notification Service

The **Notification Service** is in charge of sending timely notifications to users before the deadline of their tasks. Currently, it sends email notifications for impending deadlines.

## Reverse Proxy with Nginx

The project includes a reverse proxy implemented using **Nginx**. This Nginx server serves multiple purposes:

- **Load Balancer:**
  - Distributes incoming requests across multiple instances of microservices for load balancing.

- **Gateway:**
  - Acts as a gateway to manage and control the flow of requests to the respective microservices.

- **Route Security:**
  - Secures specific routes, ensuring that sensitive    functionalities are appropriately protected.


**Integration with Kafka**

Similar to other services, the Notification Service utilizes **Apache Kafka** for communication and coordination.

## Docker Setup

The project includes **Docker** configurations for **PostgreSQL** and **Apache Kafka**. Use the provided Docker Compose file to orchestrate these services.

### Prerequisites

- **Docker**
- **Docker Compose**

### Usage

1. **Start PostgreSQL and Apache Kafka containers:**

    ```bash
    docker-compose up -d
    ```



