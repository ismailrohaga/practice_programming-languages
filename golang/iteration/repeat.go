package iteration

import "fmt"

// TODO: PRACTICE EXCERCISE
// - [x] Change the test so a caller can specify how many times the character is repeated and then fix the code
// - [x] Write ExampleRepeat to document your function
// - [x] Have a look through the strings package. Find functions you think could be useful and experiment with them by writing tests like we have here. Investing time learning the standard library will really pay off over time.
func Repeat(s string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated = repeated + s
	}
	return repeated
}

// ways to declare variables in go
func declareVariables() {
	a := "initial" // a is a string
	fmt.Println(a)

	// can declare multiple variables at once.
	var b, c int = 1, 2 // b and c is an int
	fmt.Println(b, c)

	d := true // d is a bool
	fmt.Println(d)

	var e int
	fmt.Println(e) // will print 0 because e is not initialized, thus it's "zero-valued"

	f := "apple" // shorthand declaration for var f string = "apple"
	fmt.Println(f)
}
