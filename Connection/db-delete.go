package connection

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DelKey() {
	db, err := sql.Open("mysql", "root:zrUi53bJzOy7k6jOAMQEgvg+wJEfhHFZ@tcp(127.0.0.1:8989)/VAULT")
	if err != nil {
		panic(err.Error())
	}
	hash_k := DelQueryHash()
	_, errQuery := db.Query("DELETE FROM Alice WHERE value = ?", hash_k)
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	fmt.Printf("[+] Entry Deleted")
	defer db.Close()
}

func DelQueryHash() (out string) {
	var k string
	fmt.Println("[+] Enter Query Data")
	fmt.Scan(&k)
	hash_k := sha256.Sum256([]byte(k))
	str := hex.EncodeToString(hash_k[:])
	return str
}

func DelValue() {
	db, err := sql.Open("mysql", "root:zrUi53bJzOy7k6jOAMQEgvg+wJEfhHFZ@tcp(127.0.0.1:8989)/VAULT")
	if err != nil {
		panic(err.Error())
	}
	hash_k := DelQueryHash()
	_, errQuery := db.Query("DELETE FROM Alice WHERE credential = ?", hash_k)
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	fmt.Printf("[+] Entry Deleted\n")
	defer db.Close()
}
