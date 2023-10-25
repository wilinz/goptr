package main

import (
	"fmt"
	"github.com/wilinz/goptr"
)

func Example() {
	// Use the ptr function to create a pointer to an integer
	pt := goptr.Ptr(42)
	fmt.Println(pt) // The address of the output pointer

	// Use the unptr function to dereference pointers
	value := goptr.Unptr(pt)
	fmt.Println(value) // The value pointed to by the output pointer
}

func main() {
	Example()
}
