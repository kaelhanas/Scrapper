package service

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
)

const jsonT = "D:\\CRESUS\\bgv\\Workspace\\Scrapper_POC\\jsons\\transactions.json"
const jsonS = "D:\\CRESUS\\bgv\\Workspace\\Scrapper_POC\\jsons\\synchros.json"

func DetectJsonType(jsonPath string) (string, string){

	jsonFile, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		log.Fatal(err)
	}

	value := gjson.Get(string(jsonFile), "connections.#.accounts")

	if len(value.Array()) != 0 {
		return string(jsonFile), "transaction"
	} else {
		return string(jsonFile), "transaction"
	}

}


