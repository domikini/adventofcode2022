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
	overlap_counter := 0
	content_split := strings.Split(string(content), "\n")
	for _, row := range content_split {
		if ReadRow(row) {
			overlap_counter++
		}
	}
	fmt.Println(overlap_counter)
}

func ReadRow(row string) bool {
	row_split := strings.Split(row, ",")
	pair_1_start, pair_1_end := GetStartAndEnd(row_split[0])
	pair_2_start, pair_2_end := GetStartAndEnd(row_split[1])
	overlap := OverlapChecker(pair_1_start, pair_1_end, pair_2_start, pair_2_end)
	return overlap
}

func GetStartAndEnd(sections string) (int, int) {
	section_split := strings.Split(sections, "-")
	section_start, err := strconv.Atoi(section_split[0])
	section_end, err2 := strconv.Atoi(section_split[1])

	if err != nil || err2 != nil {
		fmt.Println("Error during conversion")
		return 0, 0
	}
	return section_start, section_end
}

func OverlapChecker(pair_1_start int, pair_1_end int, pair_2_start int, pair_2_end int) bool {
	overlap := false

	sections_1 := ArraySynthesizer(pair_1_start, pair_1_end)
	sections_2 := ArraySynthesizer(pair_2_start, pair_2_end)
	sections_1_overlap_counter := 0
	sections_2_overlap_counter := 0

	for i := 0; i < len(sections_1); i++ {
		for j := 0; j < len(sections_2); j++ {
			if sections_1[i] == sections_2[j] {
				sections_1_overlap_counter++
				sections_2_overlap_counter++
			}
		}
	}
	if sections_1_overlap_counter > 0 && sections_2_overlap_counter > 0 {
		overlap = true
	}
	return overlap
}

func ArraySynthesizer(start int, end int) []int {
	array := []int{}
	for i := start; i < end+1; i++ {
		array = append(array, i)
	}
	return array
}
