package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Tree struct {
	height       int
	scenic_score int
}

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	content_split := strings.Split(string(content), "\n")
	tree_array := [][]Tree{}
	for _, row := range content_split {
		tree_row := ReadRow(row)
		tree_array = append(tree_array, tree_row)
	}

	for i := 0; i < len(tree_array); i++ {
		for j := 0; j < len(tree_array); j++ {
			tree_array[i][j].scenic_score = ScenicScoreLookup(i, j, tree_array)
		}
	}

	// Find max scenic_score
	max_scenic_score := 0
	for i := 0; i < len(tree_array); i++ {
		for j := 0; j < len(tree_array); j++ {
			if max_scenic_score < tree_array[i][j].scenic_score {
				max_scenic_score = tree_array[i][j].scenic_score
			}
		}
	}

	fmt.Println(max_scenic_score)
}

func ReadRow(row string) []Tree {
	row_array := []Tree{}
	for _, value := range row {
		tree := Tree{height: int(value), scenic_score: 0}
		row_array = append(row_array, tree)
	}
	return row_array
}

func ScenicScoreLookup(row int, column int, tree_array [][]Tree) int {
	right_counter := 0
	left_counter := 0
	top_counter := 0
	bottom_counter := 0
	current_tree_height := tree_array[row][column].height

	// To right
	for i := column + 1; i < len(tree_array[0]); i++ {
		right_counter++
		if current_tree_height <= tree_array[row][i].height {
			break
		}
	}
	// To left
	for i := column - 1; i >= 0; i-- {
		left_counter++
		if current_tree_height <= tree_array[row][i].height {
			break
		}
	}
	// To top
	for i := row + 1; i < len(tree_array); i++ {
		top_counter++
		if current_tree_height <= tree_array[i][column].height {
			break
		}
	}
	// To bottom
	for i := row - 1; i >= 0; i-- {
		bottom_counter++
		if current_tree_height <= tree_array[i][column].height {
			break
		}
	}

	scenic_score := right_counter * left_counter * top_counter * bottom_counter
	return scenic_score
}
