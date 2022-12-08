package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Tree struct {
	height  int
	visible bool
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

	// Find visible trees left to right
	for i := 0; i < len(tree_array); i++ {
		highest_tree_so_far := 0
		for j := 0; j < len(tree_array[i]); j++ {
			if highest_tree_so_far < tree_array[i][j].height {
				tree_array[i][j].visible = true
				highest_tree_so_far = tree_array[i][j].height
			}
		}
	}

	// Find visible trees from right to left
	for i := 0; i < len(tree_array); i++ {
		highest_tree_so_far := 0
		for j := len(tree_array[i]) - 1; j >= 0; j-- {
			if highest_tree_so_far < tree_array[i][j].height {
				tree_array[i][j].visible = true
				highest_tree_so_far = tree_array[i][j].height
			}
		}
	}

	// Find visible trees from top to bottom
	for i := 1; i < len(tree_array[0])-1; i++ {
		highest_tree_so_far := 0
		for j := 0; j < len(tree_array); j++ {
			if highest_tree_so_far < tree_array[j][i].height {
				tree_array[j][i].visible = true
				highest_tree_so_far = tree_array[j][i].height
			}
		}
	}

	// Find visible trees from bottom to top
	for i := 1; i < len(tree_array[0])-1; i++ {
		highest_tree_so_far := 0
		for j := len(tree_array) - 1; j >= 0; j-- {
			if highest_tree_so_far < tree_array[j][i].height {
				tree_array[j][i].visible = true
				highest_tree_so_far = tree_array[j][i].height
			}
		}
	}

	total_visible_trees := 0
	for _, row := range tree_array {
		for _, tree := range row {
			if tree.visible == true {
				total_visible_trees += 1
			}
		}
	}
	fmt.Println(total_visible_trees)
}

func ReadRow(row string) []Tree {
	row_array := []Tree{}
	for _, value := range row {
		tree := Tree{height: int(value), visible: false}
		row_array = append(row_array, tree)
	}
	return row_array
}
