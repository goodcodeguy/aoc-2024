package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const SafetyThreshold = 3

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var reports [][]int
	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		reportData := make([]int, len(levels))

		for i, s := range levels {
			reportData[i], _ = strconv.Atoi(s)
		}

		reports = append(reports, reportData)
	}

	part1(reports)
	part2(reports)

}

func part1(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report, false) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func part2(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report, true) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isReportSafe(report []int, withDampening bool) bool {
	isSafe := true
	d := 0

	if report[0] > report[1] {
		d = -1
	} else {
		d = 1
	}

	for i, level := range report {
		if i >= (len(report) - 1) {
			break
		}

		if math.Abs(float64(level)-float64(report[i+1])) > SafetyThreshold {
			isSafe = false
			break
		}

		if level == report[i+1] {
			isSafe = false
			break
		}

		if level > report[i+1] {
			if d == 1 {
				isSafe = false
				break
			}
		} else {
			if d == -1 {
				isSafe = false
				break
			}
		}
	}

	if withDampening && !isSafe {

		// Brute Force any unsafe Reports
		for rIndex := -1; rIndex <= len(report); rIndex++ {
			newReport := make([]int, 0, len(report)-1)
			for i, level := range report {
				if i == rIndex {
					continue
				}
				newReport = append(newReport, level)
			}

			if isReportSafe(newReport, false) {
				return true
			}
		}
	}

	return isSafe
}
