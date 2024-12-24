package day14

import (
	"advent2024/util"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// const HEIGHT = 7
// const WIDTH = 11
const HEIGHT = 103
const WIDTH = 101

type Coordinate struct{ x, y int }

func Part1() int {

    path := os.Args[1]
    graph := make(map[Coordinate]int, 0)
    seconds := 100

    for _, line := range util.ReadLinesAsSlice(path) {
        fields := strings.Fields(line)
        position := strings.Split(string(fields[0][2:]), ",")
        velocity := strings.Split(string(fields[1][2:]), ",")
        px, py := atoi(position[0]), atoi(position[1])
        vx, vy := atoi(velocity[0]), atoi(velocity[1])
        endPosition := getFinalPosition(px, py, vx, vy, seconds)

        if _, ok := graph[endPosition]; ok {
            graph[endPosition]++
        } else {
            graph[endPosition] = 1
        }
    }

    q1, q2, q3, q4 := 0, 0, 0, 0
    for coor, count := range graph{
        if coor.x < WIDTH/2 && coor.y < HEIGHT/2 {
            q1 += count
        } else if coor.x > WIDTH/2 && coor.y < HEIGHT/2 {
            q2 += count
        } else if coor.x < WIDTH/2 && coor.y > HEIGHT/2 {
            q3 += count
        } else if coor.x > WIDTH/2 && coor.y > HEIGHT/2 {
            q4 += count
        }
    }

    return q1*q2*q3*q4
}

func getFinalPosition(posx, posy, velx, vely, sec int) Coordinate {
    px, py, vx, vy, s := float64(posx), float64(posy), float64(velx), float64(vely), float64(sec)
    x := math.Mod(px+(s*vx), WIDTH)
    y := math.Mod(py+(s*vy), HEIGHT)
    if x < 0 {
        x = WIDTH+x
    }
    if y < 0 {
        y = HEIGHT+y
    }
    return Coordinate{
        x: int(x),
        y: int(y),
    }
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
