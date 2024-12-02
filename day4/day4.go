package main

import (
	"bufio"
	"fmt"
	"os"
)

func Log(a ...any) {
    if len(os.Args) > 2 && os.Args[2] == "--log" {
        fmt.Println(a...)
    }
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func checkUp(graph [][]rune, row int, col int) int {
    if row-3 < 0 {
        return 0
    }

    word := ""
    for i := 0; i < 4; i++ {
        y := row - i
        word += string(graph[y][col])
    }

    // Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkDown(graph [][]rune, row int, col int) int {
    if row+3 >= len(graph) {
        return 0
    }

    word := ""
    for i := 0; i < 4; i++ {
        y := row + i
        word += string(graph[y][col])
    }

    // Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkLeft(graph [][]rune, row int, col int) int {
    if col-3 < 0 {
        return 0
    }

    word := string(graph[row][col-3:col+1])
    // Log(word)
    if Reverse(word) == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkRight(graph [][]rune, row int, col int) int {
    if col+3 >= len(graph[row]) {
        return 0
    }

    word := string(graph[row][col:col+4])
    // Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkDiagonalUpRight(graph [][]rune, row int, col int) int {
    // Check out of range
    if row-3 < 0 || col+3 >= len(graph[row]) {
        return 0
    }

    word := ""
    for i := 0; i < 4; i++ {
        x := col + i
        y := row - i
        word += string(graph[y][x])
    }

    // Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkDiagonalUpLeft(graph [][]rune, row int, col int) int {
    // Check out of range
    if row-3 < 0 || col-3 < 0 {
        return 0
    }

    word := ""
    for i := 0; i < 4; i++ {
        x := col - i
        y := row - i
        word += string(graph[y][x])
    }

    Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkDiagonalDownRight(graph [][]rune, row int, col int) int {
    // Check out of range
    if row+3 >= len(graph) || col+3 >= len(graph[row]) {
        return 0
    }

    word := ""
    for i := 0; i < 4; i++ {
        x := col + i
        y := row + i
        word += string(graph[y][x])
    }

    // Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}

func checkDiagonalDownLeft(graph [][]rune, row int, col int) int {
    // Check out of range
    if row+3 >= len(graph) || col-3 < 0 {
        return 0
    }

    word := ""
    for i := 0; i < 4; i++ {
        x := col - i
        y := row + i
        word += string(graph[y][x])
    }

    // Log(word)
    if word == "XMAS" {
        return 1
    } else {
        return 0
    }
}


func Day4Part1() int {
    location := os.Args[1]
    file, _ := os.Open(location)
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    total := 0
    graph := [][]rune{}

    for fileScanner.Scan() {
        graph = append(graph, []rune(fileScanner.Text()))
    }

    for row, line := range graph {
        for col, c := range line {
            if c == 'X' {
                total += checkUp(graph, row, col)
                total += checkDown(graph, row, col)
                total += checkLeft(graph, row, col)
                total += checkRight(graph, row, col)
                total += checkDiagonalUpLeft(graph, row, col)
                total += checkDiagonalUpRight(graph, row, col)
                total += checkDiagonalDownLeft(graph, row, col)
                total += checkDiagonalDownRight(graph, row, col)
            }
        }
    }

    return total
}

func Day4Part2() int {
    location := os.Args[1]
    file, _ := os.Open(location)
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    total := 0
    graph := [][]rune{}

    for fileScanner.Scan() {
        graph = append(graph, []rune(fileScanner.Text()))
    }

    for row, line := range graph {
        for col, c := range line {
            if c == 'A' {
                if col-1 < 0 || col+1 >= len(line) || row-1 < 0 || row+1 >= len(graph) {
                    continue
                }

                diagonalDownRight := string([]rune{graph[row-1][col-1], c, graph[row+1][col+1]})
                diagonalDownLeft  := string([]rune{graph[row-1][col+1], c, graph[row+1][col-1]})

                isDDRMas := diagonalDownRight == "MAS" || diagonalDownRight == "SAM"
                isDDLMas := diagonalDownLeft == "MAS" || diagonalDownLeft == "SAM"

                if isDDRMas && isDDLMas {
                    total++
                }
            }
        }
    }

    return total
}

func main() {
    fmt.Println("Day 4")
    fmt.Println("Part 1", Day4Part1())
    fmt.Println("Part 2", Day4Part2())
}
