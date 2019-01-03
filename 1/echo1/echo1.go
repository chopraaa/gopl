/* To execute the program:
   $ go run echo1.go

   To fetch the source code and execute the program:
   $ go get github.com/gopl/1/echo1

   Compile it once and save the compiled result:
   $ go build echo1.go
*/

// Each source file begins with a package declaration that states which package
// the file belongs to
// The package "main" is special as it defines a standalone executable program,
// not a library
package main

// Package is followed by the other packages it imports
import (
        // The "fmt" package contains functions for printing formatted output and scanning
        // input
        "fmt"
        // The "os" package provides functions and other values for dealing with the
        // operating system in a platform-independent fashion
        "os"
)

// And then the declarations of the program that are stored in that file
// The function main() is special - it's where execution of the program begins
func main() {
        var s, sep string
        // The variable os.Args is a slice of strings.
        // The first element of os.Args, os.Args[0] is the name of the command itself
        // The other elements of are the arguments that were presented to the program
        // when it started execution
        for i := 1; i < len(os.Args); i++ {
                // Creates a string "s" that adds args from os.Args as it increments
                // First iteration: arg1, second iteration: arg1 arg2 and so on...
                s += sep + os.Args[i]
                sep = " "
        }
        // Println is one of the basic output functions in "fmt"; it prints one or
        // more values, separated by spaces, with a newline character at the end so
        // that the values appear as a single line of output
        // fmt.Println(s)
        // Print the final value of "s"
        fmt.Println(s)
}
