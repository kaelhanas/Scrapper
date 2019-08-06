package config

type ViperConfiguration struct {
	Mongo MongoConfiguration
	Jsonpath JsonpathConfiguration
}

type MongoConfiguration struct {
	Server ServerConfiguration
}

type ServerConfiguration struct {
	Host string
	Port int
}

type AccountConfiguration struct {
	Username string
	Password string
	HostDB string
}

type DbsConfiguration struct {
	Scrapper ScrapperConfiguration
	Common CommonConfiguration
}

type ScrapperConfiguration struct {

}

type ScrapperCollections struct {

}

type CommonConfiguration struct {

}


type JsonpathConfiguration struct {

}