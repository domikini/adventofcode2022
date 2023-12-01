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
	position1 := Position{0, 0}
	position2 := Position{0, 0}
	position3 := Position{0, 0}
	position4 := Position{0, 0}
	position5 := Position{0, 0}
	position6 := Position{0, 0}
	position7 := Position{0, 0}
	position8 := Position{0, 0}
	tail_position := Position{0, 0}
	rope := []Position{head_position, position1, position2, position3, position4, position5, position6, position7, position8, tail_position}
	tail_positions := []Position{}

	content_split := strings.Split(string(content), "\n")
	for i := 0; i < len(content_split); i++ {
		direction, steps = ReadRow(content_split[i])
		rope, tail_positions = HeadAndTailMovement(rope, steps, direction, tail_positions)
		fmt.Println(rope)
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

func HeadAndTailMovement(rope []Position, steps int, direction string, tail_positions []Position) ([]Position, []Position) {
	if steps == 0 {
		return rope, tail_positions
	}
	switch direction {
	case "R":
		rope[0].x++
	case "L":
		rope[0].x--
	case "U":
		rope[0].y++
	case "D":
		rope[0].y--
	}
	for i := 1; i < len(rope); i++ {
		rope[i] = KnotRealign(rope[i-1], rope[i])
	}
	tail_positions = append(tail_positions, rope[9])
	steps--
	return HeadAndTailMovement(rope, steps, direction, tail_positions)
}

func KnotRealign(head_position Position, tail_position Position) Position {
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
	} else if head_position.x-tail_position.x == 2 && head_position.y-tail_position.y == 2 {
		return Position{tail_position.x + 1, tail_position.y + 1}
	} else if head_position.x-tail_position.x == -2 && head_position.y-tail_position.y == 2 {
		return Position{tail_position.x - 1, tail_position.y + 1}
	} else if head_position.x-tail_position.x == -2 && head_position.y-tail_position.y == -2 {
		return Position{tail_position.x - 1, tail_position.y - 1}
	} else if head_position.x-tail_position.x == 2 && head_position.y-tail_position.y == -2 {
		return Position{tail_position.x + 1, tail_position.y - 1}
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
