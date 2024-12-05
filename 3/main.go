package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	combinedString := ""
	for scanner.Scan() {
		combinedString += scanner.Text()
	}

	fmt.Println(part1(combinedString))
	fmt.Println(part2(combinedString))
}

func part1(input string) int {
	total := 0
	r, _ := regexp.Compile(`mul\((?P<a>\d*),(?P<b>\d*)\)`)

	for _, m := range r.FindAllStringSubmatch(input, -1) {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		total += a * b
	}

	return total
}

func part2(input string) int {
	total := 0
	enabled := true
	r, _ := regexp.Compile(`mul\((?P<a>\d*),(?P<b>\d*)\)|do\(\)|don't\(\)`)

	for _, m := range r.FindAllStringSubmatch(input, -1) {

		if enabled && strings.Contains(m[0], "mul") {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			total += a * b
		}

		if strings.Contains(m[0], "do()") {
			enabled = true
		}

		if strings.Contains(m[0], "don't()") {
			enabled = false
		}

	}

	return total
}
