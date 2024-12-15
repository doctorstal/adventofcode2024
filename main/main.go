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
	case "5":
		fifthDay()
	case "10":
		tenthDay()
	case "11":
		eleventhDay()
	case "11.2":
		eleventhDayP2()
	case "12":
		twelfthDay()
	case "14":
		fourteenthDay()
	case "13":
		thirteenthDay()
	default:
		fmt.Println("Invalid day")
	}
}
