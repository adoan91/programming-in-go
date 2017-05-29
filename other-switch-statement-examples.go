package main

import "fmt"
import "time"

func main() {
	switch time.Now().Weekday() {
		case time.Monday, time.Tuesday, 
			time.Wednesday, time.Thursday, 
			time.Friday:
			fmt.Println("It's a weekday")
		default:
			fmt.Println("it's the weekend")
	}

	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("afternoon")
	}
}