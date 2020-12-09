package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
    color string
    val int
    children []*node
}

type tree struct {
    root *node
}

func newNode(val int, color string) *node {
    node := new(node)
    node.color = color
    node.val = val
    return node
}

func newTree(root *node) *tree {
}

func main() {
    fmt.Println("Day 7")
    inputVals := readFile("./input")
    for _, line := range inputVals {
        fmt.Println(line)
    }
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()
    var returnArr []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        returnArr = append(returnArr, line)
    }
    return returnArr
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
