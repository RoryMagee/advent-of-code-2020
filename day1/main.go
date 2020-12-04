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

func main() {
    file, err := os.Open("./input")
    check(err) // Check for error reading input file
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val, err := strconv.Atoi(scanner.Text())
        check(err)
        lines = append(lines, val)
    }
    fmt.Println(lines)
}
