package main

import (
	"fmt"
	"strconv"
)

var NUM1 int = 999
var AKIN string = "Akin Unver"

func main() {

	fmt.Println("Hello World!")

	var i int = 30
	j := 25
	k := 99.
	//var j float32 = 27
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T", j, j)

	//Multiple var
	var (
		firstName string = "Akin"
		lastName  string = "Unver"
		age       int    = 24
	)
	fmt.Println(firstName, lastName, age)

	//Scope
	fmt.Println(NUM1)
	NUM1 = 0
	fmt.Println(NUM1)

	//Conversion (import strconv)
	var x int = 32
	fmt.Printf("%v, %T\n", x, x)

	var y string = strconv.Itoa(x)
	fmt.Printf("%v, %T\n", y, y)
}
