package main

import (
	"fmt"
	"time"
)

var name string = "Gopher"

func main() {
	fmt.Println("Hello, " + name)
	printTime()
}

func printTime() {
	t := time.Now()
	// day-month-year hour:minute:second
	fmt.Println(t.Format("02-01-2006 15:04:05 MST"))
}
