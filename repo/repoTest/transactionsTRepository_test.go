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


func TestInsertTrans(t *testing.T) {

	json_T := model.Json_T{
		Id_user: 9999999999,
		Connections: []model.Connection_T{
			{
				Id_connection: 9999999999,
				Last_update:   "now",
				Accounts: []model.Account{
					{
						Id_account: 9999999999,
						Last_update: "now",
						Currency: model.Currency{
							Name:   "Euro",
							Symbol: "€",
						},
						Original_name: "original name",
						Balance: 99999.99999,
						Account_type: "check",
						Transactions: []model.Transaction{
							{
								Wording: "wording",
								Original_wording: "Original wording",
								Transaction_type: "virement",
								Date: "date",
								Rdate: "Rdate",
							},
						},
					},
				},
			},
		},
	}

	repoInstance := repo.GetTInstance()

	repoInstance.InsertTrans(json_T)

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
