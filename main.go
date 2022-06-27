package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid [initialRows][initialCols]string

const (
	initialRows = 30
	initialCols = 30
	life        = "x"
	nolife      = " "
)

func main() {
	grid := Grid{}
	initializeGrid(&grid)
	printGrid(grid)
	for repeat := 0; repeat < 100; repeat++ {
		nextGrid := Grid{}
		next(grid, &nextGrid)
		grid = nextGrid
		fmt.Printf("Generation %d\n", repeat+1)
		printGrid(grid)
		time.Sleep(100 * time.Millisecond)
	}
}

func initializeGrid(grid *Grid) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < initialRows; i++ {
		for j := 0; j < initialCols; j++ {
			switch rand.Intn(2) {
			case 0:
				grid[i][j] = life
			default:
				grid[i][j] = nolife
			}
		}
	}
}

func next(currentGrid Grid, nextGrid *Grid) {
	for y := 0; y < initialRows; y++ {
		for x := 0; x < initialCols; x++ {
			if isNextAlive(currentGrid, y, x) {
				nextGrid[y][x] = life
			} else {
				nextGrid[y][x] = nolife
			}
		}
	}
}

func isNextAlive(currentGrid Grid, y, x int) bool {
	aroundCount := 0
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -1; offsetX <= 1; offsetX++ {
			if offsetY == 0 && offsetX == 0 {
				continue
			}
			currentY := y + offsetY
			currentX := x + offsetX
			if currentY < 0 || currentY >= initialRows {
				continue
			}
			if currentX < 0 || currentX >= initialCols {
				continue
			}
			if currentGrid[currentY][currentX] == life {
				aroundCount++
			}
		}
	}

	if currentGrid[y][x] == nolife && aroundCount == 3 {
		return true
	}

	if currentGrid[y][x] == life {
		if aroundCount == 2 || aroundCount == 3 {
			return true
		}
		if aroundCount <= 1 {
			return false
		}
		if aroundCount >= 4 {
			return false
		}
	}

	return false
}

func printGrid(grid Grid) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("")
}
