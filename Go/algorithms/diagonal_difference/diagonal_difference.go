package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func findPrimaryCoords(arr [][]int32) []coord {
	size := len(arr)
	coords := []coord{}
	for i := 0; i < int(size); i++ {
		coord := coord{
			x: i,
			y: i,
		}
		coords = append(coords, coord)
	}
	return coords
}

func findSecondaryCoords(arr [][]int32) []coord {
	size := len(arr)
	coords := []coord{}

	for i := 0; i < int(size); i++ {
		coord := coord{
			x: i,
			y: int(size) - 1 - i,
		}
		coords = append(coords, coord)
	}
	return coords
}

func getSumOfCoords(arr [][]int32, coords []coord) int32 {
	var sum int32
	for _, coord := range coords {
		sum += arr[coord.x][coord.y]
	}
	return sum

}

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {

	primaryDiagCoords := findPrimaryCoords(arr)
	secondaryDiagCoords := findSecondaryCoords(arr)

	primarySum := getSumOfCoords(arr, primaryDiagCoords)
	secondarySum := getSumOfCoords(arr, secondaryDiagCoords)

	difference := primarySum - secondarySum
	difference = absInt(difference)
	return difference
}

func absInt(i int32) int32 {
	absFloat := math.Abs(float64(i))
	return int32(absFloat)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
