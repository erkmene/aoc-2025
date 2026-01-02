package main

import (
	"aoc2025/utils"
	"log"
	"strconv"
	"strings"
)

func checkDoubleOccurrence(value int) bool {
	stringValue := strconv.Itoa(value)
	if len(stringValue)%2 == 0 {
		firstHalf := stringValue[:len(stringValue)/2]
		secondHalf := stringValue[len(stringValue)/2:]
		return firstHalf == secondHalf
	}
	return false
}

func checkMultipleOccurrence(value int) bool {
	stringValue := strconv.Itoa(value)
	// No need to check after the first half
	for charCount := 1; charCount <= len(stringValue)/2; charCount++ {
		toCheck := stringValue[:charCount]
		rest := stringValue[charCount:]
		if len(rest)%len(toCheck) == 0 && strings.Count(rest, toCheck) == len(rest)/len(toCheck) {
			return true
		}
	}
	return false
}

func getInvalidIds(node utils.Node, check func(int) bool) []int {
	invalidIds := make([]int, 0)

	// Ranges
	for _, rangeTuple := range node.Children {
		min, _ := strconv.Atoi(rangeTuple.Children[0].Value)
		max, _ := strconv.Atoi(rangeTuple.Children[1].Value)
		for i := min; i <= max; i++ {
			if check(i) {
				invalidIds = append(invalidIds, i)
			}
		}
	}

	return invalidIds
}

func sumInvalidIds(invalidIds []int) int {
	sum := 0
	for _, id := range invalidIds {
		sum += id
	}
	return sum
}

func main() {
	timer := utils.NewTimeRecord()

	testData := utils.ImportMultiDimensionalData("data/02.test.dat", []string{",", "-"})
	data := utils.ImportMultiDimensionalData("data/02.dat", []string{",", "-"})

	timer.LapAndLog("Data Parsing Done")

	testInvalidIds := getInvalidIds(testData, checkDoubleOccurrence)
	if len(testInvalidIds) != 8 {
		log.Fatalf("FAIL solution 1: expected 8, got %d", testInvalidIds)
	}
	testSum := sumInvalidIds(testInvalidIds)
	if testSum != 1227775554 {
		log.Fatalf("FAIL solution 1: expected 1227775554, got %d", testSum)
	}

	timer.LapAndLog("Test 1 Done")

	invalidIds := getInvalidIds(data, checkDoubleOccurrence)
	solution := sumInvalidIds(invalidIds)
	timer.LapAndLog("Solution 1 Done")

	log.Print("--------------------------------")
	log.Printf("solution 1: %d", solution)
	log.Print("--------------------------------")

	testInvalidIds = getInvalidIds(testData, checkMultipleOccurrence)
	if len(testInvalidIds) != 13 {
		log.Fatalf("FAIL solution 2: expected 13, got %d", testInvalidIds)
	}
	testSum = sumInvalidIds(testInvalidIds)
	if testSum != 4174379265 {
		log.Fatalf("FAIL solution 2: expected 4174379265, got %d", testSum)
	}

	timer.LapAndLog("Test 2 Done")

	invalidIds = getInvalidIds(data, checkMultipleOccurrence)
	solution = sumInvalidIds(invalidIds)

	log.Print("--------------------------------")
	log.Printf("solution 2: %d", solution)
	log.Print("--------------------------------")

	timer.LapAndLog("Solution 2 Done")
}
