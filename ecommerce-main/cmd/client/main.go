package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "github.com/kishorens18/ecommerce/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	r := gin.Default()

	// Connect to gRPC service
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	// response, err := client.CreateCustomer(context.Background(), &pb.CustomerDetails{})
	// if err != nil {
	// 	log.Fatalf("Failed to call CreateTransaction: %v", err, response)
	// }

	// Define a POST route
	r.POST("/signin", func(c *gin.Context) {
		var request pb.CustomerDetails

		// Parse incoming JSON
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the gRPC service
		response, err := client.CreateCustomer(c.Request.Context(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	// Start the Gin server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start Gin server:", err)
	}
}
