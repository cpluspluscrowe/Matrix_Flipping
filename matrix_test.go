package main

import (
	"fmt"
	"testing"
)

func TestSumTopCorner(t *testing.T) {
	square := [][]int{{1, 2}, {3, 4}}
	var sum = SumTopCorner(square)
	if sum != 1 {
		t.Errorf("Sum should be 1, but was %d", sum)
	}
}

func TestFlipRow(t *testing.T) {
	row := []int{1, 2, 3}
	var flipped = flipRow(row)
	if flipped[0] != 3 && flipped[1] != 2 && flipped[2] != 3 {
		t.Errorf("Row flip was unsuccessful")
	}
}

func TestFlipColumn(t *testing.T) {
	square := [][]int{{1, 2}, {3, 4}}
	var flipped = flipColumn(square, 1)
	fmt.Println(square)
	if flipped[0][1] != 4 || flipped[1][1] != 2 {
		t.Errorf("Column flip was unsuccessful")
	}
}
