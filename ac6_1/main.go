package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	content_split := strings.Split(string(content), "\n")
	for _, row := range content_split {
		fmt.Println(ReadRow(row))
	}

}

func ReadRow(row string) string {
	re := regexp.MustCompile(`(\$) (cd)(.*)`)
	result := re.FindStringSubmatch(row)
	if len(result) != 0 {
		return result[0]
	}
	re = regexp.MustCompile(`(\$) (ls)`)
	result = re.FindStringSubmatch(row)
	if len(result) != 0 {
		return result[0]
	}
	return ""
}
