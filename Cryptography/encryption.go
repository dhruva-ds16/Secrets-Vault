package cryptography

import (
	"crypto/sha256"
	"fmt"
	connection "sql-test/Connection"
)

func HashMessage(v string, c string) {
	HashValue := sha256.Sum256([]byte(v))
	HashCred := sha256.Sum256([]byte(c))
	fmt.Printf("[+] Hashed Credentials\n - Key: %x\n - Value: %x\n", HashValue, HashCred)
	var desc string
	fmt.Println("[-] Enter Description: ")
	fmt.Scanln(&desc)
	connection.DBInsert(HashValue[:], HashCred[:], desc)
}
