package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    PREAMBLE_LEN := 25
    inputVals := readFile("./input")
    prevIndex := 0
    for i := PREAMBLE_LEN; i < len(inputVals); i++ {
        preamble := inputVals[prevIndex: prevIndex + PREAMBLE_LEN]
        target := inputVals[i]
        res := sumExists(preamble, target)
        if !res {
            fmt.Println("Finished", target)
            break
        }
        prevIndex++
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

func check(e error) {
    if (e != nil) {
        panic(e)
    }
}

func sumExists(arr []int, target int) bool {
    for i := 0; i < len(arr); i++ {
        for j := i+1; j < len(arr); j++ {
            if arr[i] + arr[j] == target {
                return true
            }
        }
    }
    return false
}
