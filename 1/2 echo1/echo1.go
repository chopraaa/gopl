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
        /*  The "fmt" package contains functions for printing formatted output and scanning
            input
        */
        "fmt"
        /*  The "os" package provides functions and other values for dealing with the
            operating system in a platform-independent fashion
        */
        "os"
)

// And then the declarations of the program that are stored in that file
// The function main() is special - it's where execution of the program begins
func main() {
        /*  Declare two variables "s" and "sep" of type string. A variable can be
            intialized as part of it's declaration. Here, we implicityly intialize
            "s" and "sep" to empty strings
        */
        var s, sep string
        /*  The variable os.Args is a slice of strings.
            The first element of os.Args, os.Args[0] is the name of the command itself
            The other elements of are the arguments that were presented to the program
            when it started execution
        */

        /*  The "for" loop is the only loop statement in Go. It has a number of forms,
            one of which is illustrated here:

            for intialiation; condition; post {
                // zero or more statements
            }

            Infinite loops can be created:

            "intialization" must be a "simple" statement i.e.
                - a short variable declaration,
                - an assignment statement
                - a function call

            "condition" is a boolean expression evaluated at the beginning of each
            iteration of the loop, if "true", statements controlled by the loop are executed

            "post" statement is executed after the body of the loop, then the condition is evaluated again

            The loop ends when the condition becomes false

            A traditional while loop should contain only "condition":

            for condition {
                ...
            }

            Or a traditional infinite loop:

            for {
                ...
            }
        */
        // Creates a string "s" that adds args from os.Args as it increments
        /*  := is part of a "short variable declaration", a statement that declares
            one or more one or more variables and gives them appropriate types based
            on the initializer values
        */
        for i := 1; i < len(os.Args); i++ {
                /*  For numbers, Go provides the usual arithmetic and logical operators
                    When applied to strings however, the + operator concatenates the values
                */
                s += sep + os.Args[i]
                sep = " "
        }
        /*  Println is one of the basic output functions in "fmt"; it prints one or
            more values, separated by spaces, with a newline character at the end so
            that the values appear as a single line of output
        */
        // Print the final value of "s"
        fmt.Println(s)
}
