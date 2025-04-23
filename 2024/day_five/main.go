package main

import (
	"strconv"
	"io"
	"log"
	"os"
	"strings"
	"fmt"
	"slices"
)

func main() {
	rules, updates := getRulesAndUpdates("input.txt")
	okUpdates := make([][]int, 0)
	notOkUpdates := make([][]int, 0)
	total := 0

	for _, update := range updates {
		bad := false
		for j, num := range update {
			rule := rules[num]	
			for _, r := range rule {
				if  slices.Index(update[:j], r) != -1 {
					bad = true
				}
			}
		}
		if bad {
			notOkUpdates = append(notOkUpdates, update)
			continue
		}
		okUpdates = append(okUpdates, update)
	}

	for _, update := range okUpdates {
		total += update[len(update) / 2]
	}
	fmt.Println("Part 1:", total)

	total = 0
	
	for _, update := range notOkUpdates {
		for isUpdateOk(update, rules){
			for i, num := range update {
				rule := rules[num]	
				for _, r := range rule {
					if  idx := slices.Index(update[:i], r); idx != -1 {
						tmp := update[i]
						update[i] = update[idx]
						update[idx] = tmp
					}
				}
			}
		}
		total += update[len(update) / 2]
	}

	fmt.Println("Part 2:", total)
}

func isUpdateOk(update []int, rules map[int][]int) bool {
	bad := false
	for j, num := range update {
		rule := rules[num]	
		for _, r := range rule {
			if  slices.Index(update[:j], r) != -1 {
				bad = true
			}
		}
	}
	return bad
}

func getRulesAndUpdates(filePath string) (map[int][]int, [][]int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	rules := make(map[int][]int)
	updates := make([][]int, 0)
	updateStart := 0

	for i, line := range lines {
		if line == "" {
			updateStart = i + 1
			break
		}
		rule := strings.Split(line, "|")

		num1, _ := strconv.Atoi(rule[0])
		num2, _ := strconv.Atoi(rule[1])
		rules[num1] = append(rules[num1], num2)
	}

	for i := updateStart; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		updateInt := make([]int, 0)
		update := strings.Split(lines[i], ",")

		for _, u := range update {
			num, _ := strconv.Atoi(u)	
			updateInt = append(updateInt, num)
		}
		updates = append(updates, updateInt)
	}

	return rules, updates
}
