package main

import (
	"bufio"
	"fmt"
	"os"
	connection "sql-test/Connection"
	cryptography "sql-test/Cryptography"
	"time"
)

func main() {
	connection.DBConnCheck()
	for {
		switchStatement()
		fmt.Println("Going back to the main menu...")
		time.Sleep(3 * time.Second)
		switchStatement()
	}
}

func InsertCase() {
	var k, v string
	fmt.Printf("[-] Enter Key and Value: \n :")
	fmt.Scanln(&k, &v)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[-] Enter Description: \n :")
	text, _ := reader.ReadString('\n')
	fmt.Println("[+] Hashing...")
	cryptography.HashMessage(k, v, text)

}

func switchStatement() {
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
		time.Sleep(3 * time.Second)
	case 2:
		fmt.Println("[-] Work in progress...")
		time.Sleep(3 * time.Second)
	case 3:
		fmt.Println("[-] Work in progress...")
		time.Sleep(3 * time.Second)
	default:
		fmt.Println("[-] Invalid Option")
		time.Sleep(3 * time.Second)
	}
}
