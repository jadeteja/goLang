package main

import "fmt"

func main(){
	firstNumber := 100
	secondNumber := 200
	thirdNumber := 300

	firstNumber, secondNumber = secondNumber, firstNumber;
	thirdNumber = firstNumber + secondNumber;
	fmt.Println(firstNumber, secondNumber, thirdNumber);
}