package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsValid(prevNumber int, currentNumber int, isDecreasing bool) bool {
    if Abs(prevNumber - currentNumber) > 3 || prevNumber - currentNumber == 0 {
        return false
    } else if !isDecreasing && prevNumber < currentNumber {
        return false
    } else if isDecreasing && prevNumber > currentNumber {
        return false
    } else {
        return true
    }
}

func day2Part1() {
    filePath := os.Args[1]
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    validReports := 0
    for fileScanner.Scan() {

        words := strings.Fields(fileScanner.Text())

        isDecreasing := false
        isReportValid := true
        prevNumber := -1

        for i, word := range words {
            number, _ := strconv.Atoi(word)

            if i == 0 {
                prevNumber = number
                continue
            }

            if i == 1 && prevNumber < number {
                isDecreasing = true
            }

            if Abs(prevNumber - number) > 3 || prevNumber - number == 0 {
                isReportValid = false
                break
            } else if !isDecreasing && prevNumber < number {
                isReportValid = false
                break
            } else if isDecreasing && prevNumber > number {
                isReportValid = false
                break
            }

            prevNumber = number
        }

        if isReportValid {
            validReports = validReports + 1
        }
    }

    fmt.Println(validReports)
}

func day1Part2() {

    filePath := os.Args[1]
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    validReports := 0
    for fileScanner.Scan() {

        words := strings.Fields(fileScanner.Text())

        isDecreasing := false
        isReportValid := true
        prevNumber := -1
        removeCount := 0

        for i, word := range words {

            number, _ := strconv.Atoi(word)
            hasNextNumber := i < len(words) - 1

            var nextNumber int
            if hasNextNumber {
                value, _ := strconv.Atoi(words[i+1])
                nextNumber = value
            }

            if !hasNextNumber && removeCount == 0 {
                break
            }

            if i == 0 {
                prevNumber = number
                continue
            }

            if i == 1 && prevNumber < number {
                isDecreasing = true
            }

            if !IsValid(prevNumber, number, isDecreasing) {
                if hasNextNumber &&
                   removeCount == 0 &&
                   IsValid(number, nextNumber, isDecreasing) {
                    removeCount = removeCount + 1
                    continue
                }
                isReportValid = false
                break
            }

            prevNumber = number
        }

        if isReportValid {
            validReports = validReports + 1
        }
    }

    fmt.Println(validReports)
}

func main() {
    fmt.Println("Day2")

    fmt.Println("Part 1: ")
    day2Part1()

    fmt.Println("Part 2: ")
    day1Part2()
}
