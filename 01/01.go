package main

import (
	"aoc2025/utils"
	"log"
	"math"
	"strconv"
)

func parseData(data utils.Node) []int {
	steps := make([]int, 0, len(data.Children))
	for _, child := range data.Children {
		isLeft := child.Value[0] == 'L'
		number, err := strconv.Atoi(child.Value[1:])
		if err != nil {
			log.Fatal(err)
		}
		if isLeft {
			steps = append(steps, -number)
		} else {
			steps = append(steps, number)
		}
	}
	return steps
}

func getSign(value int) int {
	if value == 0 {
		return 0
	}
	if value > 0 {
		return 1
	}
	return -1
}

type Step struct {
	Position     int
	Change       int
	PassedZero   bool
	FullRotation int
}

func execute(data []int) []Step {
	position := 50

	results := make([]Step, 0, len(data)+1)
	results = append(results, Step{Position: position, Change: 0, PassedZero: false, FullRotation: 0})
	lastPosition := position

	for _, step := range data {
		passedZero := false

		fullRotation := int(math.Floor(math.Abs(float64(step) / 100)))
		step = step % 100
		position = position + step

		if lastPosition != 0 && (position < 0 || position > 100) {
			passedZero = true
		}
		if position < 0 {
			position = 100 + position
		}
		position = position % 100
		lastPosition = position

		results = append(results,
			Step{
				Position:     position,
				Change:       step,
				PassedZero:   passedZero,
				FullRotation: fullRotation})
	}
	return results
}

func countZeroes(data []Step) int {
	count := 0
	for _, step := range data {
		if step.Position%100 == 0 {
			count++
		}
	}
	return count
}

func countZeroClicks(data []Step) int {
	count := 0

	for index, step := range data {
		if index == 0 {
			continue
		}
		count += int(step.FullRotation)
		if step.PassedZero || step.Position == 0 {
			count += 1
		}
	}

	return count
}

func main() {
	timer := utils.NewTimeRecord()

	testData := parseData(utils.ImportMultiDimensionalData("data/01.test.dat", []string{"\n"}))
	testSteps := execute(testData)
	data := parseData(utils.ImportMultiDimensionalData("data/01.dat", []string{"\n"}))
	steps := execute(data)

	timer.LapAndLog("Data Parsing Done")

	testZeroes := countZeroes(testSteps)
	if testZeroes != 3 {
		log.Fatalf("FAIL solution 1: expected 3, got %d", testZeroes)
	}

	timer.LapAndLog("Test 1 Done")

	solutionOne := countZeroes(steps)
	log.Print("--------------------------------")
	log.Printf("solution 1: %d", solutionOne)
	log.Print("--------------------------------")

	timer.LapAndLog("Solution 1 Done")

	testZeroClicks := countZeroClicks(testSteps)
	if testZeroClicks != 6 {
		log.Fatalf("FAIL solution 2: expected 6, got %d", testZeroClicks)
	}

	timer.LapAndLog("Test 2 Done")

	solutionTwo := countZeroClicks(steps)
	log.Print("--------------------------------")
	log.Printf("solution 2: %d", solutionTwo)
	log.Print("--------------------------------")

	timer.LapAndLog("Solution 2 Done")
}
