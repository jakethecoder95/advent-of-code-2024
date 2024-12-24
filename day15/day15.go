package main

import (
	"advent2024/util"
	"fmt"
	"strings"
)

func main() {
    part1()
    part2()
}

func part1() {
    input := strings.Split(util.ReadFile("input.txt"), "\n\n")

    directions := strings.Join(strings.Split(input[1], "\n"), "")
    water := make([][]rune, 0)
    for _, line := range strings.Split(input[0], "\n") {
        water = append(water, []rune(line))
    }

    robotLoc := [2]int{-1, -1}
    for i, line := range water {
        for j, area := range line {
            if area == '@' {
                robotLoc = [2]int{i,j}
                break
            }
        }
        if robotLoc[0] != -1 {
            break
        }
    }

    for _, dir := range directions {
        if dir == '<' {
            moveX(water, &robotLoc, -1, 0)
        }
        if dir == '>' {
            moveX(water, &robotLoc, 1, len(water[0])-1)
        }
        if dir == '^' {
            moveY(water, &robotLoc, -1, 0)
        }
        if dir == 'v' {
            moveY(water, &robotLoc, 1, len(water)-1)
        }
    }

    total := 0
    for y, line := range water {
        for x, area := range line {
            if area == 'O' {
                total += 100 * y + x
            }
        }
    }

    fmt.Println("Part One:", total)
}

func part2() {
}

func moveX(water [][]rune, robotLoc *[2]int, dir, end int) {
    nextInd := robotLoc[1]+dir
    next := &water[robotLoc[0]][nextInd]
    if *next == 'O' {
        for i := nextInd+dir; i != end; i+=dir {
            nextNext := &water[robotLoc[0]][i]
            if *nextNext == '#' {
                break
            }
            if *nextNext == '.' {
                *nextNext = 'O'
                *next = '.'
                break
            }
        }
    }
    if *next == '.' {
        *next = '@'
        water[robotLoc[0]][robotLoc[1]] = '.'
        robotLoc[1] = nextInd
    }
}

func moveY(water [][]rune, robotLoc *[2]int, dir, end int) {
    nextInd := robotLoc[0]+dir
    next := &water[nextInd][robotLoc[1]]
    if *next == 'O' {
        for i := nextInd+dir; i != end; i+=dir {
            nextNext := &water[i][robotLoc[1]]
            if *nextNext == '#' {
                break
            }
            if *nextNext == '.' {
                *nextNext = 'O'
                *next = '.'
                break
            }
        }
    }
    if *next == '.' {
        *next = '@'
        water[robotLoc[0]][robotLoc[1]] = '.'
        robotLoc[0] = nextInd
    }
}

func show(water [][]rune, dir rune) {
    fmt.Println(string(dir))
    for _, line := range water {
        fmt.Println(string(line))
    }
}
