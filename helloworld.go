package main

import "fmt"

func main() {
	var name string

	name = "Raditya"
	fmt.Println(name)

	name = "Raditya Dwisatrio"
	fmt.Println(name)

	var age = 10 // Integer
	fmt.Println(age)

	var isMarried = true // Boolean
	fmt.Println(isMarried)

	// Tanpa menggunakan var
	secondName := "Budi"
	fmt.Println(secondName)

	secondName = "Budi Anduk"
	fmt.Println(secondName)

	// Multiple variable
	var (
		firstName = "Tono"
		lastName  = "Santoso"
	)
	fmt.Println(firstName, lastName)

	// Konstan

	const nilaiTetap int = 100
	fmt.Println(nilaiTetap)

}
