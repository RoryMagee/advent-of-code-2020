package main

import (
    "fmt"
)

func main() {
    fmt.Println("day 9")
    arr := make([]int, 10)
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    arr[3] = 4
    arr[4] = 5
    for _, i := range arr {
        fmt.Println(i)
    }
}
