package main

import (
    "os"
    "bufio"
)

func readFile(path string) []string {
    file, _ := os.Open(path)
    defer file.Close()
    var retVal []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        retVal = append(retVal, line)
    }
    return retVal
}

func abs(num int) int {
    if num < 0 {
        num = num - num * 2
    }
    return num
}

/* Clockwise rotation
east: 0, 1
sound: 1, 0
west: 0, -1
north: -1, 0
*/

func rotate(waypoint* [2]int, dir string, dis int) {
    var noOfTurns int
    //noOfTurns := dis / 90
    if dir == "R" {
        noOfTurns = dis / 90
    } else {
        noOfTurns = (360 - dis) / 90
    }
    for x := 0; x < noOfTurns; x++ {
        i := waypoint[0]
        j := waypoint[1]
        waypoint[0] = -j
        waypoint[1] = i
    }
    /*
    i       j 
    ----------
    4       10
    -10     4
    -4      -10
    10      -4
    */
}

func increment(i* [2]int) {
    iIdx := i[0]
    jIdx := i[1]
    if iIdx == 0 && jIdx == 1 { // East
        i[0] = 1
        i[1] = 0
    } else if iIdx == 1 && jIdx == 0 {
        i[0] = 0
        i[1] = -1
    } else if iIdx == 0 && jIdx == -1 {
        i[0] = -1
        i[1] = 0
    } else if iIdx == -1 && jIdx == 0 {
        i[0] = 0
        i[1] = 1
    }
}

func updateCurrentDirection(dir string, dis int, currentDirection* [2]int) {
    var noOfTurns int
    //noOfTurns := dis / 90
    if dir == "R" {
        noOfTurns = dis / 90
    } else {
        noOfTurns = (360 - dis) / 90
    }
    for i := 0; i < noOfTurns; i++ {
        increment(currentDirection)
    }
}
