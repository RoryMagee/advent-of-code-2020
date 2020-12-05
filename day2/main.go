package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    fmt.Println("Solution 1")
    inputVals := readFile("./input")
    sl1CorrectPasswords := 0
    sl2CorrectPasswords := 0
    for i := 0; i < len(inputVals); i++ {
        pass := inputVals[i]
        r, c, p := splitPass(pass)
        answerSln1 := solution1(r, p, c)
        if (answerSln1) {
            sl1CorrectPasswords++
        }
        answerSln2 := solution2(r, p, c)
        if (answerSln2) {
            sl2CorrectPasswords++
        }
    }
    fmt.Println("Solution1:", sl1CorrectPasswords)
    fmt.Println("Solution2:", sl2CorrectPasswords)
}

func countOccurances(passString string, char string) int {
    charArr := strings.Split(passString, "")
    occurances := 0
    for _, c := range(charArr) {
        if char == string(c) {
            occurances++
        }
    }
    return occurances
}

func splitPass(passString string) (rule [2]int, char string, password string) {
    // 1-12 m: wmjwgmmlmmmm
    split := strings.Split(passString, " ")
    ruleSplit := strings.Split(split[0], "-")
    lowerCount, _ := strconv.Atoi(ruleSplit[0])
    upperCount, _ := strconv.Atoi(ruleSplit[1])
    r := [2]int{lowerCount, upperCount} // {1, 2}
    c := split[1][:1] // m
    p := split[2] // wmkwgmmlmmm
    return r, c, p
}

func solution1(rule [2]int, passString string, char string) bool {
    occurances := countOccurances(passString, char)
    if occurances >= rule[0] && occurances <= rule[1] {
        return true
    } else {
        return false
    }
}

func checkPositions(lowerChar string, upperChar string, char string) bool {
    if (lowerChar == char && upperChar != char) || (lowerChar != char && upperChar == char) {
        return true
    } else {
        return false
    }
}

func solution2(rule [2]int, passString string, char string) bool {
    fmt.Println(passString, char, rule[0], rule[1])
    lowerCharIdx := rule[0]
    upperCharIdx := rule[1]
    lowerChar := passString[lowerCharIdx-1: lowerCharIdx]
    upperChar := passString[upperCharIdx-1: upperCharIdx]
    return checkPositions(lowerChar, upperChar, char)
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
