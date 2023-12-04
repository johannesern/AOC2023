package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Day1Sol1()
	Day1Sol2()
}

func Day1Sol1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	var input []string
	for _, line := range lines {
		if line != "" {
			input = append(input, line)
		}
	}

	numbers := "0123456789"

	//map[string]int{}
	// newNumbers := map[string]int{
	// 	"zero":  0,
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// 	"four":  4,
	// 	"five":  5,
	// 	"six":   6,
	// 	"seven": 7,
	// 	"eight": 8,
	// 	"nine":  9,
	// }

	sum := 0

	for _, line := range input {
		lineInts := []string{}
		for _, char := range string(line) {
			if strings.Contains(numbers, string(char)) {
				lineInts = append(lineInts, string(char))
			}
		}
		if len(lineInts) == 1 {
			tmp, _ := strconv.Atoi(lineInts[0] + lineInts[0])
			sum += tmp
		}
		if len(lineInts) == 2 {
			tmp, _ := strconv.Atoi(lineInts[0] + lineInts[1])
			sum += tmp
		}
		if len(lineInts) > 2 {
			tmp, _ := strconv.Atoi(lineInts[0] + lineInts[len(lineInts)-1])
			sum += tmp
		}
	}

	fmt.Println(sum)
}

func Day1Sol2() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	var input []string
	for _, line := range lines {
		if line != "" {
			input = append(input, line)
		}
	}

	// numbers := "0123456789"

	// map[string]int{}
	numbersMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0

	for _, line := range input {
		m := make(map[int]int)
		for key := range numbersMap {
			fmt.Println("KEY:", key)
			index := getIndex(line, key)
			for _, i := range index {
				m[i] = numbersMap[key]
			}

		}
		for value := range numbersMap {
			if strings.Contains(line, value) {
				index := strings.Index(line, string(value))
				m[index] = numbersMap[value]
			}
		}
		fmt.Println("MAP PRINT:", m)
		// if len(m) == 1 {
		// 	tmp := m[0] + m[0]
		// 	sum += tmp
		// }
		// if len(m) == 2 {
		// 	keys := make([]int, 0, len(m))
		// 	for k := range m {
		// 		keys = append(keys, k)
		// 	}
		// 	sort.Ints(keys)
		// 	tmp := m[keys[0]] + m[keys[1]]
		// 	sum += tmp
		// }
		// if len(m) > 2 {
		// 	keys := make([]int, 0, len(m))
		// 	for k := range m {
		// 		keys = append(keys, k)
		// 	}
		// 	sort.Ints(keys)
		// 	tmp := m[keys[0]] + m[keys[len(keys)-1]]
		// 	sum += tmp
		// }
	}

	fmt.Println(sum)

}

func getIndex(line, key string) []int {
	var indexes []int
	for i, _ := range line {
		index := strings.Index(line[i:], key)
		if index == -1 {
			break
		}

		indexes = append(indexes, index)
	}
	return indexes
}
