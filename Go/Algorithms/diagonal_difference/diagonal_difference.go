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

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {

	primaryPoints := []int32{arr[0][0], arr[1][1], arr[2][2]}
	secondaryPoints := []int32{arr[0][2], arr[1][1], arr[2][0]}

	p := addPoints(primaryPoints)
	p = absInt(p)

	s := addPoints(secondaryPoints)
	s = absInt(s)

	diff := p - s

	return absInt(diff)

}

func absInt(i int32) int32 {
	absFloat := math.Abs(float64(i))
	return int32(absFloat)
}

func addPoints(points []int32) int32 {
	var sum int32
	for _, v := range points {
		sum += v
	}
	return sum
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
