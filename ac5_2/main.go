package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	stacks := InitialStacks()
	fmt.Println(stacks)

	content_split := strings.Split(string(content), "\n")
	for _, row := range content_split {
		amount, from, to := ReadRow(row)
		// Minus 1 to match array indexes
		stacks = Move(stacks, amount, from-1, to-1)
	}

	// Display last crate on each stack
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
}

func InitialStacks() [][]string {
	stack1 := []string{"W", "M", "L", "F"}
	stack2 := []string{"B", "Z", "V", "M", "F"}
	stack3 := []string{"H", "V", "R", "S", "L", "Q"}
	stack4 := []string{"F", "S", "V", "Q", "P", "M", "T", "J"}
	stack5 := []string{"L", "S", "W"}
	stack6 := []string{"F", "V", "P", "M", "R", "J", "W"}
	stack7 := []string{"J", "Q", "C", "P", "N", "R", "F"}
	stack8 := []string{"V", "H", "P", "S", "Z", "W", "R", "B"}
	stack9 := []string{"B", "M", "J", "C", "G", "H", "Z", "W"}
	return [][]string{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9}
}

func ReadRow(row string) (int, int, int) {
	re := regexp.MustCompile(`(\d+).*(\d+).*(\d+)`)
	result := re.FindStringSubmatch(row)
	amount, _ := strconv.Atoi(result[1])
	from, _ := strconv.Atoi(result[2])
	to, _ := strconv.Atoi(result[3])
	return amount, from, to
}

func Move(stacks [][]string, amount int, from int, to int) [][]string {
	start_position := len(stacks[from]) - amount
	stacks_to_move := stacks[from][start_position:]
	stacks[to] = append(stacks[to], stacks_to_move...)
	stacks[from] = stacks[from][:start_position]
	return stacks
}
