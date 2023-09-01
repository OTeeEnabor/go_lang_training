package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)
func main () {
	_, err := userInput()

	if err != nil {
		fmt.Println("Invalid number try again")
		_, err := userInput()
		if err != nil {
			fmt.Println("game over")
		}
	}
}
func userInput () (int8, error){
	// fmt.Println("Hello, my name is Oselu")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	strAge,err := reader.ReadString('\n')
	if err != nil {
		return -1, err

		// fmt.Println("Please enter a number")
	}
	cleanAge := strings.Replace(strAge, "\n", "", -1)
	// remove spaces
	cleanAge = strings.TrimSpace(cleanAge)
	// conv string to int
	numberAge,err := strconv.ParseInt(cleanAge,10,8)

	return numberAge, err
}
func rowLogic(numberAge int8) {
	if numberAge ==18 {
		fmt.Println("Welcome to BC Club")
	} else if (numberAge > 18) && ( numberAge < 30) {
		fmt.Println("row2")
	}else if (numberAge >= 30) && ( numberAge < 60) {
		fmt.Println("row3")
	}else if numberAge >= 60 {
		fmt.Println("vip")
	}else {
		fmt.Println("Sorry bro, you a youngin! get outta here")
	}
}