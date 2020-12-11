package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
    counted bool
    color string
    children []*node
}

func main() {
    inputVals := readFile("./testinput")
    parsedInput := parseInput(inputVals)
    nodeArr := make(map[string] *node)
    for _, line := range parsedInput {
        currentNode := getNode(nodeArr, line[0])
        /* Here we want to loop through the rest of the items in the string
           and add them to currenNode's children array
        */
        for i := 1; i < len(line); i++ {
            bag := getNode(nodeArr, line[i])
            currentNode.children = append(currentNode.children, bag)
        }
    }
    fmt.Println(nodeArr["light red"].children)
}



func parseInput(input []string) [][]string {
    parsedInput := [][]string{}
    for idx, line := range input {
        parsedLine := []string{}
        splitLine := strings.Split(line, " ")
        startBag := strings.Join(splitLine[:2], " ")
        parsedLine = append(parsedLine, startBag)
        index := 5
        fmt.Println(index, len(splitLine))
        for index + 4 < len(splitLine) {
            nextBag := strings.Join(splitLine[index:index+2], " ")
            fmt.Println(idx, "next bag >", nextBag)
            index = index + 4
            parsedLine = append(parsedLine, nextBag)
        }
        parsedInput = append(parsedInput, parsedLine)
    }
    return parsedInput
}

func getNode(nodeArr map[string] *node, nodeName string) *node{
    if nodeArr[nodeName] == nil {
        newNode := new(node)
        newNode.color = nodeName
        newNode.counted = false
        newNode.children = []*node{}
        nodeArr[nodeName] = newNode
    }
    return nodeArr[nodeName]
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
