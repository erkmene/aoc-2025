package main

import (
	"aoc2025/utils"
	"log"
	"math"
	"strconv"
)

func parseData(data []string) []int {
	steps := make([]int, 0, len(data))
	for _, line := range data {
		isLeft := line[0] == 'L'
		number, err := strconv.Atoi(line[1:])
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
	testData := parseData(utils.ImportNewlineSeparatedData("01.test.dat"))
	testSteps := execute(testData)
	data := parseData(utils.ImportNewlineSeparatedData("01.dat"))
	steps := execute(data)

	testZeroes := countZeroes(testSteps)
	if testZeroes != 3 {
		log.Fatalf("FAIL solution 1: expected 3, got %d", testZeroes)
	}

	solutionOne := countZeroes(steps)
	log.Printf("solution 1: %d", solutionOne)

	testZeroClicks := countZeroClicks(testSteps)
	if testZeroClicks != 6 {
		log.Fatalf("FAIL solution 2: expected 6, got %d", testZeroClicks)
	}

	solutionTwo := countZeroClicks(steps)
	log.Printf("solution 2: %d", solutionTwo)
}
