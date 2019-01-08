package main

import (
        "fmt"
        "os"

        // required for the strings.Join function
        "strings"
)

func main() {
        /*  Method 1:
            - Short variable declaration
            - It is the most compact
            - It can be used only within a function, not for package-level variables
            s := ""

            Method 2:
            - Relies on default initialization to the zero value for strings, which is ""
            var s string


            Method 3:
            - Rarely used except when using multiple variables
            var s = ""

            Method 4:
            - Explicit about the variable's type, which is redundant when it is
              the same as that of the initial value but necessary in other cases where they are not of the same type
            var s string = ""
        */

        s, sep := "", ""

        /*  Range produces a pair of values:
            - index
            - value of the element at that index

            Since we don't need the index, we use the blank identifier denoted by _ (underscore)
            The blank identifier may be used whenever syntax requires a variable name but program logic does not
        */

        for _, arg := range os.Args[1:] {
                s += sep + arg
                sep = " "
        }
        fmt.Println(s)

        /* This is an expensive operation. Use the Join function of the strings package for large amounts of data */

        fmt.Println(strings.Join(os.Args[1:], " "))

        // ...or just...

        fmt.Println(os.Args[1:])

        // prints [arg1 arg2]
        // yes, with []
}
