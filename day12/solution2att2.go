package main
import (
    "fmt"
    "strconv"
)

func solution1att2() {
    fmt.Println("attemp2")
    inputVals := readFile("./testinput")
    eastWestTotal := 0
    northSouthTotal := 0
    currentDirection := [2]int{0, 1}

    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        if dir == "N" || dir == "S" {
            northSouthTotal = northSouthTotal + dis
        }
        if dir == "E" || dir == "W" {
            eastWestTotal = eastWestTotal + dis
        }
        if dir == "L" || dir == "R" {
            // Here we need to update currentDirection value
            fmt.Println("Before", currentDirection)
            updateCurrentDirection(dir, dis, currentDirection)
            fmt.Println("After", currentDirection)
        }
        if dir == "F" {
            // Here we need to move dist in whatever direction currentDirection is
            fmt.Println("FFFF")
        }
    }
}

func updateCurrentDirection(dir string, dis int, currentDirection [2]int) {
    var noOfTurns int
    //noOfTurns := dis / 90
    if dir == "R" {
        noOfTurns = dis / 90
    } else {
        noOfTurns = (360 - dis) / 90
    }
    fmt.Println("NoOfTurns", noOfTurns)
    for i := 0; i < noOfTurns; i++ {
        increment(&currentDirection)
    }
    fmt.Println("currentDirection", currentDirection)
}

func abs(num int) int {
    if num < 0 {
        num = num - num * 2
    }
    return num
}

func increment(i* [2]int) {
    if i[0] == 0 {
        if i[1] == 1 {
            // return 1, 0
            i[0] = 1
            i[1] = 0
        } else { // [0, -1]
            // return -1, 0
            i[0] = -1
            i[1] = 0
        }
    } else { // [1, ?]
        if i[1] == 0 { // [1, 0]
            // return 0, -1
            i[0] = 0
            i[1] = -1
        } else {
            // return 0, 1
            i[0] = 0
            i[1] = 1
        }
    }
}

/* Clockwise rotation
0, 1
1, 0
0, -1
-1, 0
*/
