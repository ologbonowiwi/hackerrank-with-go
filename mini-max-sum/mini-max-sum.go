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

type SortableInt32 []int32

func (arr SortableInt32) Len() int {
	return len(arr)
}

func (arr SortableInt32) Less(a, b int) bool {
	return arr[a] < arr[b]
}

func (arr SortableInt32) Swap(a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func remove(slice []int32, index int) []int32 {
	return append(slice[:index], slice[index+1:]...)
}

func calcSums(arr SortableInt32) []int32 {
	lastIndex := len(arr) - 1

	sort.Sort(arr)

	smallest, greatest, rest := arr[0], arr[lastIndex], arr[1:lastIndex]

	accumulator := int32(0)

	for _, value := range rest {
		accumulator += value
	}

	return []int32{accumulator + smallest, accumulator + greatest}
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
