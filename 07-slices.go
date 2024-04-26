package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, slices are typed only by the elements they contain,
	// not the number of elements.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Set and get values just like arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[1])

	fmt.Println("len:", len(s))

	// Append returns a new slice with a new value(s) appended.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// Slices can also be copied.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Slices support index slicing, similar to Python with the following syntax:
	// slice[low, high]
	// Low is inclusive, high is exclusive.
	l := s[2:5]
	fmt.Println("slice:", l)

	// This slices up to (but excluding) s[5]:
	l = s[:5]
	fmt.Println("slice:", l)

	// This slices up from (and including) s[2]:
	l = s[2:]
	fmt.Println("slice:", l)

	// We can declare and initialize a slice in one line as well:
	t := []string{"g", "h", "i"}
	fmt.Println("declare:", t)

	// The slices package provides some useful tools:
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be constructed into multi-dimensional data structures.
	// The length of inner slices can vary, unlike with arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
