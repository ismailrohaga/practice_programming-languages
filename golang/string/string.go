package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "this is A string"
	fmt.Println(strings.ToUpper(ss))
	fmt.Println(ss) // it doesnt change the original string
	fmt.Println(strings.ToLower(ss))
	fmt.Println(strings.HasPrefix(ss, "this"))
	fmt.Println(strings.HasPrefix(ss, "is"))
	fmt.Println(strings.HasSuffix(ss, "string"))
	fmt.Println(strings.Contains(ss, "is"))
	fmt.Println(strings.Contains(ss, "A"))
	fmt.Println(strings.Contains(ss, "a")) // will return false, case sensitive
	fmt.Println(strings.Count(ss, "A"))
	fmt.Println(len(ss))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))

	sp := strings.Split(ss, " ")
	fmt.Printf("%q", sp)

	fmt.Println()

	fields := strings.Fields(ss)
	fmt.Printf("%q", fields)

	fmt.Println()
	fmt.Println(strings.ReplaceAll(ss, "is", "was"))
	// Output: thwas was A string

	fmt.Println(strings.Index(ss, "is")) // return 2
	fmt.Println(strings.Index(ss, "A"))  // return 8
}
