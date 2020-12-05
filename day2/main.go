package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    fmt.Println("Solution 1")
    inputVals := readFile("./input")
    //validPasses := 0
    for i := 0; i < len(inputVals); i++ {
        pass := inputVals[i]
        r, c, p := splitPass(pass)
        fmt.Println(r, c, p)
    }
}

func splitPass(passString string) (rule string, char string, password string) {
    // 7-19 l: llvllllllclllflllll
    //fmt.Println(passString)
    split := strings.Split(passString, " ")
    r := split[0]
    c := split[1][:1]
    p := split[2]
    return r, c, p
}

func validPass(rule string, char string, password string) bool {
    return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(path string) []string {
    file, err := os.Open(path)
    check(err)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val := scanner.Text()
        lines = append(lines, val)
    }
    return lines
}
