package main

import (
    "fmt"
    "strconv"
)

func solution1() {
    fmt.Println("Solution 1")
    inputVals := readFile("./testinput")
    currDir := 90
    currLocation := [2]int{0, 0}
    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        if dir == "L" || dir == "R" {
            updateDirection(currDir, dir, dis)
        } else {
            if dir == "N" {
                currLocation[0] = currLocation[0] - dis
            } else if dir == "E" {
                currLocation[1] = currLocation[1] + dis
            } else if dir == "S" {
                currLocation[0] = currLocation[0] + dis
            } else if dir == "W" {
                currLocation[1] = currLocation[1] - dis
            } else if dir == "F" {
                dirArr := angleToCo(currDir)
                currLocation[dirArr[0]] = currLocation[dirArr[0]] + dis
                currLocation[dirArr[1]] = currLocation[dirArr[1]] + dis
            }
        }
    }
    fmt.Println(currLocation)
}

func angleToCo(angle int) [2]int{
    co := [2]int{}
    if angle == 0 {
        co[0] = -1
        co[1] = 0
    } else if angle == 90 {
        co[0] = 0
        co[1] = 1
    } else if angle == 180 {
        co[0] = 1
        co[1] = 0
    } else if angle == 270 {
        co[0] = 0
        co[1] = -1
    }
    return co
}

func coToAngle(co[2]int) int {
    if co[0] == -1 && co[1] == 0 {
        return 0
    } else if co[0] == 0 && co[1] == -1 {
        return 270
    } else if co[0] == 0 && co[1] == 1 {
        return 90
    } else {
        return 180
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
