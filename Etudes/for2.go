package main

import "fmt"

func main() {
    for i := 0; i <= 10; i++ {
        switch i {
        case 0: fmt.Printf("Zero ")
        case 1: fmt.Printf("One ")
        case 2: fmt.Printf("Two ")
        case 3: fmt.Printf("Three ")
        case 4: fmt.Printf("Four ")
        case 5: fmt.Printf("Five ")
        case 6: fmt.Printf("Six ")
        case 7: fmt.Printf("Seven ")
        case 8: fmt.Printf("Eight ")
        default: fmt.Printf("%d ", i)
        }
    }

    fmt.Printf("\n")
}

