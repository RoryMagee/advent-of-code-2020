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
    inputVals := readFile("./testinput2")
    bagMap := s2ParseInput(inputVals)
    count := 0
    for _, bag := range bagMap["shiny gold"] {
        count = count + walk(bag, count, bagMap)
    }
    fmt.Println("COunt", count)
}

func walk(b *bag, count int, bagMap map[string][]*bag) int {
    c := count
    if b.counted == false {
        c = c + b.numberOf
        for _, newBag := range bagMap[b.color] {
            walk(newBag, c, bagMap)
        }
        b.counted = false
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
