package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	a := 1
	b := 2
	got := Calculate(a, b)
	want := a + b
	assertNumber(t, got, want)
}

// basically testable examples is just programmer making sure that the behavior must behave like how're they defined
// in the below example, i'm saying that "hey when you provide 1 and 5 as input, the output should be 6"
// and the ccode should behave like that
// it's just like a test case in leetcode or any other cp problem, here is the "expected output"
// the name should be "ExampleXYZ"
func ExampleCalculate() {
	sum := Calculate(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertNumber(t *testing.T, got int, want int) {
	t.Helper() // a flag to indicate that the function is a helper function
	if got != want {
		// instead of using %q, we can use %d to print the integer
		t.Errorf("Got %d want %d", got, want)
	}
}
