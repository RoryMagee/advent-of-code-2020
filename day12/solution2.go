package main

import (
	"fmt"
	"strconv"
)

func solution2() {
    fmt.Println("Solution2")
    inputVals := readFile("./testinput")
    shipPosition := [2]int{0, 0}
    waypointPosition := [2]int{1, 10}
    currentDirection := [2]int{0, 1}
    for i := 0; i < len(inputVals); i++ {
        dir := string(inputVals[i][0])
        dis, _ := strconv.Atoi(inputVals[i][1:])
        switch dir {
        case "N":

        }
        case "E":

        }
        case "S":

        }
        case "W":

        }
    }
}
