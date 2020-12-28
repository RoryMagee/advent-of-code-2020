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
    //currDir := 90
    currDir := [2]int{0,1}
    currLocation := [2]int{0,0}
    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        if dir == "L" || dir == "R" {
            updateDirection(currDir, dir, dis)
        } else {
            // dir can be N / E / S / W / F
            if dir == "N" {
                currLocation[0] = currLocation[0] - dis
            } else if dir == "E" {
                currLocation[1] = currLocation[1] + dis
            } else if dir == "S" {
                currLocation[0] = currLocation[0] + dis
            } else if dir == "W" {
                currLocation[1] = currLocation[1] - dis
            } else if dir == "F" {
                currLocation
            }
        }
    }
}

func updateDirection2(currDirection [2]int, dir string, dis int) {
    noOfTurns := dis / 90

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
