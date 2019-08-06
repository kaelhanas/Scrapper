package repo

import (
	"Scrapper_POC/model"
	"Scrapper_POC/mongoConnections"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type transactionsRepo struct {
	collection *mongo.Collection
}

var trans_instance *transactionsRepo


func GetTransInstance() *transactionsRepo {

	once.Do(func() {
		scrapperDB := mongoConnections.GetScrapperDB()
		trans_instance = &transactionsRepo{collection: scrapperDB.Collection(trans_collectionName)}
	})

	return trans_instance

}

func (tr *transactionsRepo) getCollection() *mongo.Collection {
	return tr.collection
}

func (tr *transactionsRepo) InsertTrans(newTrans model.Json_T) {
	insertResult, err := tr.getCollection().InsertOne(context.TODO(), newTrans)

	if err != nil {
		log.Println("transactionsRepository -> insertion failed")
		log.Fatal(err)
	}

	log.Printf("New document inserted in collection %v : %v\n",trans_collectionName, insertResult.InsertedID)
}

// Deprecated
/*func (tr *transactionsRepo) collectionExist() bool {
	result := false

	// get mongo client
	mClient := mongoConnections.getScrapperClientInstance().GetScrapperClient()

	// get the "transactions" collection if it exists
	collectionOptions := options.ListCollections()

	filter := bson.D{{"name", trans_collectionName}}
	cursor, err := mClient.Database(scrapperDBName).ListCollections(context.TODO(), filter, collectionOptions)
	if err != nil {
		log.Fatal(err)
	}

	// if the collection exist, cursor contains ONE element
	for cursor.Next(context.TODO()) {
		result = true
	}

	return result
}*/