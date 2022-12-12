package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	number                int
	items                 []int
	operation_operator    string
	operation_number      int
	test_divisbile_number int
	if_true_to_monkey     int
	if_false_to_monkey    int
	inspection_times      int
}

func main() {

	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	monkeys := []Monkey{}
	content_split := strings.Split(string(content), "\n")
	for i := 0; i < len(content_split); i = i + 7 {
		monkey := ReadRow(content_split[i] + content_split[i+1] + content_split[i+2] + content_split[i+3] + content_split[i+4] + content_split[i+5])
		monkeys = append(monkeys, monkey)
	}

	rounds := 20
	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			for k := 0; k < len(monkeys[j].items); k++ {
				monkey := monkeys[j]
				item := monkey.items[k]
				monkeys[j].inspection_times++
				operation_number := 0
				if monkey.operation_number == 0 {
					operation_number = item
				} else {
					operation_number = monkey.operation_number
				}
				switch monkey.operation_operator {
				case "*":
					item = item * operation_number
				case "+":
					item = item + operation_number
				}
				item = item / 3
				if item%monkey.test_divisbile_number == 0 {
					monkeys[monkey.if_true_to_monkey].items = append(monkeys[monkey.if_true_to_monkey].items, item)
					monkey.items = RemoveIndex(monkey.items, len(monkey.items)-1)
				} else {
					monkeys[monkey.if_false_to_monkey].items = append(monkeys[monkey.if_false_to_monkey].items, item)
					monkey.items = RemoveIndex(monkey.items, len(monkey.items)-1)
				}
			}
			// Resetting items for the monkey
			monkeys[j].items = []int{}
		}
	}

	inspection_times := []int{}
	for _, monkey := range monkeys {
		inspection_times = append(inspection_times, monkey.inspection_times)
	}

	// Find max value and second max value and multiply them
	max_value := 0
	second_max_value := 0
	for i, e := range inspection_times {
		if i == 0 || e > max_value {
			max_value = e
		}
	}

	for i, e := range inspection_times {
		if e == max_value {
			inspection_times[i] = 0
		}
	}

	for i, e := range inspection_times {
		if i == 0 || e > second_max_value {
			second_max_value = e
		}
	}

	fmt.Println(max_value * second_max_value)
}

func ReadRow(row string) Monkey {
	re := regexp.MustCompile(`Monkey (.*):  Starting items: (.*)  Operation: new = old (.*) (.*)  Test: divisible by (.*)    If true: throw to monkey (.*)    If false: throw to monkey (.*)`)
	result := re.FindStringSubmatch(row)
	if len(result) > 0 {
		monkey := Monkey{}
		monkey.number, _ = strconv.Atoi(result[1])
		items_list := strings.Split(result[2], ", ")
		for _, i := range items_list {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			monkey.items = append(monkey.items, j)
		}
		monkey.operation_operator = result[3]
		monkey.operation_number, _ = strconv.Atoi(result[4])
		monkey.test_divisbile_number, _ = strconv.Atoi(result[5])
		monkey.if_true_to_monkey, _ = strconv.Atoi(result[6])
		monkey.if_false_to_monkey, _ = strconv.Atoi(result[7])
		return monkey
	}
	return Monkey{}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
