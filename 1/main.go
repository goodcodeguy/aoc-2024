package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left []int
	var right []int

	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		leftInt, err := strconv.ParseInt(words[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rightInt, err := strconv.ParseInt(words[1], 10, 64)
		if err != nil {
			panic(err)
		}

		left = append(left, int(leftInt))
		right = append(right, int(rightInt))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	part1(left, right)
	part2(left, right)
}

func part1(left []int, right []int) {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i, leftItem := range left {
		totalDistance = totalDistance + int(math.Abs(float64(leftItem-right[i])))
	}

	fmt.Println(totalDistance)
}

func part2(left []int, right []int) {
	simScore := 0
	for _, leftItem := range left {
		m := 0

		for _, rightItem := range right {
			if leftItem == rightItem {
				m = m + 1
			}
		}
		simScore = simScore + (leftItem * m)
	}

	fmt.Println(simScore)
}
