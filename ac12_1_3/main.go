package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Step struct {
	Row       int
	Column    int
	Elevation byte
	Up        *Step
	Right     *Step
	Down      *Step
	Left      *Step
	Parent    *Step
	Visited   bool
}

func NewStep(row int, column int, number_of_rows int, number_of_columns int) *Step {
	if row < 0 || row > number_of_rows-1 || column < 0 || column > number_of_columns-1 {
		return nil
	}
	return &Step{Row: row, Column: column}
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var height_map [41][143]Step
	content_split := strings.Split(string(content), "\n")
	for i := 0; i < len(content_split); i++ {
		for j := 0; j < len(content_split[i]); j++ {
			height_map[i][j].Elevation = content_split[i][j]
			height_map[i][j].Row = i
			height_map[i][j].Column = j
			height_map[i][j].Up = NewStep(i-1, j, len(content_split), len(content_split[i]))
			height_map[i][j].Right = NewStep(i, j+1, len(content_split), len(content_split[i]))
			height_map[i][j].Down = NewStep(i+1, j, len(content_split), len(content_split[i]))
			height_map[i][j].Left = NewStep(i, j-1, len(content_split), len(content_split[i]))
		}

	}

	// Populate height_map with data
	for i := 0; i < len(height_map); i++ {
		for j := 0; j < len(height_map[0]); j++ {
			if height_map[i][j].Up != nil {
				height_map[i][j].Up.Elevation = height_map[i-1][j].Elevation
				height_map[i][j].Up = &height_map[i-1][j]
			}
			if height_map[i][j].Right != nil {
				height_map[i][j].Right.Elevation = height_map[i][j+1].Elevation
				height_map[i][j].Right = &height_map[i][j+1]
			}
			if height_map[i][j].Down != nil {
				height_map[i][j].Down.Elevation = height_map[i+1][j].Elevation
				height_map[i][j].Down = &height_map[i+1][j]
			}
			if height_map[i][j].Left != nil {
				height_map[i][j].Left.Elevation = height_map[i][j-1].Elevation
				height_map[i][j].Left = &height_map[i][j-1]
			}
		}
	}

}

func BFS(step *Step) []int {
	queue := []*Step{}
	queue = append(queue, step)
	result := []int{}
	return BFSUtil(queue, result)
}

func BFSUtil(queue []*Step, res []int) []int {
	if len(queue) == 0 {
		return res
	}
	if queue[0].Up != nil {
		queue = append(queue, queue[0].Up)
	}

	return BFSUtil(queue[1:], res)
}

func PrintHeightMap(height_map [41][143]Step) {
	fmt.Println("")
	for i := 0; i < len(height_map); i++ {
		fmt.Println("")
		for j := 0; j < len(height_map[0]); j++ {
			if height_map[i][j].Visited == true {
				fmt.Print("X")
			} else {
				fmt.Print("0")
			}
		}
	}
	fmt.Println("")
}
