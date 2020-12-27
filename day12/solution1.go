package main

import (
    "fmt"
)

func solution1() {
    fmt.Println("Solution 1")
    inputVals := readFile("./testinput")
    for i := 0; i < len(inputVals); i++ {
        fmt.Println(inputVals[i])
    }
}
