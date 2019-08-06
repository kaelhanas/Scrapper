package mongoConnections

import (
	"Scrapper_POC/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type scrapperConnection struct {
	commonDB *mongo.Database
	scrapperDB     *mongo.Database
}

var scrapperInstance *scrapperConnection

func getScrapperConnection() *scrapperConnection {

	once.Do(func() {

		v := config.GetViper()

		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d",
			v.GetString("mongo.server.host"), v.GetInt("mongo.server.port")))
		credOptions := options.Credential{
			Username: v.GetString("mongo.account.username"),
			Password: v.GetString("mongo.account.password"),
			AuthSource: v.GetString("mongo.account.hostDB")}
		clientOptions.SetAuth(credOptions)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		scrapper_db := client.Database(v.GetString("mongo.dbs.scrapper.dbName"))
		common_db := client.Database(v.GetString("mongo.dbs.common.dbName"))

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connected to scrapperDB!")
		}

		scrapperInstance = &scrapperConnection{scrapperDB: scrapper_db, commonDB: common_db}
	})

	return scrapperInstance

}

func GetCommonDB() *mongo.Database {
	return getScrapperConnection().commonDB
}

func GetScrapperDB() *mongo.Database {
	return getScrapperConnection().scrapperDB
}
