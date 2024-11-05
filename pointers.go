package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	// var agePointer *int // declaration
	// agePointer = &age   // instantiation

	fmt.Println("Age Address:", agePointer) // 0x1400009a02
	fmt.Println("Age Value:", *agePointer)  // 28

	// adultYears := getAdultYears(&age)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
