package main

import (
	"github.com/meshachdamilare/banking-api/database"
	"log"
)

func main() {
	//err := database.MigratePostgres()
	//if err != nil {
	//	panic(err)
	//}
	err := database.MigrateImmuDB()
	if err != nil {
		log.Println(err)
	}
}
