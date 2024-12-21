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
	case "6":
		sixthDay()
	case "10":
		tenthDay()
	case "11":
		eleventhDay()
	case "11.2":
		eleventhDayP2()
	case "12":
		twelfthDay()
	case "13":
		thirteenthDay()
	case "14":
		fourteenthDay()
	case "15":
		fifteenthDay()
	case "16":
		sixteenthDay()
	case "17":
		seventeenthDay()
	case "18":
		eighteenthDay()
	case "19":
		ninteenthDay()
	case "20":
		twentiethDay()
	case "21":
		twentyFirstDay()
	default:
		fmt.Println("Invalid day")
	}
}
