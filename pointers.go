package main

import "fmt"

func main() {
	age := 28
	agePointer := &age

	fmt.Println("Age Address:", agePointer) // 0x1400009a02
	fmt.Println("Age Value:", *agePointer)  // 28

	editAgeToAdultYears(agePointer)

	fmt.Println("Age New Value:", age)
}

func editAgeToAdultYears(age *int) {
	// *age = *age - 18
	*age -= 18
}
