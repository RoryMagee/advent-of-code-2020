package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    inputVals := readFile("./input")
    highest := 0
    for _, seat := range inputVals {
        row := weirdBinarySearchTing(127, seat[:7], "F", "B")
        col := weirdBinarySearchTing(7, seat[7:], "L", "R")
        result := (row * 8) + col
        if result > highest {
            highest = result
        }
    }
    fmt.Println(highest)
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
