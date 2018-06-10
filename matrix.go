package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
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

func runGame() {
	//	var largestNumbers []int
	var n = readNumber(reader)
	square := make([][]int, n)
	for line := 0; line < 2*n; line++ {
		square[line] = getNextRow()
	}
}

func flipColumn(column []int) []int {
	fmt.Println(column)
	var capacity int = cap(column)
	for i := 0; i < capacity/2; i++ {
		var oppositeIndex int = capacity - i - 1
		var temp = column[i]
		column[i] = column[oppositeIndex]
		column[oppositeIndex] = temp
	}
	return column
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

func TestSimple(t *testing.T) {

}
