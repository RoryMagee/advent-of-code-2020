package main

import (
    "fmt"
    "strings"
    "strconv"
)

type bag struct {
    color string
    numberOf int
}

func solution2() {
    inputVals := readFile("./testinput")
    mp := s2ParseInput(inputVals)
    for _, m := range mp {
        fmt.Println(m[0].color)
        fmt.Println(m[0].numberOf)
    }
}

func s2ParseInput(input []string) map[string][]*bag  {
    m := map[string][]*bag{}
    for _, line := range input {
        splitLine := strings.Split(line, " ")
        startBag := strings.Join(splitLine[:2], " ")
        index := 4
        for index + 3 < len(splitLine) {
            nextBag := bag{}
            nextBag.numberOf, _ = strconv.Atoi(strings.Join(splitLine[index:index+1], " "))
            nextBag.color = strings.Join(splitLine[index+1:index+2], " ")
            m[startBag] = append(m[startBag], &nextBag)
            index = index + 3
        }
    }
    return m
}
