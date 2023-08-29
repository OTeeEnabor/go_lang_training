// entry point into the
package main

// import necessary packages
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	// DAY 2
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name: ")
	name,_ := reader.ReadString('\n')
	cleanName := strings.Replace(name, "\n","",-1)
	
	fmt.Println("Enter current year: ")
	year,_ := reader.ReadString('\n')
	cleanYear := strings.Replace(year, "\n","",-1)

	fmt.Println("Enter your date of birth: ")
	dob,_ := reader.ReadString('\n')
	cleanDOB := strings.Replace(dob, "\n","",-1)

	// remove spaces
	cleanName = strings.TrimSpace(cleanName)
	cleanYear = strings.TrimSpace(cleanYear)
	cleanDOB = strings.TrimSpace(cleanDOB)
	
	// convert to numbers
	convYear,_ := strconv.ParseInt(cleanYear, 10, 64)
	convDOB,_ := strconv.ParseInt(cleanDOB, 10, 64)



	fmt.Printf("name: %v \n", cleanName)
	fmt.Printf("year :%v \n", convYear)
	fmt.Printf("dob: %v \n", convDOB)
	// const studentNumbebr int = 211155217
	// var firstName string = "Oselu"
	// var lastName string = "Enabor"
	// var age int8 = 51
	// var purchase float32 = 48.40
	// var isFinalYear bool = false
	// var favLetter byte ='a'
	// var countryCurr rune = '$'
	// var name string = fmt.Sprintf("%v %v", firstName, lastName)

	/*
	var age int8 = 64 // long form de:claration
	year := 1959 // short-form declaration, only worksk in function
	var breadPrice float32 = 19.99

	fmt.Println(age, year,breadPrice)
	*/
	/*
	DISPLAYING INFORMATION
	fmt.Println("My name is Oselu.")
	fmt.Println("My school is UJ")
	fmt.Print("My campus is APK")
	*/
}