# Golang with RabbitMQ Boilerplate
The **go-with-rabbitmq** service is a simple Go application that connects to a RabbitMQ message broker and publishes a message to a queue.

![Dashboard RabbitMQ](https://imgur.com/xCUd8bh.png)

# Getting Started
To get started with this project, you will need to have Docker and Golang installed on your machine.

## Prerequisites
- Docker
- Docker compose

## Installation
- Clone this repository
- Navigate to the project directory
- Run **docker-compose up --build** to start the containers

## For new projects
- Remove the **go.mod** and **go.sum**
- Start this project
- Run **docker exec -it <container-app> bash** in your terminal
- Into the container, you need to run **go mod init github.com\/\<user\|company\>/\<project\>**
- After this execute **go mod tidy**

## Usage
After the installation, the following containers will be running:

- **go-with-rabbitmq**
    > The Go application in **go-with-rabbitmq** container can be modified to suit your needs.
- **rabbitmq**
    > The rabbitmq service is a RabbitMQ message broker that is used by the **go-with-rabbitmq** service. 
    
    The service is **based on the official RabbitMQ Docker image** and is configured with a *default user (guest) and password (guest)*. The service exposes ports **5672** (RabbitMQ messaging port), **15672** (RabbitMQ management console port), and **15692** (RabbitMQ metrics port).

## Built with
- Golang
- RabbitMQ
- Docker

## To execute this project in manual mode
```bash
root@go-with-rabbitmq:/go/go-with-rabbitmq# go run cmd/consumer/main.go

WORKER - INFO: RabbitMQ worker has started!
CONSUMER - INFO: RabbitMQ consumer has started!
```

> **PS:** *If you want to execute this project automatic, you need modified this **Dockerfile**.*

## To execute all tests

This will run all the tests in your project, including the **TestSavingOrder** test in **OrderRepositoryTestSuite**.

> **go test -v ./...**

**PS:** *The **-v** flag will enable verbose output, which will show you the details of each test that is run.*

## Authors
- **Will3g**

## Query used in this project for orders.db with sqlite3
> CREATE TABLE orders (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key (id))
