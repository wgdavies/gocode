package main

import (
    "fmt"
    "os"
)

func main() { /*
    const hello = "Hello"
    f := "Fizz"
    var b string = "Buzz"
*/
    for i := 1; i <= 15; i++ {
        fmt.Printf("%3d ", i)

        if i % 3 == 0 {
            fmt.Printf("Fizz")
        }
        if i % 5 == 0 {
            fmt.Printf("Buzz")
        }

        fmt.Printf("\n")
    }
    var b bool = false

    fmt.Printf("%T: %v\n", b, b)
    fmt.Printf("%T: %v\n", 0, 0)
    fmt.Printf("%T: %v\n", 0.0, 0.0)
    fmt.Printf("%T: %v\n", 'x', 'x')
    fmt.Printf("%T: %v\n", "x", "x")
    fmt.Printf("%T: %v\n", 0i, 0i)
    os.Exit(0)
}
