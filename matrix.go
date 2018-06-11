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
	square := make([][]int, n)
	for line := 0; line < 2*n; line++ {
		square[line] = getNextRow()
	}
	return square
}

func runGame() {
	var n = readNumber(reader)
	var matrix = fillMatrix(n)
	var changed = false
	var currentSum = SumTopCorner(matrix)
	for true {
		fmt.Println(changed, currentSum)
	}
	fmt.Println(matrix)
}

func SumTopCorner(square [][]int) int {
	var length = len(square)
	var sum int = 0
	for i := 0; i < length/2; i++ {
		for j := 0; j < length/2; j++ {
			sum += square[i][j]
		}
	}
	return sum
}

func flipRow(row []int) []int {
	var capacity int = cap(row)
	for i := 0; i < capacity/2; i++ {
		var oppositeIndex int = capacity - i - 1
		var temp = row[i]
		row[i] = row[oppositeIndex]
		row[oppositeIndex] = temp
	}
	return row
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
	fmt.Println("vim-go")
}
