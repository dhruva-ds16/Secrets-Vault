package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnCheck() {
	db, err := sql.Open("mysql", "root:zrUi53bJzOy7k6jOAMQEgvg+wJEfhHFZ@tcp(127.0.0.1:8989)/VAULT")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println("[+] Database Connection Established...")
	}
}
