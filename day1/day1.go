package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1Part1() {
    filePath := os.Args[1]
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    var numGroup1 []int
    var numGroup2 []int

    for fileScanner.Scan() {
        words := strings.Fields(fileScanner.Text())

        num1, _ := strconv.Atoi(words[0])
        num2, _ := strconv.Atoi(words[1])

        numGroup1 = append(numGroup1, num1)
        numGroup2 = append(numGroup2, num2)
    }

    sort.Ints(numGroup1)
    sort.Ints(numGroup2)

    result := 0
    for i := 0; i < len(numGroup1); i++ {
        num1 := numGroup1[i]
        num2 := numGroup2[i]

        if num1 == num2 {
            continue
        }

        if num1 > num2 {
            result += num1 - num2
        } else {
            result += num2 - num1
        }
    }

    fmt.Println(result)
}

func day1Part2() {
    filePath := os.Args[1]
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    var numGroup1 []int
    var numGroup2 = map[int]int{}

    for fileScanner.Scan() {
        words := strings.Fields(fileScanner.Text())

        num1, _ := strconv.Atoi(words[0])
        num2, _ := strconv.Atoi(words[1])

        numGroup1 = append(numGroup1, num1)

        val, ok := numGroup2[num2]
        if ok {
            numGroup2[num2] = val + 1
        } else {
            numGroup2[num2] = 1
        }
    }

    result := 0

    for _, v := range numGroup1 {
        occurrances, ok := numGroup2[v]
        if ok {
            result += occurrances * v
        }
    }

    fmt.Println(result)
}

func main() {
    fmt.Println("Day1")

    fmt.Println("Part 1: ")
    day1Part1()

    fmt.Println("Part 2: ")
    day1Part2()
}
