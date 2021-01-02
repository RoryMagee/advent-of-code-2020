package main

import (
	"fmt"
	"strconv"
)

func solution2() {
    fmt.Println("Solution2")
    inputVals := readFile("./input")
    shipPosition := [2]int{0, 0}
    waypointPosition := [2]int{1, 10}
    //currentDirection := [2]int{0, 1}

    /*
    Formatting for positional arrays
    [North/South][East/West]
    */

    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        switch dir {
        case "N":
            waypointPosition[0] = waypointPosition[0] + dis
        case "E":
            waypointPosition[1] = waypointPosition[1] + dis
        case "S":
            waypointPosition[0] = waypointPosition[0] - dis
        case "W":
            waypointPosition[1] = waypointPosition[1] - dis
        case "R", "L":
            rotate(&waypointPosition, dir, dis)
        case "F":
            nsMovement := dis * waypointPosition[0]
            ewMovement := dis * waypointPosition[1]
            shipPosition[0] = shipPosition[0] + nsMovement
            shipPosition[1] = shipPosition[1] + ewMovement
        }
    }
    answer := abs(shipPosition[0]) + abs(shipPosition[1])
    fmt.Printf("%d, %d, %d\n", shipPosition[0], shipPosition[1], answer)
}
