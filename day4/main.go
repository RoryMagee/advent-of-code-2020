package main

import (
	"bufio"
	"os"
    "fmt"
    "strings"
)

func test() {

}

func main() {
    inputVals := readFile("./input")
    solution1 := 0
    solution2 := 0
    for _, p := range inputVals {
        fields := getPassportFields(p)
        if checkFieldsExist(fields) {
            solution1++
        }
        if checkFieldsExist(fields) && validateFields(fields) {
            solution2++
        }
    }
    fmt.Println("Solution1:", solution1)
    fmt.Println("Solution2:", solution2)
}

func validateFields(fields [][]string) bool {
    for _, field := range fields {
        if !inputIsValid(field[0], field[1]) {
            return false
        }
    }
    return true
}

func checkFieldsExist(fields [][]string) bool {
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

func isValueInArr(arr [][]string, key string) bool {
    for _, val := range arr {
        if val[0] == key {
            return true
        }
    }
    return false
}

func getPassportFields(passportDetails string) [][]string {
    parsedDetails := strings.Fields(passportDetails)
    var fields [][]string
    for _, field := range parsedDetails {
        splitFields := strings.Split(field, ":")
        fields = append(fields, splitFields) // [hgt, byr, icd, ...]
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
