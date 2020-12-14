package main

import (
    "fmt"
    "strings"
    "strconv"
)

type bag struct {
    color string
    numberOf int
    counted bool
}

func solution2() {
    inputVals := readFile("./input")
    bagMap := s2ParseInput(inputVals)
    fmt.Println(walk(bagMap["shiny gold"], 0, bagMap))
}

func walk(bagList []*bag, count int, bagMap map[string][]*bag) int {
    c := count
    if len(bagList) < 1 {
        return c
    }
    for _, bag := range bagList {
        if bag.counted == false {
            c = c + bag.numberOf
            walk(bagMap[bag.color], c, bagMap)
        }
    }
    return c
}

func s2ParseInput(input []string) map[string][]*bag  {
    m := map[string][]*bag{}
    for _, line := range input {
        splitLine := strings.Split(line, " ")
        startBag := strings.Join(splitLine[:2], " ")
        index := 4
        for index + 3 < len(splitLine) {
            nextBag := bag{}
            nextBag.numberOf, _ = strconv.Atoi(splitLine[index])
            nextBag.color = strings.Join(splitLine[index+1:index+3], " ")
            nextBag.counted = false
            m[startBag] = append(m[startBag], &nextBag)
            index = index + 4
        }
    }
    return m
}
