package main

import (
    "fmt"
)

func inputIsValid(fieldName string, fieldVal string) bool {
    switch fieldName {
    case "byr":
        fmt.Println("byr")
    case "iyr":
        fmt.Println("iyr")
    case "eyr":
        fmt.Println("eyr")
    case "hgt":
        fmt.Println("hgt")
    }
    return true
}
