package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    inputVals := readFile("./input")
    solution1 := 0
    for _, v := range inputVals {
        fmt.Println(v)
        solution1 = solution1 + findUnique(v)
    }
    fmt.Println("solution1:", solution1)
}


func findUnique(group []string) int {
    m := make(map[string] bool)
    str := ""
    for _, arr := range group {
        str = str + arr
    }
    for _, g := range str {
        m[string(g)] = true
    }
    return len(m)
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
