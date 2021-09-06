package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var (
	maxRows = flag.Int("max_rows", 500, "maximum rows in world")
	maxCols = flag.Int("max_cols", 500, "maximum columns in world")
	displayWorld = flag.Bool("display_world", false, "display world on console")
)

func rowLine() {
	fmt.Print("\n")
	for i := 0; i < *maxCols; i++ {
		fmt.Print(" -----")
	}
	fmt.Print("\n")
}

func countLiveNeighbourCell(arr [][]int, curr_row int, curr_col int) int {
	count := 0

	for i := curr_row-1; i <= curr_row+1; i++ {
		for j := curr_col-1; j <= curr_col+1; j++ {
			if (i==curr_row && j==curr_col) || (i<0 || j<0) || (i>=*maxRows || j>=*maxCols) {
				continue
			}
			if arr[i][j]==1 {
                count += 1
            }
		}
	}

	return count
}

func nextWorldState(world01 *[][]int, world02 *[][]int) {
	neighbour_live_cell := 0
	
	for i := 0; i < *maxRows; i++ {
		for j := 0; j < *maxCols; j++ {
			neighbour_live_cell = countLiveNeighbourCell(*world01, i, j)

			if (*world01)[i][j]==1 && (neighbour_live_cell==2 || neighbour_live_cell==3) {
                (*world02)[i][j] = 1
            } else if (*world01)[i][j]==0 && neighbour_live_cell==3 {
                (*world02)[i][j] = 1
            } else {
                (*world02)[i][j] = 0
            }
		}
	}
}

func printWorld(world *[][]int) {
	rowLine()
    for i := 0; i < *maxRows; i++ {
        fmt.Print(":")
        for j := 0; j < *maxCols; j++ {
            fmt.Printf("  %d  :", (*world)[i][j])
        }
        rowLine()
    }
}


func main() {
	flag.Parse()

	world01 := make([][]int, *maxRows)
	for i := 0; i < *maxRows; i++ {
		world01[i] = make([]int, *maxCols)
	}

	for i := 0; i < *maxRows; i++ {
		for j := 0; j < *maxCols; j++ {
			world01[i][j] = rand.Intn(32767) % 2
		}
	}

	world02 := make([][]int, *maxRows)
	for i := 0; i < *maxRows; i++ {
		world02[i] = make([]int, *maxCols)
	}

	for i := 0; i < *maxRows; i++ {
		for j := 0; j < *maxCols; j++ {
			world02[i][j] = 0
		}
	}

	if *displayWorld {
		printWorld(&world01)
	}
	

	nextWorldState(&world01, &world02)

	if *displayWorld {
		fmt.Println("Next Generation:")
		printWorld(&world02)
	}
}