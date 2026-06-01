package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	printRoom(Beginner())
}

func printRoom(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				fmt.Print("███")
			} else {
				fmt.Print(" □ ")
			}
		}
		fmt.Println()
	}
}

func Beginner() [][]int {
	return [][]int{
		{0, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 1, 1},
		{1, 1, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 0},
	}
}
