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
				PrintToCRT(cycle, WhichPixel(x_value, cycle))

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
			PrintToCRT(cycle, WhichPixel(x_value, cycle))

			if CheckIfNeedRead(cycle) {
				signal_strength += x_value * cycle
			}
		}
	}
}

func WhichPixel(x_value int, cycle int) string {
	crt_pixel_position := cycle % 40
	if crt_pixel_position >= x_value && crt_pixel_position <= x_value+2 {
		return "#"
	}
	return "."
}

func PrintToCRT(cycle int, pixel string) {
	if CheckIfPrintNewLine(cycle) {
		fmt.Println(pixel)
	} else {
		fmt.Print(pixel)
	}
}

func CheckIfNeedRead(cycle int) bool {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return true
	}
	return false
}

func CheckIfPrintNewLine(cycle int) bool {
	if cycle == 40 || cycle == 80 || cycle == 120 || cycle == 160 || cycle == 200 || cycle == 240 {
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
