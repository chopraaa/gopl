package main

import (
        "fmt"
        "os"
        "strings"
)

func main() {
        // Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command
        // that invoked it.
        fmt.Println(strings.Join(os.Args[0:], " "))

        // Exercise 1.2: Modify the echo program to print the index and value of each of its arguments,
        // one per line.
        for i, s := range os.Args[1:] {
                fmt.Println(i, " ", s)
        }

        // Exercise 1.3: Experiment to measure the difference in running time between our potentially
        // inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the
        // time package, and Section 11.4 shows how to write benchmark tests for systematic performance
        // evaluation.)
}
