package cryptography

import (
	"crypto/sha256"
	connection "sql-test/Connection"
)

func HashMessage(v string, c string, desc string) {
	HashValue := sha256.Sum256([]byte(v))
	HashCred := sha256.Sum256([]byte(c))
	//fmt.Printf("[+] Hashed Credentials\n - Key: %x\n - Value: %x\n", HashValue, HashCred)
	connection.DBInsert(HashValue[:], HashCred[:], desc)
}
