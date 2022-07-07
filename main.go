package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Grid [][]string

const (
	maxLength = 30
	life      = "x"
	nolife    = " "
)

func main() {
	var gridLength int
	for {
		var inputLength string
		fmt.Println("Please input grid length.")
		fmt.Print(">")
		fmt.Scan(&inputLength)
		i, err := strconv.Atoi(inputLength)
		if err != nil {
			fmt.Println("The length must be number.")
			continue
		}
		if i < 1 || i > maxLength {
			fmt.Println("The length must be greater than or equal to 1 and less than or equal to 30.")
			continue
		}
		gridLength = i
		break
	}

	var repeatNumber int
	for {
		var inputRepeat string
		fmt.Println("Please input repeat number.")
		fmt.Print(">")
		fmt.Scan(&inputRepeat)
		i, err := strconv.Atoi(inputRepeat)
		if err != nil {
			fmt.Println("The length must be number.")
			continue
		}
		if i < 1 {
			fmt.Println("The length must be greater than or equal to 1.")
			continue
		}
		repeatNumber = i
		break
	}

	grid := make(Grid, gridLength)
	initializeGrid(grid, gridLength)
	printGrid(grid)
	for repeat := 0; repeat < repeatNumber; repeat++ {
		nextGrid := deepCopySlice(grid, gridLength)
		next(grid, nextGrid, gridLength)
		grid = nextGrid
		fmt.Printf("Generation %d\n", repeat+1)
		printGrid(grid)
		time.Sleep(100 * time.Millisecond)
	}
}

func initializeGrid(grid Grid, gridLength int) {
	for i := 0; i < gridLength; i++ {
		grid[i] = make([]string, gridLength)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < gridLength; i++ {
		for j := 0; j < gridLength; j++ {
			switch rand.Intn(2) {
			case 0:
				grid[i][j] = life
			default:
				grid[i][j] = nolife
			}
		}
	}
}

func next(currentGrid Grid, nextGrid Grid, gridLength int) {
	for y := 0; y < gridLength; y++ {
		for x := 0; x < gridLength; x++ {
			if isNextAlive(currentGrid, y, x, gridLength) {
				nextGrid[y][x] = life
			} else {
				nextGrid[y][x] = nolife
			}
		}
	}
}

func isNextAlive(currentGrid Grid, y, x, gridLength int) bool {
	aroundCount := 0
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -1; offsetX <= 1; offsetX++ {
			if offsetY == 0 && offsetX == 0 {
				continue
			}
			currentY := y + offsetY
			currentX := x + offsetX
			if currentY < 0 || currentY >= gridLength {
				continue
			}
			if currentX < 0 || currentX >= gridLength {
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

func deepCopySlice(originalGrid Grid, gridLength int) Grid {
	grid := make(Grid, gridLength)
	for i := 0; i < gridLength; i++ {
		dst := make([]string, gridLength)
		grid[i] = dst
		copy(grid[i], originalGrid[i])
	}
	return grid
}
