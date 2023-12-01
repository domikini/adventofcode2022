package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	content_split := strings.Split(string(content), "\n")
	row_string := []string{}
	for _, row := range content_split {
		row_string = ReadRow(row)
	}

	character_count := 0
	type void struct{}
	var member void
end:
	for c := 0; c < len(row_string)-14; c++ {
		compare_slice := row_string[c : c+14]
		set := make(map[string]void)
		for _, character := range compare_slice {
			set[character] = member
		}
		if len(set) == 14 {
			character_count = c + 14
			break end
		}
	}

	fmt.Println(character_count)

}

func ReadRow(row string) []string {
	return strings.Split(row, "")
}
