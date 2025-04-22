package main

import(
	"os"
	"log"
	"io"
	"strings"
	"strconv"
	"fmt"
	"sort"
	"math"
)

func main() {
	list1, list2 := getListsFromFile("input.txt")
	total := 0
	sort.Slice(list1, func(i, j int) bool {return list1[i] < list1[j]})
	sort.Slice(list2, func(i, j int) bool {return list2[i] < list2[j]})

	for i, num := range list1 {
		total += int(math.Abs(float64(num) - float64(list2[i])))
	}
	fmt.Println("Part 1:", total)

	similarityScores := 0

	for _, num := range list1 {
		similarityScores +=  num * numOfTimesInList(num, list2)
	}
	fmt.Println("Part 2:", similarityScores)
}

func numOfTimesInList(target int, list []int) int {
	total := 0
	for _, num := range list {
		if num == target {
			total += 1
		}
	}
	return total
}


func getListsFromFile(filePath string) ([]int, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lists := strings.Split(string(data), "\n")

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, item := range lists {
		if item == ""{
			break
		}
		item1 := strings.Fields(item)[0]
		num1, err := strconv.ParseInt(item1, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, int(num1))

		item2 := strings.Fields(item)[1]
		num2, err := strconv.ParseInt(item2, 10, 64)
		if err != nil {
			log.Fatal(err)
		} 
		list2 = append(list2, int(num2))
	}
	return list1, list2
}
