package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    inputVals := readFile("./input")
    fmt.Println(inputVals)
    total := 0
    for _, v := range inputVals {
        total = total + findUnique(v)
    }
    fmt.Println(total)
}

func findUnique(group string) int {
    m := make(map[string] bool)
    for _, g := range group {
        m[string(g)] = true
    }
    return len(m)
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var str string
    var groups []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            groups = append(groups, str)
            str = ""
        } else {
            str = str + line
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
