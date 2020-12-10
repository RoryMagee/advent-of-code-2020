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
    inputVals := readFile("./input")
    parsedInput := parseInput(inputVals)
    for idx, input := range parsedInput {
        fmt.Println(idx, input)
    }
    //nodeArr := make(map[string] *node)
    //for _, line := range inputVals {
        //splitLine := strings.Split(line, " ")
        //color := strings.Join(splitLine[:2], " ")
        //_ := getNode(nodeArr, color)
        /* Here we want to loop through the rest of the items in the string
           and add them to currenNode's children array
        */
        //restOfString := strings.Split(strings.Join(splitLine[4:], " "), ", ")
        //for _, remaining := range restOfString {
        //    bag := 
        //    currentNode.children = append(currentNode.children, bag)
        //}
    //}
}

func parseInput(input []string) []string { // [startBag, bag2, bag3]
    parsedInput := []string{}
    for _, line := range input {
        splitLine := strings.Split(line, " ")
        startBag := strings.Join(splitLine[:2], " ")
        parsedInput = append(parsedInput, startBag)
        index := 5
        for index + 4 < len(splitLine) {
            nextBag := strings.Join(splitLine[index:index+2], " ")
            index = index + 4
            parsedInput = append(parsedInput, nextBag)
        }
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
