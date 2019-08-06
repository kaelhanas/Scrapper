package repo

import "sync"

var once sync.Once

//TO-DELETE bdd de test
const s_collectionName = "synchro"
const t_collectionName = "transaction"

const sync_collectionName = "synchro"
const trans_collectionName = "transactions"

const scrapperDBName = "scrapperDB"