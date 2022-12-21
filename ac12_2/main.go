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
	Visited   bool
}

type QueueObject struct {
	Step  *Step
	Count int
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

	var height_map [41][147]Step
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

	start_points := [][2]int{}
	for i := 0; i < len(height_map); i++ {
		for j := 0; j < len(height_map[i]); j++ {
			if height_map[i][j].Elevation == 83 {
				start_points = append(start_points, [2]int{i, j})
			} else if height_map[i][j].Elevation == 97 {
				start_points = append(start_points, [2]int{i, j})
			}
		}
	}

	var path_lengths []int
	for _, start_point := range start_points {
		path_lengths = append(path_lengths, FindShortestPath(height_map, start_point[0], start_point[1]))
	}

	fmt.Println(findMinElement(path_lengths))
}

func FindShortestPath(height_map [41][147]Step, first_step_row int, first_step_column int) int {
	queue := []QueueObject{}
	queue = append(queue, QueueObject{Step: &height_map[first_step_row][first_step_column], Count: 0})
	visited := map[[2]int]bool{}
	path_length := 0

found:
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if visited[[2]int{first.Step.Row, first.Step.Column}] {
			continue
		}
		visited[[2]int{first.Step.Row, first.Step.Column}] = true
		if first.Step.Elevation == 69 {
			path_length = first.Count
			break found
		}
		if first.Step.Right != nil {
			if (first.Step.Right.Elevation <= first.Step.Elevation && first.Step.Right.Elevation >= byte('a')) ||
				first.Step.Elevation+1 == first.Step.Right.Elevation ||
				first.Step.Elevation-first.Step.Right.Elevation == 52 ||
				first.Step.Elevation-first.Step.Right.Elevation == 53 ||
				first.Step.Elevation+first.Step.Right.Elevation == 180 ||
				first.Step.Elevation+first.Step.Right.Elevation == 181 {
				queue = append(queue, QueueObject{Step: first.Step.Right, Count: first.Count + 1})
			}
		}
		if first.Step.Up != nil {
			if (first.Step.Up.Elevation <= first.Step.Elevation && first.Step.Up.Elevation >= byte('a')) ||
				first.Step.Elevation+1 == first.Step.Up.Elevation ||
				first.Step.Elevation-first.Step.Up.Elevation == 52 ||
				first.Step.Elevation-first.Step.Up.Elevation == 53 ||
				first.Step.Elevation+first.Step.Up.Elevation == 180 ||
				first.Step.Elevation+first.Step.Up.Elevation == 181 {
				queue = append(queue, QueueObject{Step: first.Step.Up, Count: first.Count + 1})
			}
		}
		if first.Step.Down != nil {
			if (first.Step.Down.Elevation <= first.Step.Elevation && first.Step.Down.Elevation >= byte('a')) ||
				first.Step.Elevation+1 == first.Step.Down.Elevation ||
				first.Step.Elevation-first.Step.Down.Elevation == 52 ||
				first.Step.Elevation-first.Step.Down.Elevation == 53 ||
				first.Step.Elevation+first.Step.Down.Elevation == 180 ||
				first.Step.Elevation+first.Step.Down.Elevation == 181 {
				queue = append(queue, QueueObject{Step: first.Step.Down, Count: first.Count + 1})
			}
		}
		if first.Step.Left != nil {
			if (first.Step.Left.Elevation <= first.Step.Elevation && first.Step.Left.Elevation >= byte('a')) ||
				first.Step.Elevation+1 == first.Step.Left.Elevation ||
				first.Step.Elevation-first.Step.Left.Elevation == 52 ||
				first.Step.Elevation-first.Step.Left.Elevation == 53 ||
				first.Step.Elevation+first.Step.Left.Elevation == 180 ||
				first.Step.Elevation+first.Step.Left.Elevation == 181 {
				queue = append(queue, QueueObject{Step: first.Step.Left, Count: first.Count + 1})
			}
		}
	}
	return path_length
}

func findMinElement(arr []int) int {
	min_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min_num && arr[i] != 0 {
			min_num = arr[i]
		}
	}
	return min_num
}
