package main

import (
    "os"
    "bufio"
    "strconv"
)

func readFile(path string) []map[string]int {
    file, _ := os.Open(path)
    defer file.Close()
    retVal := []map[string]int{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        m := make(map[string]int)
        line := scanner.Text()
        direction := string(line[0])
        num, _ := strconv.Atoi(line[1:])
        m[direction] = num
        retVal = append(retVal, m)
    }
    return retVal
}
