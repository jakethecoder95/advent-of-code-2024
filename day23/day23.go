package main

import (
	"advent2024/util"
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Computer struct {
    key string
    connections map[*Computer]struct{}
}

type Pair struct { a, b string }

func main() {
    computers, candidates := getComputers("input.txt")
    fmt.Println(part1(candidates))
    fmt.Println(part2(computers))
}

func part1(candidates map[*Computer]struct{}) int {
    total := 0
    visited := make(map[string]bool, 0)
    for candidate := range candidates {
        for conn1 := range candidate.connections {
            for conn2 := range conn1.connections {
                trio := GetKey(candidate, conn1, conn2)
                if visited[trio] {
                    continue
                }
                if _, ok := conn2.connections[candidate]; ok {
                    total++
                }
                visited[trio] = true
            }
        }
    }

    return total
}

func part2(computers map[string]*Computer) string {
    longestConnection := make([]*Computer, 0)
    visited := make(map[string]bool, 0)
    for _, computer := range computers {
        prev := []*Computer{computer}
        connection := getLongest(computer.connections, prev, visited)
        if len(connection) > len(longestConnection) {
            longestConnection = connection
        }
    }
    return GetKey(longestConnection...)
}

func getComputers(path string) (map[string]*Computer, map[*Computer]struct{}) {
    lines := util.ReadLinesAsSlice(path)

    computers := make(map[string]*Computer, 0)
    candidates := make(map[*Computer]struct{}, 0)

    for _, line := range lines {
        fields := strings.Split(line, "-")

        var c1, c2 *Computer
        var ok bool

        if c1, ok = computers[fields[0]]; !ok {
            c1 = &Computer{ key: fields[0], connections: make(map[*Computer]struct{}, 0) }
            computers[c1.key] = c1
        }
        if c2, ok = computers[fields[1]]; !ok {
            c2 = &Computer{ key: fields[1], connections: make(map[*Computer]struct{}, 0) }
            computers[c2.key] = c2
        }

        c1.connections[c2] = struct{}{}
        c2.connections[c1] = struct{}{}

        if c1.key[0] == 't' {
            candidates[c1] = struct{}{}
        }
        if c2.key[0] == 't' {
            candidates[c2] = struct{}{}
        }
    }

    return computers, candidates
}

func getLongest(connections map[*Computer]struct{}, prev []*Computer, visited map[string]bool) []*Computer {
    longestConnection := make([]*Computer, 0)
    for nextComputer := range connections {
        key := GetKey(append(prev, nextComputer)...)
        i := slices.IndexFunc(prev, func (c *Computer) bool {
            return c.key == nextComputer.key
        })
        if visited[key] || i >= 0 {
            continue
        }
        connection := getLongestConnection(nextComputer, prev, visited)
        visited[key] = true
        if len(connection) > len(longestConnection) {
            longestConnection = connection
        }
    }
    return longestConnection
}

func getLongestConnection(computer *Computer, prev []*Computer, visited map[string]bool) []*Computer {
    for _, prevComputer := range prev {
        if _, ok := prevComputer.connections[computer]; !ok || prevComputer.key == computer.key {
            return prev
        }
    }
    next := append(append([]*Computer{}, prev...), computer)
    return getLongest(computer.connections, next, visited)
}

type sortStrings []string

func (s sortStrings) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortStrings) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortStrings) Len() int {
    return len(s)
}

func GetKey(computers ...*Computer) string {
    var trio []string
    for _, computer := range computers {
        trio = append(trio, computer.key)
    }
    sort.Sort(sortStrings(trio))
    return strings.Join(trio, ",")
}
