package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := getInput("input.txt")
	total := 0
	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			total += num1 * num2
		}	
	}
	fmt.Println("Part 1:", total)
	total = 0
	re, err = regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatal(err)
	}
	doing := true 

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				doing = true
				continue
			} else if match[0] == "don't()" {
				doing = false
				continue
			}
			
			if doing {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				total += num1 * num2
			}
		}
	}
	
	fmt.Println("Part 2:", total)
}

func getInput(filePath string) []string{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
