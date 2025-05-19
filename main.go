package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"work_with_db/internal/dbs/postgres"
)

func main() {
	conf := postgres.NewConfig()
	sql, err := postgres.NewDb(conf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sql)

}
