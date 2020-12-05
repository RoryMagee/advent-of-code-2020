package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    inputVals := readFile("./input")
    for _, v := range inputVals {
        fmt.Println(v)
    }
}

func readFile(path string) [][]string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var lines [][]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val := strings.Split(scanner.Text(), "")
        lines = append(lines, val)
    }
    return lines
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
