package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var wordSearch []string
	for scanner.Scan() {
		wordSearch = append(wordSearch, scanner.Text())
	}

	part1(wordSearch)
	part2(wordSearch)
}

func part1(wordSearch []string) {
	totalChristmas := 0
	for yIndex, line := range wordSearch {
		for xIndex, r := range line {
			if r == 'X' {
				totalChristmas += findChristmas(xIndex, yIndex, wordSearch)
			}
		}
	}

	fmt.Println(totalChristmas)
}

func findChristmas(x int, y int, wordSearch []string) int {
	totalChristmas := 0

	// check horizontal
	currentLine := wordSearch[y]

	// backwards
	if x > 2 {
		if currentLine[x-1] == 'M' && currentLine[x-2] == 'A' && currentLine[x-3] == 'S' {
			totalChristmas++
		}
	}

	// forwards
	if (x + 3) < len(currentLine) {
		if currentLine[x+1] == 'M' && currentLine[x+2] == 'A' && currentLine[x+3] == 'S' {
			totalChristmas++
		}
	}

	// check vertical

	// up
	if y > 2 {
		mLine := wordSearch[y-1]
		aLine := wordSearch[y-2]
		sLine := wordSearch[y-3]

		if mLine[x] == 'M' && aLine[x] == 'A' && sLine[x] == 'S' {
			totalChristmas++
		}
	}

	// down
	if (y + 3) < len(wordSearch) {
		mLine := wordSearch[y+1]
		aLine := wordSearch[y+2]
		sLine := wordSearch[y+3]

		if mLine[x] == 'M' && aLine[x] == 'A' && sLine[x] == 'S' {
			totalChristmas++
		}
	}

	if y > 2 {
		mLine := wordSearch[y-1]
		aLine := wordSearch[y-2]
		sLine := wordSearch[y-3]

		// diagonal up, forward
		if (x + 3) < len(currentLine) {
			if mLine[x+1] == 'M' && aLine[x+2] == 'A' && sLine[x+3] == 'S' {
				totalChristmas++
			}
		}

		// diagonal up, backward
		if x > 2 {
			if mLine[x-1] == 'M' && aLine[x-2] == 'A' && sLine[x-3] == 'S' {
				totalChristmas++
			}
		}
	}

	// diagonal down
	if (y + 3) < len(wordSearch) {
		mLine := wordSearch[y+1]
		aLine := wordSearch[y+2]
		sLine := wordSearch[y+3]

		// diagonal down, forward
		if (x + 3) < len(currentLine) {
			if mLine[x+1] == 'M' && aLine[x+2] == 'A' && sLine[x+3] == 'S' {
				totalChristmas++
			}
		}

		// diagonal down, backward
		if x > 2 {
			if mLine[x-1] == 'M' && aLine[x-2] == 'A' && sLine[x-3] == 'S' {
				totalChristmas++
			}
		}

	}

	return totalChristmas
}

func part2(wordSearch []string) {
	totalChristmas := 0
	for yIndex, line := range wordSearch {
		for xIndex, r := range line {
			if r == 'M' {
				totalChristmas += findChristmas2(xIndex, yIndex, wordSearch)
			}
		}
	}

	fmt.Println(totalChristmas)
}

func findChristmas2(x int, y int, wordSearch []string) int {
	totalChristmas := 0

	// check horizontal
	currentLine := wordSearch[y]

	if x+2 < len(currentLine) && y+2 < len(wordSearch) {
		aLine := wordSearch[y+1]
		sLine := wordSearch[y+2]

		if aLine[x+1] == 'A' && sLine[x+2] == 'S' {
			if currentLine[x+2] == 'M' && sLine[x] == 'S' {
				totalChristmas++
			}

			if currentLine[x+2] == 'S' && sLine[x] == 'M' {
				totalChristmas++
			}
		}
	}

	if x > 1 && y > 1 {
		aLine := wordSearch[y-1]
		sLine := wordSearch[y-2]

		if aLine[x-1] == 'A' && sLine[x-2] == 'S' {
			if currentLine[x-2] == 'M' && sLine[x] == 'S' {
				totalChristmas++
			}

			if currentLine[x-2] == 'S' && sLine[x] == 'M' {
				totalChristmas++
			}
		}
	}

	return totalChristmas
}
