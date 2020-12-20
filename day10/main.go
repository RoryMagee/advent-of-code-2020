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
    sortArr(inputVals)
    oneJump := 0
    threeJump := 0
    for i := 0; i < len(inputVals)-1; i++ {
        if inputVals[i+1] - inputVals[i] == 1 {
            oneJump++
        } else {
            threeJump++
        }
    }
    fmt.Println(oneJump * threeJump)
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
