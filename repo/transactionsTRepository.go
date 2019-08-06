package repo

import (
	"Scrapper_POC/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"

	"Scrapper_POC/mongoConnections"
)

type transRepo struct {
	collection *mongo.Collection
}

var t_instance *transRepo
var t_once sync.Once


func GetTInstance() *transRepo {

	t_once.Do(func() {
		var mongoClient= mongoConnections.GetInstance().GetMongoClient()
		t_instance = &transRepo{collection: mongoClient.Database("test").Collection(t_collectionName)}
	})

	return t_instance

}

func (tr *transRepo) GetCollection() *mongo.Collection {
	return tr.collection
}

func (tr *transRepo) InsertTrans(newTrans model.Json_T) {
	insertResult, err := tr.GetCollection().InsertOne(context.TODO(), newTrans)

	if err != nil {
		log.Println("transactionsRepository -> insertion failed")
		log.Fatal(err)
	}


	log.Printf("New document inserted in collection %v : %v\n",t_collectionName, insertResult.InsertedID)
}