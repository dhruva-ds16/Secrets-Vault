package connection

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func KeySearch() {
	db, err := sql.Open("mysql", "root:zrUi53bJzOy7k6jOAMQEgvg+wJEfhHFZ@tcp(127.0.0.1:8989)/VAULT")
	if err != nil {
		panic(err.Error())
	}
	hash_k := QueryHash()
	var ans string
	errQuery := db.QueryRow("SELECT credential FROM Alice WHERE value = ?", hash_k).Scan(&ans)
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	fmt.Printf("Found: %s\n", ans)
	defer db.Close()
}

func QueryHash() (out string) {
	var k string
	fmt.Println("[+] Enter Query Data")
	fmt.Scan(&k)
	hash_k := sha256.Sum256([]byte(k))
	str := hex.EncodeToString(hash_k[:])
	return str
}

func ValueSearch() {
	db, err := sql.Open("mysql", "root:zrUi53bJzOy7k6jOAMQEgvg+wJEfhHFZ@tcp(127.0.0.1:8989)/VAULT")
	if err != nil {
		panic(err.Error())
	}
	hash_k := QueryHash()
	var ans string
	errQuery := db.QueryRow("SELECT value FROM Alice WHERE credential = ?", hash_k).Scan(&ans)
	if errQuery != nil {
		panic(errQuery)
	}
	fmt.Printf("Found: %s\n", ans)
	defer db.Close()
}
