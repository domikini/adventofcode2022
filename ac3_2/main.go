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
	total_sum := 0
	for i := 0; i < len(content_split); i += 3 {
	out:
		for _, item1 := range content_split[i] {
			for _, item2 := range content_split[i+1] {
				for _, item3 := range content_split[i+2] {
					if item1 == item2 && item1 == item3 && item2 == item3 {
						total_sum += GetItemTypePriority(byte(item1))
						break out
					}
				}
			}
		}
	}
	fmt.Println(total_sum)
}

func GetItemTypePriority(item byte) int {
	switch item {
	case 65:
		return 27
	case 66:
		return 28
	case 67:
		return 29
	case 68:
		return 30
	case 69:
		return 31
	case 70:
		return 32
	case 71:
		return 33
	case 72:
		return 34
	case 73:
		return 35
	case 74:
		return 36
	case 75:
		return 37
	case 76:
		return 38
	case 77:
		return 39
	case 78:
		return 40
	case 79:
		return 41
	case 80:
		return 42
	case 81:
		return 43
	case 82:
		return 44
	case 83:
		return 45
	case 84:
		return 46
	case 85:
		return 47
	case 86:
		return 48
	case 87:
		return 49
	case 88:
		return 50
	case 89:
		return 51
	case 90:
		return 52
	case 97:
		return 1
	case 98:
		return 2
	case 99:
		return 3
	case 100:
		return 4
	case 101:
		return 5
	case 102:
		return 6
	case 103:
		return 7
	case 104:
		return 8
	case 105:
		return 9
	case 106:
		return 10
	case 107:
		return 11
	case 108:
		return 12
	case 109:
		return 13
	case 110:
		return 14
	case 111:
		return 15
	case 112:
		return 16
	case 113:
		return 17
	case 114:
		return 18
	case 115:
		return 19
	case 116:
		return 20
	case 117:
		return 21
	case 118:
		return 22
	case 119:
		return 23
	case 120:
		return 24
	case 121:
		return 25
	case 122:
		return 26
	}
	return 0
}
