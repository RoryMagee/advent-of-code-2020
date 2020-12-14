package main

import (
    "fmt"
    "strings"
    "strconv"
)

type bag struct {
    color string
    bagCount int
    counted bool
}

func solution2() {
    inputVals := readFile("./testinput2")
    bagMap := s2ParseInput(inputVals)
    count := 1
    fmt.Println("LEN", len(bagMap["shiny gold"]))
    count = walk(bagMap["shiny gold"], bagMap)
    fmt.Println("Count", count)
}

/*
Go to array of bags inside shiny gold
Go to the bottom "layer" of bags in each of those bags recursively
when you're at the bottom of a bag 
*/

func walk(b []*bag, bagMap map[string][]*bag) int {
    //if len(b) < 1 {
    //    return count
    //}
    count := 0
    for _, bag := range b {
        fmt.Println(bag.color)
        fmt.Println(count)
        count = count * walk(bagMap[bag.color], bagMap)
        fmt.Println(count)
    }
    return count
}

func s2ParseInput(input []string) map[string][]*bag  {
    m := map[string][]*bag{}
    for _, line := range input {
        splitLine := strings.Split(line, " ")
        startBag := strings.Join(splitLine[:2], " ")
        index := 4
        for index + 3 < len(splitLine) {
            nextBag := bag{}
            nextBag.bagCount, _ = strconv.Atoi(splitLine[index])
            nextBag.color = strings.Join(splitLine[index+1:index+3], " ")
            nextBag.counted = false
            m[startBag] = append(m[startBag], &nextBag)
            index = index + 4
        }
    }
    return m
}
