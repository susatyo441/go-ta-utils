package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var dbName string

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

// ConnectMongo initializes MongoDB client once
func ConnectMongo() *mongo.Client {
	mongoUri := os.Getenv("MONGO_URI")

	if mongoUri == "" {
		mongoUri = "mongodb://localhost:27017"
	}

	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(mongoUri)
		var err error
		clientInstance, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		err = clientInstance.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}

		fmt.Println("MongoDB client initialized!")
	})

	return clientInstance
}

func ConnectToCustomDb(db string) {
	Client = ConnectMongo()
	dbName = db
}

func ConnectToCompanyDb(companyCode string) {
	Client = ConnectMongo()
	dbName = fmt.Sprintf("%s_tagsamurai", companyCode)
}

func ConnectToPartnerDb(partnerId string) {
	Client = ConnectMongo()
	dbName = fmt.Sprintf("%s_admin_tagsamurai", partnerId)
}

func ConnectToGlobalDb() {
	Client = ConnectMongo()
	dbName = "tagsamurai"
}

func ConnectToAdminDb() {
	Client = ConnectMongo()
	dbName = "admin_tagsamurai"
}

func ConnectToShopVisionDb() {
	Client = ConnectMongo()
	dbName = "shop_vision"
}

func GetCollection(collection string) *mongo.Collection {
	return Client.Database(dbName).Collection(collection)
}
