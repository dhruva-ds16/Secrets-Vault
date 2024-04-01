package main

import (
	"fmt"
	connection "sql-test/Connection"
	cryptography "sql-test/Cryptography"
)

func main() {
	connection.DBConnCheck()
	for {
		fmt.Println("[+] Select an option")
		fmt.Println(" - 1. Insert")
		fmt.Println(" - 2. View")
		fmt.Println(" - 3. Delete")
		fmt.Printf(":")
		var opt int32
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			InsertCase()
		case 2:
			fmt.Println("[-] Work in progress...")
		case 3:
			fmt.Println("[-] Work in progress...")
		default:
			fmt.Println("[-] Invalid Option")
		}

	}
}

func InsertCase() {
	var k, v string
	fmt.Println("[-] Enter Key and Value: ")
	fmt.Scan(&k, &v)
	fmt.Println("[+] Hashing...")
	cryptography.HashMessage(k, v)
}
