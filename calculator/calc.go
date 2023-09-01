package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalAge () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name: ")
	name, _ := reader.ReadString('\n')
	cleanName := strings.Replace(name, "\n", "", -1)

	fmt.Println("Enter current year: ")
	year, _ := reader.ReadString('\n')
	cleanYear := strings.Replace(year, "\n", "", -1)

	fmt.Println("Enter your date of birth: ")
	dob, _ := reader.ReadString('\n')
	cleanDOB := strings.Replace(dob, "\n", "", -1)

	// remove spaces
	cleanName = strings.TrimSpace(cleanName)
	cleanYear = strings.TrimSpace(cleanYear)
	cleanDOB = strings.TrimSpace(cleanDOB)

	// convert to numbers
	convYear, _ := strconv.ParseInt(cleanYear, 10, 64)
	convDOB, _ := strconv.ParseInt(cleanDOB, 10, 64)

	//  calculate age
	var age int64 = convYear - convDOB

	fmt.Printf("name: %v \n", cleanName)
	fmt.Printf("year :%v \n", convYear)
	fmt.Printf("dob: %v \n", convDOB)
	fmt.Println("")
	fmt.Printf("Your name is %v and you are %v years old.\n", cleanName, age)
}
