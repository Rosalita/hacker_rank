package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {

	size := int32(len(arr))

	countPos, countNeg, countZero := int32(0), int32(0), int32(0)

	for _, v := range arr {

		if isPos(v) {
			countPos++
		} else if isNeg(v) {
			countNeg++
		} else {
			countZero++
		}
	}

	posFrac := float64(countPos) / float64(size)
	negFrac := float64(countNeg) / float64(size)
	zeroFrac := float64(countZero) / float64(size)

	fmt.Println(posFrac)
	fmt.Println(negFrac)
	fmt.Println(zeroFrac)

}

func isPos(i int32) bool {
	return i > 0
}

func isNeg(i int32) bool {
	return i < 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
