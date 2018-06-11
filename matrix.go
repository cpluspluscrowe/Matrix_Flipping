package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNumber(reader *bufio.Reader) int {
	var input, _, _ = reader.ReadLine()
	var number, _ = strconv.Atoi(string(input))
	return number
}

func getNextRow() []int {
	var line, _, _ = reader.ReadLine()
	var rowStrings []string = strings.Split(string(line), " ")
	var numbers []int
	for _, value := range rowStrings {
		var number, _ = strconv.Atoi(value)
		numbers = append(numbers, number)
	}
	return numbers
}

func fillMatrix(n int) [][]int {
	square := [][]int{}
	for line := 0; line < 2*n; line++ {
		square = append(square, getNextRow())
	}
	return square
}

func runGame() {
	var n = readNumber(reader)
	var matrix = fillMatrix(n)
	var currentSum = sumTopCorner(matrix)
	for true {
		var changed = false
		for rowIndex, _ := range matrix {
			flipRow(matrix, rowIndex)
			var newSum int = sumTopCorner(matrix)
			if newSum > currentSum {
				changed = true
				currentSum = newSum
			} else {
				flipRow(matrix, rowIndex)
			}
		}
		for columnIndex, _ := range matrix {
			flipColumn(matrix, columnIndex)
			var newSum int = sumTopCorner(matrix)
			if newSum > currentSum {
				changed = true
				currentSum = newSum
			} else {
				flipRow(matrix, columnIndex)
			}
		}
		if changed == false {
			break
		}
	}
	fmt.Println(currentSum)
}

func sumTopCorner(square [][]int) int {
	var length = len(square)
	var sum int = 0
	for i := 0; i < length/2; i++ {
		for j := 0; j < length/2; j++ {
			sum += square[i][j]
		}
	}
	return sum
}

func flipRow(square [][]int, rowIndex int) [][]int {
	var capacity int = len(square)
	for i := 0; i < capacity/2; i++ {
		var oppositeIndex int = capacity - i - 1
		var temp = square[rowIndex][i]
		square[rowIndex][i] = square[rowIndex][oppositeIndex]
		square[rowIndex][oppositeIndex] = temp
	}
	return square
}

func flipColumn(square [][]int, columnIndex int) [][]int {
	var capacity int = len(square)
	for i := 0; i < capacity/2; i++ {
		var oppositeIndex int = capacity - i - 1
		var temp = square[i][columnIndex]
		square[i][columnIndex] = square[oppositeIndex][columnIndex]
		square[oppositeIndex][columnIndex] = temp
	}
	return square
}

var reader = bufio.NewReaderSize(os.Stdin, 1024*1024)

func main() {
	var numberOfProblems = readNumber(reader)
	//var c []string = strings.Split(string(input), " ")
	for i := 0; i < numberOfProblems; i++ {
		runGame()
	}
}
