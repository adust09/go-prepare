package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var num_of_positive_elements = 0
	var num_of_negative_elements = 0
	var num_of_zero_elements = 0
	var num_of_elements = 0

	for i, element := range arr {
		if element < 0 {
			num_of_negative_elements++
		} else if element > 0 {
			num_of_positive_elements++
		}
		if element == 0 {
			num_of_zero_elements++
		}
		num_of_elements = i
	}
	num_of_elements = len(arr)
	var rate_of_positive = float64(num_of_positive_elements) / float64(num_of_elements)
	var rate_of_negative = float64(num_of_negative_elements) / float64(num_of_elements)
	var rate_of_zero = float64(num_of_zero_elements) / float64(num_of_elements)

	fmt.Printf("%f\n", rate_of_positive)
	fmt.Printf("%f\n", rate_of_negative)
	fmt.Printf("%f\n", rate_of_zero)

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
