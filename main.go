package main

import (
	"fmt"
	connection "sql-test/Connection"
)

func main() {
	connection.DBConnCheck()
	fmt.Println("Finally got the connection running!!")
}
