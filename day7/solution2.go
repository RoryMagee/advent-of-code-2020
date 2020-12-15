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
    inputVals := readFile("./input")
    bagMap := s2ParseInput(inputVals)
    //count := walk(bagMap["shiny gold"][0], bagMap, -1)
    count := 0
    for _, bg := range bagMap["shiny gold"] {
        count = count + walk(bg, bagMap, -1)
    }
    fmt.Println("Count", count)
}

/*
Go to array of bags inside shiny gold
Go to the bottom "layer" of bags in each of those bags recursively
when you're at the bottom of a bag 
*/

func walk(b *bag, bagMap map[string][]*bag, num int) int {
    if num == -1 {
        num = 1
    }
    if len(bagMap[b.color]) < 1 {
        return num * b.bagCount
    }
    num = num * b.bagCount
    count := num
    for _, nextBag := range bagMap[b.color] {
        t := walk(nextBag, bagMap, num)
        count = count + t
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
