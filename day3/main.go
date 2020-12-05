package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    inputVals := readFile("./input")
    startPosition := [2]int{0, 0}
    // startPosition[0] = vertical
    // startPosition[1] = horizontal
    // inputvals[row][column]
    treeCount := 0
    for startPosition[0] < len(inputVals) {
        if inputVals[startPosition[0]][startPosition[1]] == "#" {
            treeCount++
        }
        // Increment positions
        diff := len(inputVals[1]) - (startPosition[1] + 3)
        if diff <= 0 {
            startPosition[1] = abs(diff)
        } else {
            startPosition[1] = startPosition[1] + 3
        }
        startPosition[0] = startPosition[0] + 1
        fmt.Println("startPosition[0]", startPosition[0], "startPosition[1]",
        startPosition[1])
    }
    fmt.Println(treeCount)
}

func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
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
