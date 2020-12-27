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
