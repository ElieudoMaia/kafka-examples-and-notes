# Simple Go Project: Kafka Consumer and Producer

This project serves as a simple example of how to set up, connect, publish, and consume messages using Kafka in a Go application. The entire environment is fully containerized using Docker.

## Features

- **Kafka Producer:** Publishes messages to a Kafka topic.
- **Kafka Consumer:** Consumes messages from a Kafka topic.
- **Containerized Environment:** Everything is set up and run within Docker containers for ease of use and consistency.

## Prerequisites

Ensure you have the following installed on your system:

- Docker
- Docker Compose

## Getting Started

### Running the Application

To start the Kafka environment and the Go application, simply run the following command in your terminal:

```bash
docker-compose up -d
```

This command will:

- Build and start Docker containers for Kafka, Zookeeper, and the Go application.
- Set up the necessary environment to produce and consume messages.

### Creating a Kafka Topic

Before publishing or consuming messages, you need to create a Kafka topic. Run the following command to create a topic named `test` with 3 partitions:

```bash
docker exec -it <kafka-container-name> kafka-topics --create --bootstrap-server=localhost:9092 --topic=test --partitions=3
```

Replace `<kafka-container-name>` with the actual name of your Kafka container. You can find the container name by running `docker ps`.

### Consuming Messages from the CLI

To consume messages from the `test` topic using the Kafka console consumer, run:

```bash
docker exec -it <kafka-container-name> kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test
```

This command will start consuming messages from the specified topic and display them in your terminal.

## Additional Notes

- **Kafka Configuration:** The Kafka broker is configured to run on `localhost:9092`, and Zookeeper runs on `localhost:2181`. You can customize these settings in the `docker-compose.yml` file.
- **Environment Variables:** You can adjust the environment variables for Kafka and the Go application as needed in the `docker-compose.yml` file.

## Project Structure

- `main.go`: The main application file that contains the Go code for the Kafka consumer and producer.
- `Dockerfile`: Builds the Go application into a Docker image.
- `docker-compose.yml`: Defines the Docker containers for Kafka, Zookeeper, and the Go application.

## Troubleshooting

- **Docker Logs:** If you encounter any issues, you can check the logs of the running containers using `docker logs <container-name>`.
- **Kafka CLI Tools:** The Kafka CLI tools are available inside the Kafka container, which you can access using `docker exec`.
