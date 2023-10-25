# goptr

`goptr` is a Go package that provides utility functions for working with pointers. The package allows you to create pointers to non-addressable values and dereference pointers to obtain their underlying values.

## Installation

To use `goptr`, you need to have Go installed and set up on your machine. The package requires Go version 1.18 or higher due to its use of generics. You can check your Go version by running `go version` in your terminal. If you have an older version of Go, please update it before proceeding.

To install the package, use the `go get` command:

```bash
go get github.com/wilinz/goptr
```

## Usage

Import the `goptr` package in your Go program to use the provided utility functions:

```go
import (
"fmt"
"github.com/wilinz/goptr"
)

type Person struct {
Name   *string
Age    *int
Height *float64
}

func main() {

// Create a Person struct with pointers to basic data types
person := Person{
Name:   goptr.Ptr("John Doe"),
Age:    goptr.Ptr(30),
Height: goptr.Ptr(1.75),
}

// Dereference the pointers to access the values
fmt.Println(*person.Name)   // John Doe
fmt.Println(*person.Age)    // 30
fmt.Println(*person.Height) // 1.75
}
```

In the example above, we define a `Person` struct that contains pointers to basic data types such as `string`, `int`, and `float64`. By using the `Ptr` function from the `goptr` package, we can create pointers to these values and assign them to the corresponding fields in the `Person` struct.

The benefit of using `goptr` in this scenario is that it allows us to work with pointers to basic data types that are not normally addressable. By making these values addressable using `goptr.Ptr`, we can store them as pointers in a struct or pass them as arguments to functions that expect pointer values. This can be useful in situations where you need to work with libraries or APIs that require pointer values, or when you want to modify the underlying values indirectly through their pointers.

## Function Reference

### `Ptr[T any](v T) *T`

The `Ptr` function takes a value `v` of any type `T` and returns a pointer to that value. It allows you to obtain the address of non-addressable values, such as constants, literals, or the return values of functions or methods.

### `Unptr[T any](ptr *T) T`

The `Unptr` function takes a pointer `ptr` to a value of type `T` and returns the underlying value. It allows you to dereference pointers and obtain the value they point to.

## Why Use `goptr`?

Using `goptr` provides the advantage of obtaining pointers directly from literals without the need to define separate variables and take their addresses. Let's consider the following code as an example of the conventional approach:

```go
import (
	"fmt"
	"github.com/wilinz/goptr"
)

type Person struct {
	Name   *string
	Age    *int
	Height *float64
}

func main() {
	name := "John Doe"
	age := 30
	height := 1.75

	// Create a Person struct with direct values
	person := Person{
		Name:   &name,
		Age:    &age,
		Height: &height,
	}

	// Access the values directly
	fmt.Println(*person.Name)   // John Doe
	fmt.Println(*person.Age)    // 30
	fmt.Println(*person.Height) // 1.75
}
```

In the code above, we define variables `name`, `age`, and `height` to hold the values we want to store in the `Person` struct. We then take the address of each variable and assign the pointers to the corresponding fields in the `Person` struct. While this approach works, it requires additional variable declarations and explicit address-taking, which can be cumbersome and less readable.

By using `goptr`, we can simplify the code and avoid the need for separate variables:

```go
import (
	"fmt"
	"github.com/wilinz/goptr"
)

type Person struct {
	Name   *string
	Age    *int
	Height *float64
}

func main() {

	// Create a Person struct with pointers to basic data types
	person := Person{
		Name:   goptr.Ptr("John Doe"),
		Age:    goptr.Ptr(30),
		Height: goptr.Ptr(1.75),
	}

	// Dereference the pointers to access the values
	fmt.Println(*person.Name)   // John Doe
	fmt.Println(*person.Age)    // 30
	fmt.Println(*person.Height) // 1.75
}
```

In this updated code, we use the `goptr.Ptr` function to directly obtain pointers from the literals `"John Doe"`, `30`, and `1.75`. This eliminates the need for separate variable declarations and explicit address-taking, resulting in cleaner and more concise code.

The `goptr` package provides a convenient and efficient way to work with pointers to non-addressable values, allowing you to simplify your code and improve readability.

## Contribution

Contributions to `goptr` are welcome! If you find a bug, have a suggestion, or want to contribute new features or improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/wilinz/goptr).

## License

`goptr` is licensed under the [MIT License](https://github.com/wilinz/goptr/blob/main/LICENSE). Feel free to use, modify, and distribute the package according to the terms of the license.

## Acknowledgments

The `goptr` package was created by [wilinz](https://github.com/wilinz) and inspired by the need to work with non-addressable values and pointers in Go.

If you have any questions or need further assistance, please don't hesitate to reach out. Happy coding!