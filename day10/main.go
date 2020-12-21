package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "time"
)

func main() {
    inputVals := readFile("./testinput")
    start := time.Now()
    solution1 := countNumberOfLinks(append([]int(nil), inputVals...))
    finish := time.Since(start)
    solution2 := countCombinations(append([]int(nil), inputVals...))
    fmt.Println("Solution 1", solution1, finish)
    fmt.Println("Solution 2", solution2)
}

func countCombinations(inputVals []int) int {
    sortArr(inputVals)
    return 1
}

func countNumberOfLinks(inputVals []int) int {
    sortArr(inputVals)
    inputVals = append(inputVals, inputVals[len(inputVals)-1] + 3)
    inputVals = append([]int{0}, inputVals...)
    diffIsOne := 0
    diffIsThree := 0
    for i := 0; i < len(inputVals)-1; i++ {
        curr := inputVals[i]
        next := inputVals[i+1]
        diff := next - curr
        if diff == 1 {
            diffIsOne++
        } else if diff == 3 {
            diffIsThree++
        }
    }
    return diffIsOne * diffIsThree
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sortArr(inputVals []int) {
    for x := 1; x  < len(inputVals); x++ {
        current := inputVals[x]
        for y :=x-1; y >= 0 && inputVals[y] > current; y-- {
            inputVals[y+1] = inputVals[y]
            inputVals[y] = current
        }
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