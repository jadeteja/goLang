package main

import "fmt"

const (
	sunday = iota + 1
	monday 
	tuesday
	wednesday
	thursday
	friday
	saturday
)

type logLevel int

const(
	logError logLevel = iota  
)

func main() {
	fmt.Println(monday, 
	tuesday,
	wednesday,
	thursday,
	friday,
	saturday);
}
