package util

import (
	"bufio"
	"fmt"
	"os"
)

type Solution interface {
	Solve(string)
}

func ReadLinesAsSlice(path string) []string {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    slice := make([]string, 0)
    for fileScanner.Scan() {
        slice = append(slice, fileScanner.Text())
    }

    return slice
}

func ReadLinesInFile(path string, cb func (line string, index int)) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    index := 0
    for fileScanner.Scan() {
        cb(fileScanner.Text(), index)
        index++
    }
}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(b)
}
