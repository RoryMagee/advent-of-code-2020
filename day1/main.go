package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    inputVals := readFile()
    solution1(inputVals)
    solution2(inputVals)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile() []int {
    file, err := os.Open("./input")
    check(err)
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val, err := strconv.Atoi(scanner.Text())
        check(err)
        lines = append(lines, val)
    }
    return lines
}

func solution1(inputVals []int) {
    m := make(map[int] bool)
    for i := 0; i < len(inputVals); i++ {
        if  m[2020 - inputVals[i]] {
            fmt.Println("solution1:", (2020 - inputVals[i]) * inputVals[i])
            break
        }
        m[inputVals[i]] = true
    }

}

func solution2(inputVals []int) {
    m := make(map[int]map[int]int)
    for i := 0; i < len(inputVals); i++ {
        for j := i + 1; j < len(inputVals) ; j++ {
            a := inputVals[i]
            b := inputVals[j]
            m[a + b] = map[int]int{a: b}
        }
    }
    for k := 0; k < len(inputVals); k++ {
        remainingVal := 2020 - inputVals[k]
        if m[remainingVal] != nil {
            var key int
            for l := range m[remainingVal] {
                key = l
            }
            fmt.Println("solution2:", inputVals[k] * key * m[remainingVal][key])
            break
        }
    }
}
