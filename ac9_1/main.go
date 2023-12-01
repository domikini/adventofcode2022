package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	direction := ""
	steps := 0
	head_position := Position{0, 0}
	tail_position := Position{0, 0}
	tail_positions := []Position{}

	content_split := strings.Split(string(content), "\n")
	for i := 0; i < len(content_split); i++ {
		direction, steps = ReadRow(content_split[i])
		head_position, tail_position, tail_positions = HeadAndTailMovement(head_position, tail_position, steps, direction, tail_positions)
	}

	// Collect unique tail positions
	unique := []Position{}
	for _, v := range tail_positions {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	fmt.Println(len(unique))
}

func HeadAndTailMovement(head_position Position, tail_position Position, steps int, direction string, tail_positions []Position) (Position, Position, []Position) {
	if steps == 0 {
		return head_position, tail_position, tail_positions
	}
	switch direction {
	case "R":
		head_position.x++
	case "L":
		head_position.x--
	case "U":
		head_position.y++
	case "D":
		head_position.y--
	}
	tail_position = TailReAlign(head_position, tail_position)
	tail_positions = append(tail_positions, tail_position)
	steps--
	return HeadAndTailMovement(head_position, tail_position, steps, direction, tail_positions)
}

func TailReAlign(head_position Position, tail_position Position) Position {
	if (head_position.x-tail_position.x <= 1 && head_position.x-tail_position.x >= -1) &&
		(head_position.y-tail_position.y <= 1 && head_position.y-tail_position.y >= -1) {
		return tail_position
	} else if head_position.x-tail_position.x == 0 && head_position.y-tail_position.y == 2 {
		return Position{tail_position.x, head_position.y - 1}
	} else if head_position.y-tail_position.y == 0 && head_position.x-tail_position.x == 2 {
		return Position{head_position.x - 1, tail_position.y}
	} else if head_position.x-tail_position.x == 0 && head_position.y-tail_position.y == -2 {
		return Position{tail_position.x, head_position.y + 1}
	} else if head_position.y-tail_position.y == 0 && head_position.x-tail_position.x == -2 {
		return Position{head_position.x + 1, tail_position.y}
	} else if (head_position.x-tail_position.x == 1 && head_position.y-tail_position.y == 2) ||
		(head_position.x-tail_position.x == 2 && head_position.y-tail_position.y == 1) {
		return Position{tail_position.x + 1, tail_position.y + 1}
	} else if (head_position.x-tail_position.x == 1 && head_position.y-tail_position.y == -2) ||
		(head_position.x-tail_position.x == 2 && head_position.y-tail_position.y == -1) {
		return Position{tail_position.x + 1, tail_position.y - 1}
	} else if (head_position.x-tail_position.x == -1 && head_position.y-tail_position.y == 2) ||
		(head_position.x-tail_position.x == -2 && head_position.y-tail_position.y == 1) {
		return Position{tail_position.x - 1, tail_position.y + 1}
	} else if (head_position.x-tail_position.x == -1 && head_position.y-tail_position.y == -2) ||
		(head_position.x-tail_position.x == -2 && head_position.y-tail_position.y == -1) {
		return Position{tail_position.x - 1, tail_position.y - 1}
	}
	return tail_position
}

func ReadRow(row string) (string, int) {
	re := regexp.MustCompile(`(.*) (.*)`)
	result := re.FindStringSubmatch(row)
	if len(result) != 0 {
		steps, _ := strconv.Atoi(result[2])
		return result[1], steps
	}
	return "", 0
}
