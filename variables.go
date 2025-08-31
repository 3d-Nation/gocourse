package main

import (
	"fmt"
)

// global
var globalName = "woozy"

func main() {
	var age int
	var name string = "Jobob"
	var string2 = "jehozabad"

	age = 30
	age2 := 35
	lastName := "Smoby"
	globalName := "notwoozy" //this version of globalName shadows the global variable...no way to access the global version of it in main now

	fmt.Println("Hello, my name is", name, lastName, "and I am", age2, "years old.")
	fmt.Println(age, string2)

	fmt.Println("My global name is", globalName)
	// fmt.Println("My real global name is", main.globalName)

}
