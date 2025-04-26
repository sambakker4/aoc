package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

type point struct {
	i, j int
}

func main() {
	maze := getMaze("input.txt")
	i, j := getStartPos(maze)
	visited := []point{point{i: i, j: j}}
	up := true
	down := false
	right := false
	left := false

	for {
		if up && i-1 < 0 {
			break
		}
		if down && i+1 == len(maze) {
			break
		}
		if left && j-1 < 0 {
			break
		}
		if right && j+1 == len(maze[0]) {
			break
		}

		if up {
			if string(maze[i-1][j]) == "#" {
				right = true
				up = false
				continue
			} else {
				i--
				if slices.Index(visited, point{i: i, j: j}) == -1 {
					visited = append(visited, point{i: i, j: j})
				}
			}
		}

		if down {
			if string(maze[i+1][j]) == "#" {
				left = true
				down = false
				continue
			} else {
				i++
				if slices.Index(visited, point{i: i, j: j}) == -1 {
					visited = append(visited, point{i: i, j: j})
				}
			}
		}

		if right {
			if string(maze[i][j+1]) == "#" {
				right = false
				down = true
				continue
			} else {
				j++
				if slices.Index(visited, point{i: i, j: j}) == -1 {
					visited = append(visited, point{i: i, j: j})
				}
			}
		}

		if left {
			if string(maze[i][j-1]) == "#" {
				up = true
				left = false
				continue
			} else {
				j--
				if slices.Index(visited, point{i: i, j: j}) == -1 {
					visited = append(visited, point{i: i, j: j})
				}
			}
		}
	}

	fmt.Println("Part 1:", len(visited))
	total := 0
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if string(maze[i][j]) == "#" || string(maze[i][j]) == "^"{
				continue
			}
			str := maze[i]
			bytes := []byte(str)
			bytes[j] = '#'
			str = string(bytes)
			maze[i] = str
			if isGuardStuckInALoop(maze) {
				total++
			}
			str = maze[i]
			bytes = []byte(str)
			bytes[j] = '.'
			str = string(bytes)
			maze[i] = str
		}
	}
	fmt.Println("Part 2:", total)
}

func isGuardStuckInALoop(maze []string) bool {
	i, j := getStartPos(maze)
	type visitedDirection struct {
		i, j      int
		direction string
	}

	visited := make(map[string]bool)
	direction := "up"

	for {
		if direction == "up" && i-1 < 0 {
			break
		}
		if direction == "down" && i+1 == len(maze) {
			break
		}
		if direction == "left" && j-1 < 0 {
			break
		}
		if direction == "right" && j+1 == len(maze[0]) {
			break
		}

		switch direction {
			case "up":
				if string(maze[i - 1][j]) == "#" {
					direction = "right"
				} else {
					i--	
				}
				
				if visited[fmt.Sprintf("%d,%d,%s", i, j, direction)]{
					return true
				}
				visited[fmt.Sprintf("%d,%d,%s", i, j, direction)] = true

			case "down":
				if string(maze[i + 1][j]) == "#" {
					direction = "left"
				} else {
					i++
				}
				
				if visited[fmt.Sprintf("%d,%d,%s", i, j, direction)]{
					return true
				}
				visited[fmt.Sprintf("%d,%d,%s", i, j, direction)] = true

			case "right":
				if string(maze[i][j + 1]) == "#" {
					direction = "down"
				} else {
					j++
				}
				
				if visited[fmt.Sprintf("%d,%d,%s", i, j, direction)]{
					return true
				}
				visited[fmt.Sprintf("%d,%d,%s", i, j, direction)] = true

			case "left":
				if string(maze[i][j - 1]) == "#" {
					direction = "up"
				} else {
					j--
				}
				
				if visited[fmt.Sprintf("%d,%d,%s", i, j, direction)]{
					return true
				}
				visited[fmt.Sprintf("%d,%d,%s", i, j, direction)] = true
		}
	}
	return false 
}


func getStartPos(maze []string) (int, int) {
	for i, line := range maze {
		for j := range line {
			if string(maze[i][j]) == "^" {
				return i, j
			}
		}
	}
	return 0, 0
}

func getMaze(filePath string) []string {
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
