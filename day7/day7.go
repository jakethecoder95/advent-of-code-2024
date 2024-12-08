package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Equation struct{
    testValue int
    numbers []int
}

func assessNumber(num int, carry []int, testValue int, concat bool) ([]int, bool) {
    nextCarry := []int{}
    containsTestValue := false
    for _, c := range carry {
        mult := num*c
        sum := num+c
        joined, _ := strconv.Atoi(strconv.Itoa(c) + strconv.Itoa(num))
        if mult <= testValue {
            nextCarry = append(nextCarry, mult)
        }
        if sum <= testValue {
            nextCarry = append(nextCarry, sum)
        }
        if concat && joined <= testValue {
            nextCarry = append(nextCarry, joined)
        }
        if sum == testValue || mult == testValue || (concat && joined == testValue) {
            containsTestValue = true
        }
    }
    return nextCarry, containsTestValue
}

func getEquation(line string) Equation {
    fields := strings.Fields(line)
    testValue, _ := strconv.Atoi(string(fields[0][:len(fields[0])-1]))
    numbers := []int{}
    for _, v := range fields[1:] {
        number, _ := strconv.Atoi(v)
        numbers = append(numbers, number)
    }
    return Equation{
        testValue: testValue,
        numbers: numbers,
    }
}

func Part1() int64 {

    path := os.Args[1]
    file, _ := os.Open(path)
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    validEquations := []int{}

    for fileScanner.Scan() {
        line := fileScanner.Text()
        equation := getEquation(line)
        carry := []int{equation.numbers[0]}
        for i, num := range equation.numbers[1:] {
            nextCary, containsTarget := assessNumber(num, carry, equation.testValue, false)
            carry = nextCary
            if containsTarget && i == len(equation.numbers)-2 {
                validEquations = append(validEquations, equation.testValue)
            }
        }
    }

    var result int64 = 0
    for _, v := range validEquations {
        result += int64(v)
    }

    return result
}

func Part2() int64 {

    path := os.Args[1]
    file, _ := os.Open(path)
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    validEquations := []int{}

    for fileScanner.Scan() {
        line := fileScanner.Text()
        equation := getEquation(line)
        carry := []int{equation.numbers[0]}
        for i, num := range equation.numbers[1:] {
            nextCary, containsTarget := assessNumber(num, carry, equation.testValue, true)
            carry = nextCary
            if containsTarget && i == len(equation.numbers)-2 {
                validEquations = append(validEquations, equation.testValue)
            }
        }
    }

    var result int64 = 0
    for _, v := range validEquations {
        result += int64(v)
    }

    return result
}
