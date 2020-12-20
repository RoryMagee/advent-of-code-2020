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
    solution1 := findVulnerability(inputVals, PREAMBLE_LEN)
    solution2 := solution2(inputVals, solution1)
    fmt.Println("solution1", solution1)
    fmt.Println("solution2", solution2)
}

func solution2(inputVals []int, target int) int {
    for i := 0; i < len(inputVals); i++ {
        for j := i + 1; j < len(inputVals); j++ {
            arr := inputVals[i:j]
            if sumSlice(arr) == target {
                return getLargest(arr) + getSmallest(arr)
            }
        }
    }
    return -1
}

func findVulnerability(inputVals []int, PREAMBLE_LEN int) int{
    prevIndex := 0
    returnVal := -1
    for i := PREAMBLE_LEN; i < len(inputVals); i++ {
        preamble := inputVals[prevIndex:prevIndex + PREAMBLE_LEN]
        target := inputVals[i]
        res := sumExists(preamble, target)
        if !res {
            // Answer found
            returnVal = target
            break
        }
        prevIndex++
    }
    return returnVal
}

func getLargest(slice []int) int {
    var largest int
    for _, num := range slice {
        if num > largest {
            largest = num
        }
    }
    return largest
}

func getSmallest(slice []int) int {
    smallest := slice[0]
    for _, num := range slice[1:] {
        if num < smallest {
            smallest = num
        }
    }
    return smallest
}

func sumSlice(slice []int) int {
    total := 0
    for _, num := range slice {
        total = total + num
    }
    return total
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
