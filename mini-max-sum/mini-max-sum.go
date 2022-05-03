package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortableInt64 []int64

func (arr SortableInt64) Len() int {
	return len(arr)
}

func (arr SortableInt64) Less(a, b int) bool {
	return arr[a] < arr[b]
}

func (arr SortableInt64) Swap(a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func remove(slice []int64, index int) []int64 {
	return append(slice[:index], slice[index+1:]...)
}

func calcSums(arr []int32) []int64 {
	lastIndex := len(arr) - 1

	parsedArr := make(SortableInt64, int(len(arr)))

	for index, value := range arr {
		parsedArr[index] = int64(value)
	}

	sort.Sort(parsedArr)

	smallest, greatest, rest := parsedArr[0], parsedArr[lastIndex], parsedArr[1:lastIndex]

	accumulator := int64(0)

	for _, value := range rest {
		accumulator += value
	}

	return []int64{accumulator + smallest, accumulator + greatest}
}

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	summed := calcSums(arr)

	fmt.Println(summed[0], summed[1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
