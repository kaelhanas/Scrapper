package main

import (
	"Scrapper_POC/model"
	"Scrapper_POC/repo"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"os"
)

const jsonT = "D:\\CRESUS\\bgv\\Workspace\\Scrapper_POC\\jsons\\transactions.json"
const jsonS = "D:\\CRESUS\\bgv\\Workspace\\Scrapper_POC\\jsons\\synchros.json"

// Non fonctionelle depuis la destruction du struct User
func readJsonByUnMarshalling() {
	/*
		jsonFile, err := os.Open("D:\\CRESUS\\bgv\\Workspace\\Scrapper_POC\\jsons\\test.json")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Successfully opened test.json")

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var users Users

		json.Unmarshal(byteValue, &users)

		for i := 0; i < len(users.Users); i++ {
			fmt.Println("User Type: " + users.Users[i].Type)
			fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
			fmt.Println("User Name: " + users.Users[i].Name)
			fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		}

	*/
}

func readJsonByUnstructuredData() {
	jsonFile, err := os.Open("D:\\CRESUS\\bgv\\Workspace\\Scrapper_POC\\jsons\\test.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully opened test.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["users"])
}

func initSynchro() model.Json_S {
	json_S := model.Json_S{
		Id_user: 1,
		Connections: []model.Connection_S{
			{
				Id_connection: 11,
				Last_update:   "now",
				Id_accounts:   []int{100, 101, 102},
			},
		},
	}

	return json_S
}

func initTransaction() model.Json_T {
	json_T := model.Json_T{
		Id_user: 9999999999,
		Connections: []model.Connection_T{
			{
				Id_connection: 9999999999,
				Last_update:   "now",
				Accounts: []model.Account{
					{
						Id_account:  9999999999,
						Last_update: "now",
						Currency: model.Currency{
							Name:   "Euro",
							Symbol: "€",
						},
						Original_name: "original name",
						Balance:       99999.99999,
						Account_type:  "check",
						Transactions: []model.Transaction{
							{
								Wording:          "wording",
								Original_wording: "Original wording",
								Transaction_type: "virement",
								Date:             "date",
								Rdate:            "Rdate",
							},
						},
					},
				},
			},
		},
	}

	return json_T
}

// non fonctionnelle depuis les changements apportés à la connexion mongo
func testCollectionExistence() {
	/*
		mClient := singleton.GetInstance().GetMongoClient()
		listDbOptions := options.ListDatabases()

		dbList, err := mClient.ListDatabaseNames(context.TODO(), bson.D{}, listDbOptions)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Mongo contains 'test' database ? -> %t\n", utils.ContainString("test", dbList))

		collectionOptions := options.ListCollections()
		cursor, err := mClient.Database("test").ListCollections(context.TODO(), bson.D{{"name", "snchro"}}, collectionOptions)
		if err != nil {
			log.Fatal(err)
		}
		result := false
		for cursor.Next(context.TODO()) {
			result = true
		}

		fmt.Printf("%t",result)*/
}

func A0_testMongoConnectionWithContext() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx = context.WithValue(ctx, "host", "localhost:27017")
	ctx = context.WithValue(ctx, "username", "scrapper")
	ctx = context.WithValue(ctx, "password", "scrapperpwd")
	ctx = context.WithValue(ctx, "database", "scrapperDB")

	db, err := A0_configDB(ctx)
	if err != nil {
		log.Fatalf("database configuration failed : %v", err)
	}

	_, err = db.Collection("synchro").InsertOne(context.TODO(), initSynchro())
	if err != nil {
		log.Fatal(err)
	}

	client, err := A0_configMongoConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	syncInsert, err := client.Database("scrapperDB").Collection("synchro").InsertOne(context.TODO(),
		initSynchro())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Document inserter in ScrapperDB::synchro -> \n%+v", syncInsert)

	filter := bson.D{{"id_user", 9999999999}}
	result := model.Json_S{}
	err = client.Database("common").Collection("users").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Document inserter in ScrapperDB::synchro -> \n%+v", result)
}

func A0_configDB(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value("username"),
		ctx.Value("password"),
		ctx.Value("host"),
		ctx.Value("database"), )

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to mongo : %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client coulnd't connect with background context : %v", err)
	}

	db := client.Database("scrapperDB")
	return db, nil
}

func A0_configMongoConnection(ctx context.Context) (*mongo.Client, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		ctx.Value("username"),
		ctx.Value("password"),
		ctx.Value("host"),
		ctx.Value("database"), )

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to mongo : %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client coulnd't connect with background context : %v", err)
	}

	return client, nil
}

func main() {

	//service.DetectJsonType(jsonS)

	//repo.GetSInstance().InsertSynchro(test)

	repo.GetSyncInstance().InsertSynchro(initSynchro())

}