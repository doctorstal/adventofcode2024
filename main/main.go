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
	case "3":
		thirdDay()
	case "4":
		fourthDay()
	case "10":
		tenthDay()
	default:
		fmt.Println("Invalid day")
	}
}
