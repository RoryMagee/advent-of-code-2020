package main

import (
    "fmt"
    "strconv"
)

const(
    north = 0
    east = 90
    south = 180
    west = 270
)

func solution1() {
    fmt.Println("Solution 1")
    fmt.Println(updateDirection(east, "R", 180))
    inputVals := readFile("./testinput")
    currDir := 90
    currLocation := [2]int{0,0}
    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        if dir == "L" || dir == "R" {
            updateDirection(currDir, dir, dis)
        } else {

        }
    }
}

func updateDirection(currentDirection int, dir string, dis int) int {
    if dir == "L" {
        currentDirection = currentDirection - dis
    } else {
        currentDirection = currentDirection + dis
    }
    if currentDirection >= 360 {
        currentDirection = currentDirection - 360
    } else if currentDirection < 0 {
        currentDirection = currentDirection + 360
    }
    return currentDirection
}
