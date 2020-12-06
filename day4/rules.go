package main

import (
    "fmt"
    "strconv"
    "unicode"
)

func inputIsValid(fieldName string, fieldVal string) bool {
    retValue := false
    switch fieldName {
    case "byr":
        retValue = byr(fieldVal)
    case "iyr":
        retValue = iyr(fieldVal)
    case "eyr":
        retValue = eyr(fieldVal)
    case "hgt":
        retValue = hgt(fieldVal)
    case "hcl":
        retValue = hcl(fieldVal)
    case "ecl":
        retValue = ecl(fieldVal)
    case "pid":
        retValue = pid(fieldVal)
    case "cid":
        retValue = true
    default:
        fmt.Println("default", fieldName)
    }
    return retValue
}

func pid(val string) bool {
    if len(val) != 9 {
        return false
    }
    for _, char := range val {
        if !unicode.IsNumber(char) {
            return false
        }
    }
    return true
}

func ecl(val string) bool {
    allowedVals := []string{
        "amb",
        "blu",
        "brn",
        "gry",
        "grn",
        "hzl",
        "oth",
    }
    for _, color := range allowedVals {
        if val == color {
            return true
        }
    }
    return false
}

// THIS MIGHT NOT WORKKKKK
func hcl(val string) bool {
    if len(val) != 7 {
        return false
    }
    if val[0:1] != "#" {
        return false
    }
    color := string(val[1:])
    for _, char := range color {
        if !unicode.IsNumber(char) &&
        char < 97 || char > 102 {
            return false
        }
    }
    return true
}

func hgt(val string) bool {
    unit := val[len(val)-2:]
    if unit == "cm" {
        measurement, err := strconv.Atoi(val[:len(val)-2])
        check(err)
        if (measurement >= 150 && measurement <= 193) {
            return true
        }
    } else if unit == "in" {
        measurement, err := strconv.Atoi(val[:len(val)-2])
        check(err)
        if (measurement >= 59 && measurement <= 76) {
            return true
        }
    }
    return false
}

func byr(val string) bool {
    valAsInt, err := strconv.Atoi(val)
    check(err)
    if valAsInt <= 2002 && valAsInt >= 1920 {
        return true
    }
    return false
}

func iyr(val string) bool {
    valAsInt, err := strconv.Atoi(val)
    check(err)
    if valAsInt <= 2020 && valAsInt >= 2010 {
        return true
    }
    return false
}

func eyr(val string) bool {
    valAsInt, err := strconv.Atoi(val)
    check(err)
    if valAsInt <= 2030 && valAsInt >= 2020 {
        return true
    }
    return false
}
