package main

import (
	// Add required Go packages
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// Add the MongoDB driver packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Batch struct {
	ID           string                    `json:"id,omitempty" bson:"_id,omitempty"`
	Operation    string                    `json:"operation" bson:"operation"`
	Purchase     int                       `json:"purchase" bson:"purchase"`
	Availability string                    `json:"availability" bson:"availability"`
	Sizes        []string                  `json:"sizes" bson:"sizes"`
	Colors       map[string]map[string]int `json:"colors" bson:"colors"`
	Variants     []Variant                 `json:"variants" bson:"variants"`
	Supplier     Supplier                  `json:"supplier" bson:"supplier"`
	CurrentStock int                       `json:"current_stock" bson:"current_stock"`
}

// Variant represents the data structure for a T-shirt variant
type Variant struct {
	ID      int         `json:"id,omitempty" bson:"id,omitempty"`
	Name    string      `json:"name" bson:"name"`
	Details VariantInfo `json:"details" bson:"details"`
	Stock   int         `json:"stock" bson:"stock"`
}

// VariantInfo represents additional details for a T-shirt variant
type VariantInfo struct {
	Material   string `json:"material" bson:"material"`
	DesignType string `json:"design_type" bson:"design_type"`
	Rating     int    `json:"rating" bson:"rating"`
}

// Supplier represents the data structure for a T-shirt supplier
type Supplier struct {
	Name     string `json:"name" bson:"name"`
	Location string `json:"location" bson:"location"`
	Rating   int    `json:"rating" bson:"rating"`
}

// Your MongoDB Atlas Connection String
const uri = "mongodb+srv://mpatel:mnanhsm333dj@level79db.ycxykby.mongodb.net/?retryWrites=true&w=majority"

// A global variable that will hold a reference to the MongoDB client
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB. If it cannot connect it will fail and the program will exit.
func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Initialize GIN
	r := gin.Default()

	// Routes
	r.GET("/batches", getBatches)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Our implementation logic for connecting to MongoDB
func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}
func getClient() *mongo.Client {
	return mongoClient
}

func getBatches(c *gin.Context) {
	var batches []Batch

	cursor, err := getClient().Database("Level79-Clothing").Collection("Tshirt-inventory").Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var batch Batch
		if err := cursor.Decode(&batch); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		batches = append(batches, batch)
	}

	c.JSON(http.StatusOK, batches)
}
