package main

import (
	"fmt"
	"testing"
)

func TestSumTopCorner(t *testing.T) {
	square := [][]int{{1, 2}, {3, 4}}
	var sum = sumTopCorner(square)
	if sum != 1 {
		t.Errorf("Sum should be 1, but was %d", sum)
	}
}

func TestFlipRow(t *testing.T) {
	square := [][]int{{1, 2}, {3, 4}}
	var flipped = flipRow(square, 1)
	if flipped[1][0] != 4 || flipped[1][1] != 3 {
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

func TestFlipColumn2(t *testing.T) {
	square := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var flipped = flipColumn(square, 1)
	fmt.Println(square)
	if flipped[0][1] != 8 || flipped[1][1] != 5 || flipped[2][1] != 2 {
		t.Errorf("Column flip was unsuccessful")
	}
}
