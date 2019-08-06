package repo

import (
	"Scrapper_POC/model"
	"Scrapper_POC/mongoConnections"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type synchroRepo struct {
	collection *mongo.Collection
}

var sync_instance *synchroRepo


func GetSyncInstance() *synchroRepo {

	once.Do(func() {
		scrapperDB := mongoConnections.GetScrapperDB()
		sync_instance = &synchroRepo{collection: scrapperDB.Collection(sync_collectionName)}
	})

	return sync_instance

}

func (sr *synchroRepo) getCollection() *mongo.Collection {
	return sr.collection
}

func (sr *synchroRepo) InsertSynchro(newSynchro model.Json_S) {
	insertResult, err := sr.getCollection().InsertOne(context.TODO(), newSynchro)

	if err != nil {
		log.Println("synchroRepository -> insertion failed")
		log.Fatal(err)
	}

	log.Printf("New document inserted in collection %v : %v\n",sync_collectionName, insertResult.InsertedID)
}

// Deprecated
func (sr *synchroRepo) collectionExist() bool { /*
	result := false

	// get mongo client
	mClient := mongoConnections.getScrapperClientInstance().GetScrapperClient()

	// get the "synchro" collection if it exists
	collectionOptions := options.ListCollections()

	filter := bson.D{{"name", sync_collectionName}}
	cursor, err := mClient.Database(scrapperDBName).ListCollections(context.TODO(), filter, collectionOptions)
	if err != nil {
		log.Fatal(err)
	}

	// if the collection exist, cursor contains ONE element
	for cursor.Next(context.TODO()) {
		result = true
	}

	return result */
}