package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    inputVals := readFile("./input")
    solution1 := 0
    solution2 := 0
    for _, v := range inputVals {
        unique := findUnique(v)
        solution1 = solution1 + len(unique)
        solution2 = solution2 + unanimous(unique, v)
    }
    fmt.Println("solution1:", solution1)
    fmt.Println("solution2:", solution2)
}

func unanimous(unique map[string] bool, group []string) int {
    noInGroup := len(unique)
    count := 0
    groupAsStr := ""
    for _, answer := range group {
        groupAsStr = groupAsStr + answer
    }
    for key := range unique {
        if countOccurances(groupAsStr, key) == noInGroup {
            count = count + 1
        }
    }
    fmt.Println(unique, group, count)
    return count
}

func countOccurances(str string, key string) int {
    count := 0
    for _, character := range str {
        if string(character) == key {
            count = count + 1
        }
    }
    return count
}

func findUnique(group []string) map[string] bool {
    m := make(map[string] bool)
    str := ""
    for _, arr := range group {
        str = str + arr
    }
    for _, g := range str {
        m[string(g)] = true
    }
    return m
}

func readFile(path string) [][]string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var str []string
    var groups [][]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            groups = append(groups, str)
            str = nil
        } else {
            str = append(str, line)
        }
    }
    groups = append(groups, str)
    return groups
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
