package main

import (
	"fmt"
	"math/rand"
)

const maxRows = 500
const maxCols = 500

func rowLine() {
	fmt.Print("\n")
	for i := 0; i < maxCols; i++ {
		fmt.Print(" -----")
	}
	fmt.Print("\n")
}

func countLiveNeighbourCell(arr [maxRows][maxCols]int, curr_row int, curr_col int) int {
	count := 0

	for i := curr_row-1; i <= curr_row+1; i++ {
		for j := curr_col-1; j <= curr_col+1; j++ {
			if (i==curr_row && j==curr_col) || (i<0 || j<0) || (i>=maxRows || j>=maxCols) {
				continue
			}
			if arr[i][j]==1 {
                count += 1
            }
		}
	}

	return count
}

func nextWorldState(world01 *[maxRows][maxCols]int, world02 *[maxRows][maxCols]int) {
	neighbour_live_cell := 0
	
	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			neighbour_live_cell = countLiveNeighbourCell(*world01, i, j)

			if (*world01)[i][j]==1 && (neighbour_live_cell==2 || neighbour_live_cell==3) {
                world02[i][j] = 1
            } else if (*world01)[i][j]==0 && neighbour_live_cell==3 {
                world02[i][j] = 1
            } else {
                world02[i][j] = 0
            }
		}
	}
}

func main() {
	world01 := [maxRows][maxCols]int{}
	world02 := [maxRows][maxCols]int{}

	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			world01[i][j] = rand.Intn(32767) % 2
		}
	}

	nextWorldState(&world01, &world02)
}