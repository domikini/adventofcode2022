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
	cycle := 0
	x_value := 1
	signal_strength := 0
	content_split := strings.Split(string(content), "\n")
	for i := 0; i < len(content_split); i++ {
		input := ReadRow(content_split[i])
		if input != "noop" {
			for j := 0; j < 2; j++ {
				cycle++
				if CheckIfNeedRead(cycle) {
					signal_strength += (x_value * cycle)
				}
				if j == 1 {
					input_int, _ := strconv.Atoi(input)
					x_value += input_int
				}
			}
		} else {
			cycle++
			if CheckIfNeedRead(cycle) {
				signal_strength += x_value * cycle
			}
		}

	}
	fmt.Println(signal_strength)
}

func CheckIfNeedRead(cycle int) bool {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return true
	}
	return false
}

func ReadRow(row string) string {
	re := regexp.MustCompile(`(noop)`)
	result := re.FindStringSubmatch(row)
	if len(result) != 0 {
		return result[0]
	}
	re = regexp.MustCompile(`addx (.*)`)
	result = re.FindStringSubmatch(row)
	if len(result) != 0 {
		return result[1]
	}
	return ""
}
