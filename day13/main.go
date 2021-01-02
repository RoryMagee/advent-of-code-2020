package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Print("Day 13\n")
    solution1()
}

func readFile(path string) []string {
    file, _ := os.Open(path);
    defer file.Close();
    scanner := bufio.NewScanner(file);
    retArr := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        retArr = append(retArr, line)
    }
    return retArr
}
