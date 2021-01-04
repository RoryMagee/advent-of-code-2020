package readfile

import(
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func readFile(path string) []int {
    fmt.Printf("Reading file: %s\n", path)
    file, _ := os.Open(path);
    defer file.Close();
    scanner := bufio.NewScanner(file)
    var returnArr []int
    for scanner.Scan() {
        line, _ := strconv.Atoi(scanner.Text());
        returnArr = append(returnArr, line)
    }
    return returnArr
}
