package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	contentSplit := strings.Split(string(content), "\n")
	sumArray := SumOfLines(contentSplit)
	iterate := 3
	fmt.Println(Iterator(sumArray, iterate))
}

func SumOfLines(input []string) []int64 {
	var sumArray []int64
	var sum int64
	for _, value := range input {
		if value != "" {
			valueInt, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				fmt.Println("Error during conversion")
				break
			}
			sum += valueInt
		} else {
			fmt.Println(sum)
			sumArray = append(sumArray, sum)
			sum = 0
		}
	}
	return sumArray
}

func Iterator(arr []int64, iterate int) int64 {
	var sum int64 = 0
	var max_sum int64 = 0
	for i := 0; i < iterate; i++ {
		max_sum, arr = FindMaxElementAndRemoveIt(arr)
		sum += max_sum
	}
	return sum
}

func FindMaxElementAndRemoveIt(arr []int64) (int64, []int64) {
	max_num := arr[0]
	index_of_max_num := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
			index_of_max_num = i
		}
	}
	RemoveIndex(arr, index_of_max_num)
	return max_num, arr
}

func RemoveIndex(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}
