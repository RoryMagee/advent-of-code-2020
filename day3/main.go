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
    solution1Rule := [2]int{1, 3}
    solution2Rules := [5][2]int {
        {1, 1},
        {1, 3},
        {1, 5},
        {1, 7},
        {2, 1},
    }
    solution2Answer := 1
    solution1Answer := countTrees(solution1Rule, startPosition, inputVals)
    for _, r := range solution2Rules {
        solution2Answer = solution2Answer * countTrees(r, startPosition, inputVals)
    }
    fmt.Println("Solution1:", solution1Answer)
    fmt.Println("Solution2:", solution2Answer)
}

func countTrees(rule [2]int, startPosition [2]int, inputVals [][]string) int {
    treeCount := 0
    for startPosition[0] < len(inputVals) {
        if inputVals[startPosition[0]][startPosition[1]] == "#" {
            treeCount++
        }
        // Increment positions
        diff := len(inputVals[1]) - (startPosition[1] + rule[1])
        if diff <= 0 {
            startPosition[1] = abs(diff)
        } else {
            startPosition[1] = startPosition[1] + rule[1]
        }
        startPosition[0] = startPosition[0] + rule[0]
    }
    return treeCount
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
