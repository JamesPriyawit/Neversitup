package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var lockInitPostgresDBPool sync.Mutex
var PostgresPool *sql.DB
var countFailPostgres int

func initPostgresPool() *sql.DB {

	lockInitPostgresDBPool.Lock()
	defer lockInitPostgresDBPool.Unlock()

	var err error

	if PostgresPool != nil {
		err = PostgresPool.Ping()

		if err != nil {
			PostgresPool.Close()
		} else {
			countFailPostgres = 0
			return PostgresPool
		}
	}
	PostgresPool, err = sql.Open("postgres", "port=5432 user=postgres password=password host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Println("Can not initial Postgres DB Pool", err.Error())
	} 
	return PostgresPool
}

func GetConnectPostgresDB() *sql.DB {

	var err error

	if PostgresPool == nil {
		PostgresPool = initPostgresPool()
	}
	err = PostgresPool.Ping()
	if err != nil {
		log.Print("Can not verify a connection to Postgres database because "+err.Error())
		countFailPostgres++
		if countFailPostgres > 20 {
			PostgresPool = initPostgresPool()
		}
	} else {
		countFailPostgres = 0
	}
	return PostgresPool
}
