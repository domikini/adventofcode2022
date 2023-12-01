package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pair struct {
	Row1       [][][]Packet
	Row2       [][][]Packet
	RightOrder bool
}

type Packet struct {
	Value int
	Level int
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	pairs := parseInput(string(content))

	fmt.Println(pairs)
}

func parseRawString(raw string) []interface{} {
	ans := []interface{}{}
	json.Unmarshal([]byte(raw), &ans)
	return ans
}

func parseInput(input string) (ans [][2][]interface{}) {
	for _, packetPairs := range strings.Split(input, "\n\n") {
		pairs := strings.Split(packetPairs, "\n")
		ans = append(ans, [2][]interface{}{
			parseRawString(pairs[0]),
			parseRawString(pairs[1]),
		})
	}
	return ans
}
