package main

import (
    "fmt"
)

func main() {
    res := weirdBinarySearchTing(127, "FBFBBFF", "F", "B")
    fmt.Println(res)
}

func weirdBinarySearchTing(max int, seq string, upperChar string, lowerChar string) int {
    min := 0
    for _, c := range seq { // Loop through seq
        /*
         * On each iteration of the loop if c points to the lower half then
         * we keep min the same and assign max as the halfway point between min
         * and max. Otherwise we keep max the same and assign min as the halfway
         * point between current min and max
         */
        val := string(c)
        if val == upperChar {
            max = (max + min) / 2
        } else {
            min = (max + min) / 2
        }
    }
    return max
}
