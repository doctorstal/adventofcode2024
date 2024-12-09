package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		firstDay()
	case "2":
		secondDay()
	default:
		fmt.Println("Invalid day")
	}
}
