package main

import (
	"io"
	"log"
	"os"
	"strings"	
	"fmt"
)

func main() {
	search := getWordSearch("input.txt")
	total := 0
	word := "XMAS"

	for i := 0; i < len(search); i++ {
		for j := 0; j < len(search[i]); j++ {
			// up
			if i - 3 >= 0 {
				if string(search[i][j]) + string(search[i - 1][j]) + string(search[i - 2][j]) + string(search[i - 3][j]) == word {
					total++
				}
			}
			// down
			if i + 3 < len(search) {
				if string(search[i][j]) + string(search[i + 1][j]) + string(search[i + 2][j]) + string(search[i + 3][j]) == word {
					total++
				}
			}

			// right
			if j + 3 < len(search[0]) {
				if string(search[i][j]) + string(search[i][j + 1]) + string(search[i][j + 2]) + string(search[i][j + 3]) == word {
					total++
				}
			}
			// left
			if j - 3 >= 0 {
				if string(search[i][j]) + string(search[i][j - 1]) + string(search[i][j - 2]) + string(search[i][j - 3]) == word {
					total++
				}
			}
			
			// up right
			if i - 3 >= 0 && j + 3 < len(search[0]){
				if (
					string(search[i][j]) + 
					string(search[i - 1][j + 1]) + 
					string(search[i - 2][j + 2]) + 
					string(search[i - 3][j + 3]) == word) {
						total++
				}
			}
			//down right
			if i + 3 < len(search) && j + 3 < len(search[0]){
				if (
					string(search[i][j]) + 
					string(search[i + 1][j + 1]) + 
					string(search[i + 2][j + 2]) + 
					string(search[i + 3][j + 3]) == word) {
						total++
				}
			}

			// up left
			if i - 3 >= 0 && j - 3 >= 0{
				if (
					string(search[i][j]) + 
					string(search[i - 1][j - 1]) + 
					string(search[i - 2][j - 2]) + 
					string(search[i - 3][j - 3]) == word) {
						total++
				}
			}
			// down left
			if i + 3 < len(search) && j - 3 >= 0{
				if (
					string(search[i][j]) + 
					string(search[i + 1][j - 1]) + 
					string(search[i + 2][j - 2]) + 
					string(search[i + 3][j - 3]) == word) {
						total++
				}
			}
		}
	}

	fmt.Println("Part 1:", total)

	total = 0

	for i := 0; i < len(search); i++ {
		for j := 0; j < len(search[i]); j++ {
			if string(search[i][j]) != "A" {
				continue
			}

			if !((i - 1 >= 0) && (i + 1 < len(search)) && (j + 1 < len(search[0])) && (j - 1 >= 0)) {
				continue
			}

			if (
				string(search[i - 1][j + 1]) == "M" &&
				string(search[i - 1][j - 1]) == "M" &&
				string(search[i + 1][j + 1]) == "S" &&
				string(search[i + 1][j - 1]) == "S"){
				total++
				continue
			}

			if (
				string(search[i - 1][j + 1]) == "S" &&
				string(search[i - 1][j - 1]) == "M" &&
				string(search[i + 1][j + 1]) == "S" &&
				string(search[i + 1][j - 1]) == "M"){
				total++
				continue
			}
			if (
				string(search[i - 1][j + 1]) == "M" &&
				string(search[i - 1][j - 1]) == "S" &&
				string(search[i + 1][j + 1]) == "M" &&
				string(search[i + 1][j - 1]) == "S"){
				total++
				continue
			}
			if (
				string(search[i - 1][j + 1]) == "S" &&
				string(search[i - 1][j - 1]) == "S" &&
				string(search[i + 1][j + 1]) == "M" &&
				string(search[i + 1][j - 1]) == "M"){
				total++
				continue
			}
		}
	}

	fmt.Println("Part 2:", total)
}

func getWordSearch(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}
