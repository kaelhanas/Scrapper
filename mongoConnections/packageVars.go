package mongoConnections

import "sync"

var once sync.Once

const mongoURI = "mongodb://localhost:27017"

const scrapperUsername = "scrapper"
const scrapperPswd = "scrapperpwd"

const scrapperDBName = "scrapperDB"
const commonDBName = "common"