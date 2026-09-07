package main

import "fmt"
import "time"

func DaysLeft(today, target time.Time) int {
	hours := target.Sub(today).Hours()
	return int(hours / 24)
}


func main() {
    	fmt.Println("Hello, Fedora!")
	targetDate := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.Local)

	left := DaysLeft(time.Now(), targetDate)
	fmt.Println(left)
}
