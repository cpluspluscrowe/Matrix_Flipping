package main

import (
	"fmt"
	"testing"
)

func TestFlipColumn(t *testing.T) {
	column := []int{1, 2, 3}
	var flipped = flipColumn(column)
	fmt.Println(flipped)
	if flipped[0] != 3 && flipped[1] != 2 && flipped[2] != 3 {
		t.Errorf("Column flip was unsuccessful")
	}
}
