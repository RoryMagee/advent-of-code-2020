package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile() []int {
    file, err := os.Open("./input")
    check(err)
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val, err := strconv.Atoi(scanner.Text())
        check(err)
        lines = append(lines, val)
    }
    return lines
}

func main() {
    inputVals := readFile()
    fmt.Println(inputVals)
}
