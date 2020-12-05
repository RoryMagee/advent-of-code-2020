package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    fmt.Println("Solution 1")
    inputVals := readFile("./input")
    validPasses := 0
    for i := 0; i < len(inputVals); i++ {
        pass := inputVals[i]
        fmt.Println(pass)

    }
}

func splitPass(passString string) rule string, char string, password string {
    split
}

func validPass(rule string, char string, password string) bool {
    return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val := scanner.Text()
        lines = append(lines, val)
    }
    return lines
}
