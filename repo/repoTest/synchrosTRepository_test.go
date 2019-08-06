package repoTest

import (
	"Scrapper_POC/model"
	"Scrapper_POC/repo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)


func TestInsertSynchro(t *testing.T) {

	json_S := model.Json_S{
		Id_user: 9999999999,
		Connections: []model.Connection_S{
			{
				Id_connection: 9999999999,
				Last_update:   "now",
				Id_accounts:   []int{100, 101, 102},
			},
		},
	}

	repoInstance := repo.GetSInstance()

	repoInstance.InsertSynchro(json_S)

	findOptions := options.Find()
	filter := bson.D{{"id_user", 9999999999}}
	var results []model.Json_S
	collection := repoInstance.GetCollection()

	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var elem model.Json_S

		if err := cursor.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
		collection.DeleteOne(context.TODO(), filter)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())

	if len(results) == 0 {
		t.Errorf("Document has not been found so problably not inserted\n")
	} else if len(results) >1 {
		t.Errorf("To many document found :\n")
		for _, elem := range results {
			t.Errorf("-> %+v\n", elem)
		}
	}


}
