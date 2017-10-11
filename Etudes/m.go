package main

import "fmt"

type Point struct{ x, y int }

func (p Point) String() string {
    return fmt.Sprintf("Point{%d, %d}", p.x, p.y)
}

func main() {
    p := Point{3, 5}
    fmt.Println(p.String())       // static dispatch
    fmt.Println(p)
}

