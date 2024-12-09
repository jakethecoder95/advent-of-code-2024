package day8

import (
	"advent2024/util"
	// "fmt"
	"os"
	"strconv"
	"unicode"
)

func coordinatesAreValid(y int, x int, graph [][]rune, antinodes map[string]bool) {
    if y >= len(graph) || y < 0 {
        return
    }
    if x >= len(graph[y]) || x < 0 {
        return
    }
    joined := strconv.Itoa(y) + strconv.Itoa(x)
    antinodes[joined] = true
}

func Part1() int {
    path := os.Args[1]

    antenasCoor := map[rune][][2]int{}
    m := [][]rune{}
    i := 0
    util.ReadLinesInFile(path, func(line string) {
        m = append(m, []rune(line))
        for j, c := range line {
        	if c == '.' {
                continue
            }
            if _, ok := antenasCoor[c]; !ok {
                antenasCoor[c] = [][2]int{}
            }
            antenasCoor[c] = append(antenasCoor[c], [2]int{i, j})
        }
        i++
    })

    antinodes := map[string]bool{}

    for _, coordinates := range antenasCoor {
        for j, lowerCoor := range coordinates {
            for _, upperCoor := range coordinates[j+1:] {
                y := upperCoor[0] - lowerCoor[0]
                x := upperCoor[1] - lowerCoor[1]

                lowerAntinodeY := lowerCoor[0]-y
                lowerAntinodeX := lowerCoor[1]-x
                upperAntinodeY := upperCoor[0]+y
                upperAntinodeX := upperCoor[1]+x

                coordinatesAreValid(lowerAntinodeY, lowerAntinodeX, m, antinodes)
                coordinatesAreValid(upperAntinodeY, upperAntinodeX, m, antinodes)
            }
        }
    }

    return len(antinodes)
}

func Part2() int {
    path := os.Args[1]

    antenasCoor := map[rune][][2]int{}
    m := [][]rune{}
    i := 0
    util.ReadLinesInFile(path, func(line string) {
        m = append(m, []rune(line))
        for j, c := range line {
        	if !unicode.IsLetter(c) && unicode.IsDigit(c) {
                return
            }
            if _, ok := antenasCoor[c]; !ok {
                antenasCoor[c] = [][2]int{}
            }
            antenasCoor[c] = append(antenasCoor[c], [2]int{i, j})
        }
    })

    total := 0

    return total
}
