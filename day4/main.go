package main

import (
	"bufio"
	"os"
    "fmt"
    "strings"
)

func main() {
    inputVals := readFile("./input")
    validPassports := 0
    for _, p := range inputVals {
        fields := getPassportFields(p)
        if checkFields(fields) {
            validPassports++
        }
    }
    fmt.Println(validPassports)
}

func checkFields(fields []string) bool {
    required := [7]string {
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid",
    }
    retValue := true
    for _, field := range required {
        if !isValueInArr(fields, field) {
            retValue = false
            break
        }
    }
    return retValue
}

func isValueInArr(arr []string, key string) bool {
    for _, val := range arr {
        if val == key {
            return true
        }
    }
    return false
}

func getPassportFields(passportDetails string) []string {
    parsedDetails := strings.Fields(passportDetails)
    var fields []string
    for _, field := range parsedDetails {
        fields = append(fields, field[0:3])
    }
    return fields
}

func readFile(path string) []string {
    file, e := os.Open(path)
    check(e)
    defer file.Close()

    var passports []string
    scanner := bufio.NewScanner(file)
    var line string
    for scanner.Scan() {
        newLine := scanner.Text()
        if newLine == "" {
            passports = append(passports, line)
            line = ""
        } else {
            line = line + " " + newLine
        }
    }
    return passports
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
