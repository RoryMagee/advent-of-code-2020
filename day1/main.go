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
    for i := 0; i < len(inputVals); i++ {
        for j := 1; j < len(inputVals); j++ {
            if (inputVals[i] + inputVals[j] == 2020) {
                fmt.Println(inputVals[i] * inputVals[j])
            }
        }
    }
}
