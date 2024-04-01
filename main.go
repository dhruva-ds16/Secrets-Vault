package main

import (
	"bufio"
	"fmt"
	"log"
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
	var k, v, text string
	fmt.Printf("[-] Enter Key and Value: \n :")
	fmt.Scanln(k, v)
	scan := bufio.NewReader(os.Stdin)
	fmt.Printf("[-] Enter Description: \n :")
	text, err := scan.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Hashing...")
	cryptography.HashMessage(k, v, text)
	switchStatement()
}

func switchStatement() {
	fmt.Println("[+] Select an option")
	fmt.Printf(" - 1. Insert\n - 2. Search\n - 3. Delete\n")
	var opt int32
	fmt.Scan(&opt)
	switch opt {
	case 1:
		InsertCase()
		time.Sleep(3 * time.Second)
	case 2:
		searchEntry()
		time.Sleep(3 * time.Second)
	case 3:
		DelEntry()
		time.Sleep(3 * time.Second)
	default:
		fmt.Println("[-] Invalid Option")
		time.Sleep(3 * time.Second)
	}
}

func searchEntry() {
	var opt int
	fmt.Println("[+] Database Query")
	fmt.Printf(" - 1. Search by Key\n - 2. Search by Value\n - 3. Main Menu\n")
	fmt.Scan(&opt)
	switch opt {
	case 1:
		connection.KeySearch()
	case 2:
		connection.ValueSearch()
	case 3:
		fmt.Println("[-] Going back to the main menu...")
		switchStatement()
	default:
		fmt.Println("[-] Invalid Option")
	}
	switchStatement()
}

func DelEntry() {
	var opt int
	fmt.Println("[+] Database Query")
	fmt.Printf(" - 1. Delete by Key\n - 2. Delete by Value\n - 3. Main Menu\n")
	fmt.Scan(&opt)
	switch opt {
	case 1:
		connection.DelKey()
	case 2:
		connection.DelValue()
	case 3:
		fmt.Println("[-] Going back to the main menu...")
		switchStatement()
	default:
		fmt.Println("[-] Invalid Option")
	}
	switchStatement()
}
