package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Playground struct {
	x_dim int
	y_dim int
}

type Robot struct {
	x_vel int
	y_vel int
	x_pos int
	y_pos int
}

func ReadFile(fileName string) []string {
	start := time.Now()
	file, _ := os.Open(fileName)
	defer file.Close()
	var returnArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		returnArray = append(returnArray, text)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("File Runtime: %v\n", duration.Seconds())
	return returnArray
}

func main() {
	start := time.Now()
	loc := ReadFile("day14.txt")
	play := Playground{x_dim: 101, y_dim: 103}
	res1 := Task1(loc, play)
	res2 := Task2(loc, play)
	fmt.Println("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Runtime: %v\n", duration.Microseconds())
}

func Task1(lines []string, playground Playground) int {
	var robots []Robot
	for _, line := range lines {
		robots = append(robots, createBot(line))
	}
	for i := range robots {
		moveTimes(&robots[i], &playground, 100)
	}
	return calcSafety(&robots, playground)
}

func Task2(lines []string, playground Playground) int {
	var robots []Robot
	for _, line := range lines {
		robots = append(robots, createBot(line))
	}
	foundSolution := findLines(&robots, playground)
	seconds := 0
	for !foundSolution {
		for i := range robots {
			move(&robots[i], &playground)
		}
		seconds++
		foundSolution = findLines(&robots, playground)
	}
	printField(&robots, playground)
	return seconds
}

func createBot(input string) Robot {
	input = strings.ReplaceAll(input, "=", "")
	input = strings.ReplaceAll(input, "p", "")
	input = strings.ReplaceAll(input, "v", "")
	inputSplit := strings.Split(input, " ")
	posSplit := strings.Split(inputSplit[0], ",")
	posx, _ := strconv.Atoi(posSplit[0])
	posy, _ := strconv.Atoi(posSplit[1])
	velSplit := strings.Split(inputSplit[1], ",")
	velx, _ := strconv.Atoi(velSplit[0])
	vely, _ := strconv.Atoi(velSplit[1])
	return Robot{x_pos: posx, y_pos: posy, x_vel: velx, y_vel: vely}
}

func moveTimes(robot *Robot, field *Playground, times int) {
	for i := 0; i < times; i++ {
		move(robot, field)
	}
}

func move(robot *Robot, field *Playground) {
	newX := (robot.x_pos + robot.x_vel + field.x_dim) % field.x_dim
	newY := (robot.y_pos + robot.y_vel + field.y_dim) % field.y_dim
	robot.x_pos = newX
	robot.y_pos = newY
}

func calcSafety(robots *[]Robot, field Playground) int {
	mid_x := field.x_dim / 2
	mid_y := field.y_dim / 2
	sum1 := 0
	sum2 := 0
	sum3 := 0
	sum4 := 0
	for _, robot := range *robots {
		if robot.x_pos < mid_x && robot.y_pos < mid_y {
			sum1++
		}
		if robot.x_pos < mid_x && robot.y_pos > mid_y {
			sum2++
		}
		if robot.x_pos > mid_x && robot.y_pos < mid_y {
			sum3++
		}
		if robot.x_pos > mid_x && robot.y_pos > mid_y {
			sum4++
		}
	}
	return sum1 * sum2 * sum3 * sum4
}

func printField(robots *[]Robot, field Playground) {
	playArray := make([][]rune, field.y_dim)
	for i := range playArray {
		playArray[i] = make([]rune, field.x_dim)
		for j := range playArray[i] {
			playArray[i][j] = '.'
		}
	}
	for _, robot := range *robots {
		playArray[robot.y_pos][robot.x_pos] = 'X'
	}
	for i := range playArray {
		fmt.Println(string(playArray[i]))
	}
}

func findLines(robots *[]Robot, field Playground) bool {
	playArr := make([][]bool, field.y_dim)
	for i := range playArr {
		playArr[i] = make([]bool, field.x_dim)
	}
	for _, robot := range *robots {
		playArr[robot.y_pos][robot.x_pos] = true
	}
	maxLen := 0
	for i := range field.y_dim {
		currLen := 0
		for j := range field.x_dim {
			if !playArr[i][j] {
				if currLen > maxLen {
					maxLen = currLen
				}
				currLen = 0
			} else {
				currLen++
			}
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen > 30
}
