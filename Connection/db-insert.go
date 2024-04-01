package connection

import (
	"database/sql"
	"encoding/hex"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBInsert(key []byte, value []byte, desc string) {
	db, err := sql.Open("mysql", "root:zrUi53bJzOy7k6jOAMQEgvg+wJEfhHFZ@tcp(127.0.0.1:8989)/VAULT")
	if err != nil {
		panic(err.Error())
	}
	str_key := hex.EncodeToString(key)
	//fmt.Println(str_key)
	str_val := hex.EncodeToString(value)
	//fmt.Println(str_val)
	//var test_desc string = "Test description"
	ins, err := db.Query("INSERT INTO Alice (value, credential, description) VALUES (?, ?, ?)", str_key, str_val, desc)
	if err != nil {
		panic(err.Error())
	}
	defer ins.Close()
	fmt.Println("[+] Credentials inserted...")
}
