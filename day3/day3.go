package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3Part1() {
    filePath := os.Args[1]
    file, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println(err)
    }

    var total = 0

    reMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

    reNum := regexp.MustCompile(`\d{1,3}`)
    for _, v := range reMul.FindAll(file, -1) {
        numbers := reNum.FindAll(v, -1)
        num1, _ := strconv.Atoi(string(numbers[0]))
        num2, _ := strconv.Atoi(string(numbers[1]))
        total += num1 * num2
    }

    fmt.Println(total)
}

func Day3Part2() {

    filePath := os.Args[1]
    file, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println(err)
    }

    var total = 0

    reMul := regexp.MustCompile(`(?:mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\))`)
    matches := reMul.FindAll(file, -1)

    reNum := regexp.MustCompile(`\d{1,3}`)
    mulEnabled := true

    for _, v := range matches {
        match := string(v)
        if match == "do()" {
            mulEnabled = true
        } else if match == "don't()" {
            mulEnabled = false
        } else if mulEnabled && match != "do()" && match != "don't()" {
            numbers := reNum.FindAll(v, -1)
            num1, _ := strconv.Atoi(string(numbers[0]))
            num2, _ := strconv.Atoi(string(numbers[1]))
            total += num1 * num2
        }
    }

    fmt.Println(total)
}

func main() {
    fmt.Println("Day 3")

    fmt.Println("Part 1: ")
    Day3Part1()

    fmt.Println("Part 2: ")
    Day3Part2()
}
