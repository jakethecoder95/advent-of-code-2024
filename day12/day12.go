package day12

import (
	"advent2024/util"
	"bufio"
	"os"
)

type Region struct {
    plot      string
	area      int
	perimeter int
    sides     int
}

type Cell struct {
	i, j int
}

func getValidPerimeter(h int, w int, garden [][]rune) [][3]int {
    surroundingIndexes := make([][3]int, 0)
    // Up
    if h > 0 {
        up := [3]int{h-1, w, 1}
        surroundingIndexes = append(surroundingIndexes, up)
    } else {
        surroundingIndexes = append(surroundingIndexes, [3]int{-1, -1, 1})
    }
    // Down
    if h < len(garden)-1 {
        down := [3]int{h+1, w, 2}
        surroundingIndexes = append(surroundingIndexes, down)
    } else {
        surroundingIndexes = append(surroundingIndexes, [3]int{-1, -1, 2})
    }
    // Left
    if w > 0 {
        down := [3]int{h, w-1, 3}
        surroundingIndexes = append(surroundingIndexes, down)
    } else {
        surroundingIndexes = append(surroundingIndexes, [3]int{-1, -1, 3})
    }
    // Right
    if w < len(garden[h])-1 {
        down := [3]int{h, w+1, 4}
        surroundingIndexes = append(surroundingIndexes, down)
    } else {
        surroundingIndexes = append(surroundingIndexes, [3]int{-1, -1, 4})
    }
    return surroundingIndexes
}

func calcPerimeter(h int, w int, garden [][]rune) int {
    plot := garden[h][w]
    perimeter := 0
    for _, ixs := range getValidPerimeter(h, w, garden) {
        if ixs[0] == -1 {
            perimeter++
            continue
        }
        adjacentPlot := garden[ixs[0]][ixs[1]]
        if adjacentPlot != plot {
            perimeter++
        }
    }
    return perimeter
}

func recurseRegion(
    h int,
    w int,
    region *Region,
    garden [][]rune,
    seen map[int]map[int]bool,
) {
	plot := garden[h][w]

    seen[h][w] = true
    region.area++
    region.perimeter += calcPerimeter(h, w, garden)

    for _, ixs := range getValidPerimeter(h, w, garden) {
        if ixs[0] == -1 {
            continue
        }
        adjacentPlot := garden[ixs[0]][ixs[1]]
        if adjacentPlot == plot && !seen[ixs[0]][ixs[1]] {
            recurseRegion(ixs[0], ixs[1], region, garden, seen)
        }
    }
}

func Part1() int {
	garden := make([][]rune, 0)
    seen := make(map[int]map[int]bool, 0)
	util.ReadLinesInFile(os.Args[1], func(line string, i int) {
        seen[i] = make(map[int]bool, 0)
		garden = append(garden, []rune(line))
	})

    regions := make([]Region, 0)

    for h := range garden {
        for w, plot := range garden[h] {
            if seen[h][w] {
                continue
            }
            region := Region{
                area: 0,
                perimeter: 0,
                plot: string(plot),
            }
            recurseRegion(h, w, &region, garden, seen)
            regions = append(regions, region)
        }
    }

    total := 0
    for _, region := range regions {
        total += region.area * region.perimeter
    }

    return total
}

func Part2() int {
	input, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var garden [][]byte
	for scanner.Scan() {
		garden = append(garden, append([]byte{}, scanner.Bytes()...))
	}

	var area int
	var plant byte
	var dfs func(int, int)
	n, m := len(garden), len(garden[0])
	visited := map[Cell]bool{}

	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || garden[i][j] != plant {
			return
		}
		area++
		garden[i][j] = '#'
		visited[Cell{i, j}] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	var res int
	for i := range n {
		for j := range m {
			if garden[i][j] == '#' {
				continue
			}
			area = 0
			plant = garden[i][j]
			dfs(i, j)
			plot, counted := plot(visited, n, m)
			sides := sides(plot, counted)
			res += area * sides
			clear(visited)
		}
	}

	return res
}

func plot(visited map[Cell]bool, n, m int) ([][]bool, [][][]bool) {
	n += 2
	m += 2
	grid := make([][]bool, n)
	counted := make([][][]bool, n)
	for i := range n {
		grid[i] = make([]bool, m)
		counted[i] = make([][]bool, m)
		for j := range m {
			counted[i][j] = make([]bool, 4)
		}
	}
	for cell := range visited {
		grid[cell.i+1][cell.j+1] = true
	}
	return grid, counted
}

func sides(grid [][]bool, counted [][][]bool) int {
	n, m := len(grid), len(grid[0])
	var res int

	for i := range n {
		for j := range m {
			if grid[i][j] {
				if !grid[i-1][j] {
					counted[i][j][0] = true
					if !counted[i][j-1][0] {
						res++
					}
				}
				if !grid[i+1][j] {
					counted[i][j][1] = true
					if !counted[i][j-1][1] {
						res++
					}
				}
				if !grid[i][j-1] {
					counted[i][j][2] = true
					if !counted[i-1][j][2] {
						res++
					}
				}
				if !grid[i][j+1] {
					counted[i][j][3] = true
					if !counted[i-1][j][3] {
						res++
					}
				}
			}
		}
	}

	return res
}
