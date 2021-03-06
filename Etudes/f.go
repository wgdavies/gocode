package main

import "fmt"

type Point struct{ x, y int }

func PointToString(p Point) string {
    return fmt.Sprintf("Point{%d, %d}", p.x, p.y)
}

func main() {
    p := Point{3, 5}
    fmt.Println(PointToString(p)) // static dispatch
    fmt.Println(p)
}

