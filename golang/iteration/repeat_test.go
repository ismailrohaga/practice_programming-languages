package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("s", 4)
	want := "ssss"

	if got != want {
		t.Errorf("Repeat() = %v, want %v", got, want)
	}

	if strings.Count(got, "s") != 4 {
		t.Errorf("Repeat() should repeat the string 4 times")
	}
}

func ExampleRepeat() {
	s := Repeat("s", 6)
	fmt.Println(s)
	// Output: ssssss
}

// to run: `go test -bench=.`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("s", 4)
	}
}
