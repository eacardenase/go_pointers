package main

import "fmt"

func main() {
	age := 28
	adultYears := getAdultYears(age)

	fmt.Println("Age: ", age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
