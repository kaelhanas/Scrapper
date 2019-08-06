package mongoConnections

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoDBConnection struct {
	mongoClient *mongo.Client
}

var instance *mongoDBConnection

func GetInstance() *mongoDBConnection {

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(mongoURI)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connected to MongoDB!")
		}

		instance = &mongoDBConnection{mongoClient: client}
	})

	return instance

}

func (mc *mongoDBConnection) GetMongoClient() *mongo.Client {
	return mc.mongoClient
}

// Juste la pour faire jolie, on l'utilisera jamais :D
func (mc *mongoDBConnection) SetMongoClient(newClient *mongo.Client) {
	mc.mongoClient = newClient
}
