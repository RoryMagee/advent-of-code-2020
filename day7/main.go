package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {

}

func main() {
    fmt.Println("Day 7")
    inputVals := readFile("./input")
    for _, line := range inputVals {
        fmt.Println(line)
    }
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var returnArr []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        returnArr = append(returnArr, line)
    }
    return returnArr
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
