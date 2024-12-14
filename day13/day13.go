package day13

import (
	"advent2024/util"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() int {
    input := util.ReadFile(os.Args[1])
    parts := strings.Split(input, "\n\n")

	tokens := 0
	for _, p := range parts {
		desc := strings.Split(p, "\n")
		A := ints(desc[0])
		B := ints(desc[1])
		prize := ints(desc[2])
		n, m, ok := solve(A, B, prize)

		if ok {
			tokens += 3*n + m
		}
	}

	return tokens
}

func Part2() int {
    input := util.ReadFile(os.Args[1])
    parts := strings.Split(input, "\n\n")

	tokens := 0
	for _, p := range parts {
		desc := strings.Split(p, "\n")
		A := ints(desc[0])
		B := ints(desc[1])
		prize := ints(desc[2])
		prize[0] += 10000000000000
		prize[1] += 10000000000000
		n, m, ok := solve(A, B, prize)
		if ok {
			tokens += 3*n + m
		}
	}

	return tokens
}

func solve(A, B, prize []int) (int, int, bool) {
	a0, a1, b0, b1, pr0, pr1 := float64(A[0]), float64(A[1]), float64(B[0]), float64(B[1]), float64(prize[0]), float64(prize[1])

	if n, m, ok := linalgSolve(a0, a1, b0, b1, pr0, pr1); ok && isIntegral(n) && isIntegral(m) {
		return int(n), int(m), true
	}

	return -1, -1, false
}

func linalgSolve(a0, a1, b0, b1, c0, c1 float64) (float64, float64, bool) {
	det := a0*b1 - a1*b0

	if det == 0.0 {
		return 0.0, 0.0, false
	}

	return (c0*b1 - c1*b0) / det, (a0*c1 - a1*c0) / det, true
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}


func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		r = append(r, atoi(e))
	}
	return
}
