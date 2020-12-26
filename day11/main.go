package main

import (
    "time"
	"bufio"
	"fmt"
	"os"
    "strings"
)

type direction int
const(
    left = -1
    right = 1
    up = -1
    down = 1
    still = 0
)

func main() {
    fmt.Println("Day11")
    //inputVals := readFile("./testinput")
    //for true {
    //    res := applyRulesS1(inputVals)
    //    if res == false {
    //        fmt.Println(countOccupied(inputVals))
    //        break
    //    }
    //}
    inputVals := readFile("./testinput")
    walkDirection2(inputVals, 9, 9, -1, -1)
    //for true {
    //    printPlan(inputVals)
    //    res := applyRulesS2(inputVals)
    //    if res == false {
    //        fmt.Println(countOccupied(inputVals))
    //        break
    //    }
    //}
}

func walkDirection2(input [][]string, i int, j int, iDirection int, 
jDirection int) int {
    retValue := 0
    for (i + iDirection < len(input) && i >= 0)  && (j + jDirection < len(input[0]) && j >= 0){
        i = i + iDirection
        j = j + jDirection
        fmt.Printf("[%d][%d]\n", i, j)
    }
    return retValue
}

func applyRulesS2(inputVals [][]string) bool {
    // [vertical][horizontal]
    left := []int{0, -1}
    right := []int{0, 1}
    up := []int{-1, 0}
    down := []int{1, 0}
    topleft := []int{-1, -1}
    topright := []int{-1, 1}
    botleft := []int{1, -1}
    botright := []int{1, 1}
    directions := [][]int{left, right, up, down, topleft, topright, botleft, botright}
    hasChanged := false
    plan := make([][]string, len(inputVals))
    for x := range inputVals {
        plan[x] = make([]string, len(inputVals[x]))
        copy(plan[x], inputVals[x])
    }
    for i := 0; i < len(plan); i++ {
        for j := 0; j < len(plan[0]); j++ {
            if plan[i][j] != "." {
                count := 0
                for i := 0; i < len(directions); i++ {
                    curr := directions[i]
                    count = count + walkDirection(plan, i, j, curr[0], curr[1])
                    //fmt.Printf("[%d][%d] = %d\n", i, j, count)
                }
                if count == 0 && inputVals[i][j] != "#" {
                    inputVals[i][j] = "#"
                    hasChanged = true
                }
                if count >= 5 && inputVals[i][j] != "L" {
                    inputVals[i][j] = "L"
                    hasChanged = true
                }
            }
        }
    }
    return hasChanged
}


func walkDirection(input [][]string, i int, j int, iDirection int, 
    jDirection int) int {
        retValue := 0
        for true {
            if iDirection == 1 { // We need to make sure i < len(input)
                if i < len(input) - 1 {
                    i = i + iDirection
                } else {
                    break
                }
            } else if iDirection == -1 { // We need to make sure i >= 0
                if i > 0 {
                    i = i + iDirection
                } else {
                    break
                }
            }

            if jDirection == 1 { // We need to make sure j < len(input[0])
                if j < len(input[0]) - 1 {
                    j = j + jDirection
                } else {
                    break
                }
            } else if jDirection == -1 { // We need to make sure j >= 0
                if j > 0 {
                    j = j + jDirection
                } else {
                    break
                }
            }
            fmt.Printf("i: %d, j: %d\n", i, j)
            time.Sleep(500 * time.Millisecond)
            if input[i][j] == "#" {
                retValue = 1
                break
            }
            if input[i][j] == "L" {
                break
            }
        }
        return retValue
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

func applyRulesS1(inputVals [][]string) bool {
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
