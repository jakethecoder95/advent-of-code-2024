package day11

import (
	"advent2024/util"
	"math"
	"os"
	"strconv"
	"strings"
)

type NumAndBlinks struct {
	Num    int
	Blinks int
}

func split_stone(num int, blinks int, memo map[NumAndBlinks]int) int {
	if blinks == 0 {
		return 1
	}

	prev := memo[NumAndBlinks{Num: num, Blinks: blinks}]
	if prev != 0 {
		return prev
	}

	var total int
	if num == 0 {
		total = split_stone(1, blinks-1, memo)
	} else if is_even, half1, half2 := is_even_length(num); is_even {
		total = split_stone(half1, blinks-1, memo) + split_stone(half2, blinks-1, memo)
	} else {
		total = split_stone(num*2024, blinks-1, memo)
	}

	memo[NumAndBlinks{Num: num, Blinks: blinks}] = total
	return total
}

func is_even_length(num int) (bool, int, int) {
	pow_of_ten := 1
	for num%int(math.Pow10(pow_of_ten)) != num {
		pow_of_ten++
	}
	if pow_of_ten%2 != 0 {
		return false, 0, 0
	}

	half1 := num % int(math.Pow10(pow_of_ten/2))
	half2 := num / int(math.Pow10(pow_of_ten/2))
	return true, half1, half2
}

func solve(nums []int, blinks int) int {
	var prev_found = make(map[NumAndBlinks]int)

	var total = 0
	for _, num := range nums {
		total += split_stone(num, blinks, prev_found)
	}
    return total
}

func Part1() int {

	var text = util.ReadFile(os.Args[1])
	var num_strs = strings.Fields(text)
	var nums = make([]int, len(num_strs))

	for i, num_str := range num_strs {
		num, err := strconv.Atoi(num_str)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}

    return solve(nums, 25)
}

func Part2() int {

	var text = util.ReadFile(os.Args[1])
	var num_strs = strings.Fields(text)
	var nums = make([]int, len(num_strs))

	for i, num_str := range num_strs {
		num, err := strconv.Atoi(num_str)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}

    return solve(nums, 75)
}
