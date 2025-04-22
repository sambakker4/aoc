package main

import (
	"slices"
	"math"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	reports := getReports("input.txt")
	safeReports := 0

	for _, report := range reports {
		safe := isReportSafe(report)
		if safe {
			safeReports++
		}
	}
	fmt.Println("Part 1:", safeReports)

	safeReports = 0

	for _, report := range reports {
		safe := isReportSafe(report)
		if safe {
			safeReports++
			continue
		}

		for i := range report {
			newReport := slices.Concat(report[:i], report[i + 1:])			
			safe = isReportSafe(newReport)
			if safe {
				safeReports++
				break
			}
		}

	}
	fmt.Println("Part 2:", safeReports)
}

func getReports(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	reports := make([][]int, 0)

	for _, str := range strings.Split(string(data), "\n") {
		if str == "" {
			continue
		}
		report := make([]int, 0)

		for _, char := range strings.Fields(str){
			num, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports
}

func isReportSafe(report []int) bool {
	ascending := false
	descending := false

	for i, num := range report {
		if i == 0 {
			continue
		}

		diff := report[i - 1] - num

		if math.Abs(float64(diff)) > 3 || math.Abs(float64(diff)) < 1{
			return false
		}

		if diff < 0{
			ascending = true	
		} else if diff > 0 {
			descending = true
		}

		if descending && ascending{
			return false
		}
	}
	return true
}
