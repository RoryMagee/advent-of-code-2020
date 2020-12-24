package main

import (
	"bufio"
	"fmt"
	"os"
    "time"
    "strings"
)

func main() {
    fmt.Println("Day11")
    inputVals := readFile("./input")
    start := time.Now()
    for true {
        res := applyRules(inputVals)
        if res == false {
            fmt.Println(countOccupied(inputVals))
            break
        }
    }
    fmt.Println(time.Since(start))
}

func countOccupied(plan [][]string) int {
    count := 0
    for i := 0; i < len(plan); i++ {
        for j := 0; j < len(plan[i]); j++ {
            if plan[i][j] == "#" {
                count = count + 1
            }
        }
    }
    return count
}

func applyRules(inputVals [][]string) bool {
    hasChanged := false
    plan := make([][]string, len(inputVals))
    for x := range inputVals {
        plan[x] = make([]string, len(inputVals[x]))
        copy(plan[x], inputVals[x])
    }
    for i := 0; i < len(plan); i++ {
        for j := 0; j < len(plan[0]); j++ {
            count := 0
            if plan[i][j] != "." {
                //Check above
                if i-1 >= 0 {
                    if plan[i-1][j] == "#" {
                        count++
                    }
                }
                //Check below
                if i+1 < len(plan) {
                    if plan[i+1][j] == "#" {
                        count++
                    }
                }
                //Check left
                if j-1 >= 0 {
                    if plan[i][j-1] == "#" {
                        count++
                    }
                }
                //Check right
                if j+1 < len(plan[0]) {
                    if plan[i][j+1] == "#" {
                        count++
                    }
                }
                //Check top left
                if i-1 >=0 && j-1 >=0 {
                    if plan[i-1][j-1] == "#" {
                        count++
                    }
                }
                //Check top right
                if i-1 >= 0 && j+1 < len(plan[0]) {
                    if plan[i-1][j+1] == "#" {
                        count++
                    }
                }
                //Check bottom left
                if i+1 < len(plan) && j-1 >= 0 {
                    if plan[i+1][j-1] == "#" {
                        count++
                    }
                }
                //Check bottom right
                if i+1 < len(plan) && j+1 < len(plan[0]) {
                    if plan[i+1][j+1] == "#" {
                        count++
                    }
                }
                //Update Node
                if count == 0 && inputVals[i][j] != "#" {
                    inputVals[i][j] = "#"
                    hasChanged = true
                }
                if count >= 4 && inputVals[i][j] != "L" {
                    inputVals[i][j] = "L"
                    hasChanged = true
                }
            }
        }
    }
    return hasChanged
}

func printPlan(plan [][]string) {
    fmt.Println("-------------")
    for i := 0; i < len(plan); i++ {
        fmt.Println(plan[i])
    }
}

func readFile(path string) [][]string {
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var returnArr [][]string
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), "")
        returnArr = append(returnArr, line)
    }
    return returnArr
}
