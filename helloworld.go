package main

import "fmt"

func main() {
	var nilai32 int32 = 32678
	var nilai64 int64 = int64(nilai32) // Success
	var nilai8 int8 = int8(nilai32)    // Fail - melebihi batas maksimum nilai int8

	fmt.Println(nilai32) // 32678
	fmt.Println(nilai64) // 32678
	fmt.Println(nilai8)  // -90

	var name = "Radit"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)    // Radit
	fmt.Println(e)       // 82
	fmt.Println(eString) // R
}
