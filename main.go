package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Your MongoDB Atlas Connection String
const (
	mongoURI       = "mongodb+srv://mananhsmpatel:CWL6x9CTDtRRXmQy@level79db.ycxykby.mongodb.net/?retryWrites=true&w=majority"
	databaseName   = "Level79-Clothing"
	collectionName = "Tshirt-inventory"
	batchOperation = "inventory/tshirt"
)

// A global variable that will hold a reference to the MongoDB client
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB. If it cannot connect, it will fail, and the program will exit.
func init() {
	if err := connectToMongoDB(); err != nil {
		fmt.Println("Could not connect to MongoDB")
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, this is the level79 Database for inventory management",
		})
	})

	r.GET("/batches", getAllBatches)
	r.GET("/batches/:batchid", getBatchByID)
	//r.GET("/batches/:batchid", getBatchByID) // New endpoint to get batch by batchid

	// Run the server
	if err := r.Run(":818"); err != nil {
		fmt.Println("Error starting server:", err)
	}
	// Run the server
	r.Run(":818")
}

// Our implementation logic for connecting to MongoDB
func connectToMongoDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	mongoClient = client
	return nil
}

// Handler function to get all batches
func getAllBatches(c *gin.Context) {
	// Simulate fetching data from MongoDB, replace with your actual logic
	collection := mongoClient.Database("Level79-Clothing").Collection("Tshirt-inventory")
	result := collection.FindOne(context.TODO(), bson.M{"operation": "inventory/tshirt"})

	var rawData map[string]interface{}
	if err := result.Decode(&rawData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	batchData, exists := rawData["batch"]
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Batch data not found"})
		return
	}

	c.JSON(http.StatusOK, batchData)
}

func getBatchByID(c *gin.Context) {
	batchID, err := strconv.Atoi(c.Param("batchid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid batchid"})
		return
	}

	collection := mongoClient.Database(databaseName).Collection(collectionName)
	result := collection.FindOne(context.TODO(), bson.M{
		"operation": batchOperation,
		"batch": bson.M{
			"$elemMatch": bson.M{"batchid": batchID},
		},
	})

	var batch bson.M
	if err := result.Decode(&batch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %v", err)})
		return
	}

	c.JSON(http.StatusOK, batch)
}
