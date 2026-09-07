package main

import "fmt"
import "time"

func main() {
    	fmt.Println("Hello, Fedora!")
	targetDate := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.Local)
	left := int(time.Until(targetDate).Hours() / 24)
	fmt.Println(left)
}
