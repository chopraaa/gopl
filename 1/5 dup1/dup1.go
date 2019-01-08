package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    /*  A map holds a set of key/value pairs and provides constant-time operations to store, retrieve, or test for an item in the set
        The key may be of any type whose values can be compared with ==, strings being the most common example
        The value may be of any type at all
    */
    // In this map, the keys are strings and the values are integers
    // The builtin function "make" creates an empty map (it has other uses too)
    counts := make(map[string]int)
    // Each type dup reads a line of input, the line is used as a key into the map and the corresponding value is incremented
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
        /*  This statement is equivalent to:

            line := input.Text()
            counts[line] = counts[line] + 1
        */
    }
    // Note: ignoring potential errors from input.Err()
    for line, n := range counts {
        /*  As with for, paranthesis are never used around the condition in an if statement,
            but braces are required for the body
            There can be an optional "else" part that is executed if the condition is false
        */
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
