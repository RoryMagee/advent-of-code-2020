package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    inputVals := readFile("./input")
    highest := 0
    var allSeatIds []int
    for _, seat := range inputVals {
        id := getSeatId(seat)
        allSeatIds = append(allSeatIds, id)
        if id > highest {
            highest = id
        }
    }
    /* 
    * At this point we will have all the seat ID's
    * We will need to sort the array and then iterate through and find the 
    * missing value
    */
    sortedIds := sortArr(allSeatIds)
    for i := 1; i < len(sortedIds) - 1; i++ {
        if sortedIds[i - 1] != sortedIds[i] - 1 {
            fmt.Println("solution 2:", sortedIds[i] - 1)
            break
        }
        if sortedIds[i + 1] != sortedIds[i] + 1 {
            fmt.Println("solution 2:", sortedIds[i] + 1)
            break
        }
    }
    fmt.Println(highest)
}

func sortArr(arr []int) []int {
    for i := 1; i < len(arr); i++ {
        curr := arr[i]
        for j := i - 1; j >= 0 && arr[j] > curr; j-- {
            arr[j + 1] = arr[j]
            arr[j] = curr
        }
    }
    return arr
}

func getSeatId(seat string) int {
    row := weirdBinarySearchTing(127, seat[:7], "F", "B")
    col := weirdBinarySearchTing(7, seat[7:], "L", "R")
    result := (row * 8) + col
    return result
}

func weirdBinarySearchTing(max int, seq string, upperChar string, lowerChar string) int {
    min := 0
    for _, c := range seq { // Loop through seq
        /*
         * On each iteration of the loop if c points to the lower half then
         * we keep min the same and assign max as the halfway point between min
         * and max. Otherwise we keep max the same and assign min as the halfway
         * point between current min and max
         */
        val := string(c)
        if val == upperChar {
            max = (max + min) / 2
        } else {
            min = (max + min) / 2
        }
    }
    return max
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
