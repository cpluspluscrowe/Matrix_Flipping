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

func getCorner(square [][]int) [][]int {
	var n = len(square) / 2
	maxValues := [][]int{}
	for line := 0; line < n; line++ {
		maxValues = append(maxValues, make([]int, n))
	}
	return maxValues
}

func GetMaxValueFromList(list []int) int {
	var max = -100000
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}

func GetMaxQuadrant(row int, column int, square [][]int) int {
	var lastIndex = len(square) - 1
	var quadrantValues []int
	quadrantValues = append(quadrantValues, square[row][column])
	quadrantValues = append(quadrantValues, square[row][lastIndex-column])
	quadrantValues = append(quadrantValues, square[lastIndex-row][column])
	quadrantValues = append(quadrantValues, square[lastIndex-row][lastIndex-column])
	return GetMaxValueFromList(quadrantValues)
}

func runGame() {
	var n = readNumber(reader)
	var matrix = fillMatrix(n)
	var sum = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += GetMaxQuadrant(i, j, matrix)
		}
	}

	fmt.Println(sum)
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
