/* To execute the program:
   $ go run hello_world.go

   To fetch the source code and execute the program:
   $ go get github.com/gopl/1/hello_world

   Compile it once and save the compiled result:
   $ go build hello_world.go
*/

// Each source file begins with a package declaration that states which package
// the file belongs to
// The package "main" is special as it defines a standalone executable program,
// not a library
package main

// Package is followed by the other packages it imports
// The "fmt" package contains functions for printing formatted output and scanning
// input
import "fmt"

// And then the declarations of the program that are stored in that file
// The function main() is special - it's where execution of the program begins
func main() {
    // Println is one of the basic output functions in "fmt"; it prints one or
    // more values, separated by spaces, with a newline character at the end so
    // that the values appear as a single line of output
    fmt.Println("Hello, ਸੰਸਾਰ")
}
