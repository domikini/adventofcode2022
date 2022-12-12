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
	number                int64
	items                 []int64
	operation_operator    string
	operation_number      int64
	test_divisbile_number int64
	if_true_to_monkey     int64
	if_false_to_monkey    int64
	inspection_times      int64
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

	lcm := int64(1)
	for _, monkey := range monkeys {
		lcm = lcm * monkey.test_divisbile_number
	}

	rounds := 10000
	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			for k := 0; k < len(monkeys[j].items); k++ {
				monkey := monkeys[j]
				item := monkey.items[k]
				monkeys[j].inspection_times++
				operation_number := int64(0)
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
				item = item % lcm
				if item%monkey.test_divisbile_number == 0 {
					monkeys[monkey.if_true_to_monkey].items = append(monkeys[monkey.if_true_to_monkey].items, item)
					monkey.items = RemoveIndex(monkey.items, len(monkey.items)-1)
				} else {
					monkeys[monkey.if_false_to_monkey].items = append(monkeys[monkey.if_false_to_monkey].items, item)
					monkey.items = RemoveIndex(monkey.items, len(monkey.items)-1)
				}
			}
			// Resetting items for the monkey
			monkeys[j].items = []int64{}
		}
	}

	inspection_times := []int64{}
	for _, monkey := range monkeys {
		inspection_times = append(inspection_times, monkey.inspection_times)
	}

	//Find max value and second max value and multiply them
	max_value := int64(0)
	second_max_value := int64(0)
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
		monkey.number, _ = strconv.ParseInt(result[1], 10, 64)
		items_list := strings.Split(result[2], ", ")
		for _, i := range items_list {
			j, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				panic(err)
			}
			monkey.items = append(monkey.items, j)
		}
		monkey.operation_operator = result[3]
		monkey.operation_number, _ = strconv.ParseInt(result[4], 10, 64)
		monkey.test_divisbile_number, _ = strconv.ParseInt(result[5], 10, 64)
		monkey.if_true_to_monkey, _ = strconv.ParseInt(result[6], 10, 64)
		monkey.if_false_to_monkey, _ = strconv.ParseInt(result[7], 10, 64)
		return monkey
	}
	return Monkey{}
}

func RemoveIndex(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}
