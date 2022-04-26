package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func toFixed(value float64, n float64) string {
	return strconv.FormatFloat(value/n, 'f', 6, 64)
}

func calculateProportions(arr []int32) []string {
	n := float64(len(arr))

	var positives, negatives, zeros float64 = 0, 0, 0

	for _, value := range arr {
		if value > 0 {
			positives += 1
		} else if value < 0 {
			negatives += 1
		} else if value == 0 {
			zeros += 1
		}
	}

	return []string{toFixed(positives, n), toFixed(negatives, n), toFixed(zeros, n)}
}

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	for _, value := range calculateProportions(arr) {
		fmt.Println(value)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

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
