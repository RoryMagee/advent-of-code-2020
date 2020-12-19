package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    fmt.Println("Day 10")
    inputVals := readFile("./testinput")
    for _, line := range inputVals {
        fmt.Println(line)
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(path string) []int {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var returnArr []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value, _ := strconv.Atoi(scanner.Text())
        returnArr = append(returnArr, value)
    }
    return returnArr
}
