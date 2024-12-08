package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction int
const (
    Up = iota
    Down
    Left
    Right
)

var nextDirection = map[Direction]Direction{
    Up: Right,
    Right: Down,
    Down: Left,
    Left: Up,
}

func getKey(guard [2]int) string {
    return strconv.Itoa(guard[0]) + strconv.Itoa(guard[1])
}

func isObsUp(guard *[2]int, obs []bool, dir Direction, touched map[string]bool) (bool, Direction) {
    var start int
    var end int
    var increment bool
    var guardChangeIndex int

    if dir == Up {
        start = guard[0]-1
        end = 0
        increment = false
        guardChangeIndex = 0
    }
    if dir == Down {
        start = guard[0]+1
        end = len(obs)
        increment = true
        guardChangeIndex = 0
    }
    if dir == Left {
        start = guard[1]-1
        end = 0
        increment = false
        guardChangeIndex = 1
    }
    if dir == Right {
        start = guard[1]+1
        end = len(obs)
        increment = true
        guardChangeIndex = 1
    }

    if increment {
        for i := start; i < end; i++ {
            guard[guardChangeIndex] = i
            if obs[i] {
                if i != start {
                    guard[guardChangeIndex] = i-1
                }
                return true, nextDirection[dir]
            }
            touched[getKey(*guard)] = true
        }
    } else {
        for i := start; i >= end; i-- {
            guard[guardChangeIndex] = i
            if obs[i] {
                if i != start {
                    guard[guardChangeIndex] = i+1
                }
                return true, nextDirection[dir]
            }
            touched[getKey(*guard)] = true
        }
    }
    return false, dir
}

func day6Part1() int {

    path := os.Args[1]
    file, _ := os.Open(path)
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    // Map values
    verticals := map[int][]bool{}
    horizontals := map[int][]bool{}
    guard := [2]int{0, 0}

    graph := [][]rune{}
    index := 0
    for fileScanner.Scan() {
        graph = append(graph, []rune{})
        horizontals[index] = []bool{}
        for i, v := range fileScanner.Text() {
            if index == 0 {
                verticals[i] = []bool{}
            }
            horizontals[index] = append(horizontals[index], v == '#')
            verticals[i] = append(verticals[i], v == '#')
            if v == '^' {
                guard[0] = index
                guard[1] = i
            }
            graph[index] = append(graph[index], v)
        }
        index++
    }

    var dir Direction = Up
    isOnMap := true
    touched := map[string]bool{ getKey(guard): true }
    for isOnMap {
        var opt []bool
        if dir == Up || dir == Down {
            opt = verticals[guard[1]]
        } else {
            opt = horizontals[guard[0]]
        }
        isOnMap, dir = isObsUp(&guard, opt, dir, touched)
    }

    return len(touched)
}

func main() {
    fmt.Println("Day 6")
    fmt.Println("Part 1:", day6Part1())
}
