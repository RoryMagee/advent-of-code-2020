package main
import (
    "fmt"
    "strconv"
)

func solution1() {
    fmt.Println("attemp2")
    inputVals := readFile("./input")
    eastWestTotal := 0
    northSouthTotal := 0
    currentDirection := [2]int{0, 1}
    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        if dir == "N" {
            northSouthTotal = northSouthTotal + dis
        }
        if dir == "S" {
            northSouthTotal = northSouthTotal - dis
        }
        if dir == "E" {
            eastWestTotal = eastWestTotal + dis
        }
        if dir == "W" {
            eastWestTotal = eastWestTotal - dis
        }
        if dir == "L" || dir == "R" {
            // Here we need to update currentDirection value
            updateCurrentDirection(dir, dis, &currentDirection)
        }
        if dir == "F" {
            // Here we need to move dist in whatever direction currentDirection is
            i := currentDirection[0]
            j := currentDirection[1]
            if (i == 0 && j == 1) { // Moving East
                eastWestTotal = eastWestTotal + dis
            } else if i == 0 && j == -1 { // Moving West
                eastWestTotal = eastWestTotal - dis
            } else if i == -1 && j == 0 { // Moving North
                northSouthTotal = northSouthTotal + dis
            } else { // Moving South
                northSouthTotal = northSouthTotal - dis
            }
        }
    }
    fmt.Printf("%d+%d=%d\n", abs(eastWestTotal), abs(northSouthTotal),
    abs(eastWestTotal) + abs(northSouthTotal))
}

