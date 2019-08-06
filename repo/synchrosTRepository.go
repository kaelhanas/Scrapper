package repo

import (
	"Scrapper_POC/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"

	"Scrapper_POC/mongoConnections"
)

type synchrosRepo struct {
	collection *mongo.Collection
}

var s_instance *synchrosRepo
var s_once sync.Once

func GetSInstance() *synchrosRepo {

	s_once.Do(func() {
		var mongoClient= mongoConnections.GetInstance().GetMongoClient()
		s_instance = &synchrosRepo{collection: mongoClient.Database("test").Collection(s_collectionName)}
	})

	return s_instance

}

func (sr *synchrosRepo) GetCollection() *mongo.Collection {
	return sr.collection
}

func (sr *synchrosRepo) InsertSynchro(newSynchro model.Json_S) {
	insertResult, err := sr.GetCollection().InsertOne(context.TODO(), newSynchro)

	if err != nil {
		log.Println("synchrosRepository -> insertion failed")
		log.Fatal(err)
	}

	log.Printf("New document inserted in collection %v : %v\n",s_collectionName, insertResult.InsertedID)
}

func (sr *synchrosRepo) collectionExist() bool {
	result := false

	// get mongo client
	mClient := mongoConnections.GetInstance().GetMongoClient()

	// get the "synchro" collection if it exists
	collectionOptions := options.ListCollections()
	cursor, err := mClient.Database("test").ListCollections(context.TODO(), bson.D{{"name", "synchro"}}, collectionOptions)
	if err != nil {
		log.Fatal(err)
	}

	// if the collection exist, cursor contains ONE element
	for cursor.Next(context.TODO()) {
		result = true
	}

	return result
}