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
	sumArray := sumOfLines(contentSplit)
	fmt.Println("\n" + strconv.Itoa(int(findMaxElement(sumArray))))
}

func sumOfLines(input []string) []int64 {
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

func findMaxElement(arr []int64) int64 {
	max_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}
