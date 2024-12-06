package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Page struct {
    after  map[int]bool
}

func day5Part1() int {
    path := os.Args[1]
    file, _ := os.Open(path)

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    pages := map[int]Page{}

    // Build the pages
    for fileScanner.Scan() {
        line := fileScanner.Text()
        if strings.TrimSpace(line) == "" {
            break
        }
        order := strings.Split(line, "|")
        first, _ := strconv.Atoi(order[0])
        second, _ := strconv.Atoi(order[1])
        if page, ok := pages[second]; ok {
            page.after[first] = true
        } else {
            pages[second] = Page{
                after: map[int]bool{},
            }
            pages[second].after[first] = true
        }
    }

    total := 0

    for fileScanner.Scan() {
        line := fileScanner.Text()
        words := strings.Split(line, ",")
        isLineValid := true
        for i, w := range words {
            page, _ := strconv.Atoi(w)
            for _, w2 := range words[:i] {
                prevPage, _ := strconv.Atoi(w2)
                if !pages[page].after[prevPage] {
                    isLineValid = false
                    break
                }
            }
            if !isLineValid {
                break
            }
        }
        if isLineValid {
            middleIndex := len(words) / 2
            middle, _ := strconv.Atoi(words[middleIndex])
            total += middle
        }
    }

    return total
}

func day5Part2() int {
    path := os.Args[1]
    file, _ := os.Open(path)

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    pages := map[int]Page{}

    // Build the pages
    for fileScanner.Scan() {
        line := fileScanner.Text()
        if strings.TrimSpace(line) == "" {
            break
        }
        order := strings.Split(line, "|")
        first, _ := strconv.Atoi(order[0])
        second, _ := strconv.Atoi(order[1])
        if page, ok := pages[second]; ok {
            page.after[first] = true
        } else {
            pages[second] = Page{
                after: map[int]bool{},
            }
            pages[second].after[first] = true
        }
    }

    total := 0

    for fileScanner.Scan() {
        line := fileScanner.Text()
        words := strings.Split(line, ",")
        isLineValid := true
        for i, w := range words {
            page, _ := strconv.Atoi(w)
            for i2, w2 := range words[:i] {
                prevPage, _ := strconv.Atoi(w2)
                if !pages[page].after[prevPage] {
                    isLineValid = false
                    hold := words[i]
                    words[i] = words[i2]
                    words[i2] = hold
                }
            }
        }
        if !isLineValid {
            middleIndex := len(words) / 2
            middle, _ := strconv.Atoi(words[middleIndex])
            total += middle
        }
    }

    return total
}

func main() {
    fmt.Println("Day 5")
    fmt.Println("Part 1", day5Part1())
    fmt.Println("Part 2", day5Part2())
}
