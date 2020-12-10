package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func newTree(val int, color string) *tree {
    root := newNode(val, color)
    tree := new(tree)
    tree.root = root
    return tree
}

func main() {
    inputVals := readFile("./input")
    for _, line := range inputVals {
        fmt.Println(line)
    }
    parseInput(inputVals)
}

func parseInput(input []string)  {
    for idx, line := range input {
        fmt.Println(idx, line)
        splitLine := strings.Split(line, "contain")
        fmt.Println(splitLine[1])
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
