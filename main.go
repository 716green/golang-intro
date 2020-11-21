package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	// fmt is common package
	fmt.Println("Hello World")

	// math package
	var absTen = math.Abs(-10)
	fmt.Println(absTen)
	fmt.Println(math.Remainder(9, 5))

	var x int
	fmt.Println(x) // will return 0

	
	// assign var while instantiating with value
	y := 5
	fmt.Println(y)
	
	x = 7
	// x = 5
	// x = 3

	if x > 6 {
		fmt.Println("More than 6")
	} else if x == 5 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	// arrays are fixed length with 0 index
	// init with 0 values
	var a [5]int
	fmt.Println(a)
	a[2] = 7
	fmt.Println(a)

	// create and init array 
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	// slices are abstractions from top of array
	c := []int{5, 4, 3, 2, 1}
	// creates new array behind scenes
	c = append(c, 13)
	fmt.Println(c)

	// maps are like objects in js or dicts in python
	vertices := make(map[string]int)
	vertices["circle"] = 0
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	// Delete value from map
	delete(vertices, "circle")
	fmt.Println(vertices)

	// Loops
	// for loop is the only loop in go
	fmt.Println("For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop 
	fmt.Println("While Loop")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}


	// Iterate through values of array
	arr := []string{"a","b","c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	// Iterate through values of map
	fmt.Println("Iterate through values")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}


	clog("Hello, World!")
	sum(9,6)

	// error handling doesn't exist in go the traditional way. This is how errors are handled
	result, err := sqrt(64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// call a struct (like a class or interface in ts)
	p := person{name: "Bob", age: 30}
	fmt.Println(p)
	fmt.Println(p.name)


	// charmap
	fmt.Println("charmap (single quotes):", 'j', "string (quotes):", "j")
	
	// pointers
	l := 7
	fmt.Println("pointers", &l)
	fmt.Println("pointers will let you modify original values instead of a copy that is discarded by using memory addresses")

	// End of main()
}

// Creating functions - outside of main()
func clog(printme string) string {
	fmt.Println(printme)
	return printme
}

func sum(q int, w int) int {
	fmt.Println(q + w)
	return q + w
}

// errors
func sqrt(e float64) (float64, error) {
	if e < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	// requires 2 values, return nil instead of error
	// fmt.Println(math.Sqrt(e), nil)
	return math.Sqrt(e), nil
}


// create a type with struct

type person struct {
	name string
	age int
}