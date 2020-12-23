package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func main() {
    fmt.Println("Day11")
    inputVals := readFile("./testinput")
    applyRules(inputVals)
    for true {
        res := applyRules(inputVals)
        if res {
            fmt.Println(inputVals)
            fmt.Println(countOccupied(inputVals))
            break
        }
    }
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
    plan := append([][]string(nil), inputVals...)
    for i := 0; i < len(plan); i++ {
        for j := 0; j < len(plan[0]); j++ {
            if plan[i][j] != "." {
                count := 0
                //Check above
                if i-1 >= 0 {
                    if plan[i-1][j] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check below
                if i+1 < len(plan) {
                    if plan[i+1][j] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check left
                if j-1 >= 0 {
                    if plan[i][j-1] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check right
                if j+1 < len(plan[0]) {
                    if plan[i][j+1] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check top left
                if i-1 >=0 && j-1 >=0 {
                    if plan[i-1][j-1] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check top right
                if i-1 >= 0 && j+1 < len(plan[0]) {
                    if plan[i-1][j+1] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check bottom left
                if i+1 < len(plan) && j-1 >= 0 {
                    if plan[i+1][j-1] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Check bottom right
                if i+1 < len(plan) && j+1 < len(plan[0]) {
                    if plan[i+1][j+1] == "#" {
                        count++
                        hasChanged = true
                    }
                }
                //Update Node
                if count == 0 {
                    plan[i][j] = "#"
                }
                if count >= 4 {
                    plan[i][j] = "L"
                }
            }
        }
        //fmt.Println(plan[i])
    }
    inputVals = plan
    return hasChanged
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
