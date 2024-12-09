package util

import (
	"bufio"
	"fmt"
	"os"
)

type Solution interface {
	Solve(string)
}

func ReadLinesInFile(path string, cb func (line string)) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan() {
        cb(fileScanner.Text())
    }
}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(b)
}
