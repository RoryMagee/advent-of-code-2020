package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "strconv"
)

type instruction struct {
    operation string
    amount int
    visited bool
}

func main() {
    fmt.Println("day 8")
    inputVals := readFile("./input")
    parsedInput := parseInput(inputVals)
    solution1Terminates, solution1Acc := checkIfProgramTerminates(parsedInput)
    fmt.Println("Solution 1", solution1Terminates, solution1Acc)
    reset(parsedInput)
    for i := 0; i < len(parsedInput); i++ {
        if parsedInput[i].operation == "jmp" {
            parsedInput[i].operation = "nop"
        } else if  parsedInput[i].operation == "nop" {
            parsedInput[i].operation = "jmp"
        }
        t, a := checkIfProgramTerminates(parsedInput)
        reset(parsedInput)
        if t == true {
            fmt.Println("solution2", a)
            break
        } else {
            if parsedInput[i].operation == "jmp" {
                parsedInput[i].operation = "nop"
            } else if  parsedInput[i].operation == "nop" {
                parsedInput[i].operation = "jmp"
            }
        }
    }
}

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func reset(instructions []*instruction) {
    for _, i := range instructions {
        i.visited = false
    }
}

func printOperations(instructions []*instruction) {
    for _, i := range instructions {
        fmt.Printf("%s, ", i.operation)
    }
    fmt.Printf("\n")
}

func checkIfProgramTerminates(gameInput []*instruction) (bool, int) {
    acc := 0
    nextInstructionIndex := 0
    for true {
        if nextInstructionIndex > len(gameInput)-1 {
            break
        }
        i := gameInput[nextInstructionIndex]
        if i.visited == true {
            return false, acc
        }
        i.visited = true
        if i.operation == "nop" {
            nextInstructionIndex++
            continue
        }
        if i.operation == "jmp" {
            nextInstructionIndex = nextInstructionIndex + i.amount
            continue
        }
        if i.operation == "acc" {
            nextInstructionIndex++
            acc = acc + i.amount
            continue
        }
    }
    return true, acc
}

func parseInput(input []string) []*instruction {
    var returnArr []*instruction
    for _, line := range input {
        splitLine := strings.Split(line, " ")
        operation := splitLine[0]
        amount, _ := strconv.Atoi(string(splitLine[1][1:]))
        if string(splitLine[1][0]) == "-" {
            amount = -amount
        }
        i := instruction{operation: operation, amount: amount, visited: false}
        returnArr = append(returnArr, &i)
    }
    return returnArr
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var returnArr []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        returnArr = append(returnArr, scanner.Text())
    }
    return returnArr
}
