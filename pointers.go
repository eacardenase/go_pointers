package main

import "fmt"

func main() {
	age := 28
	var agePointer *int // declaration
	agePointer = &age   // instantiation

	fmt.Println("Age Address:", agePointer) // 0x1400009a02
	fmt.Println("Age Value:", *agePointer)  // 28

	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
