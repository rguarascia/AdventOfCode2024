package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func a() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}
	SEPERATOR := "   "
	totalDifference := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, SEPERATOR)
		_left, _ := strconv.Atoi(splitLine[0])
		_right, _ := strconv.Atoi(splitLine[1])
		left = append(left, _left)
		right = append(right, _right)
	}

	// We can use left as the index, only because same length
	for i := 0; i < len(left); i++ {
		for j := i + 1; j < len(left); j++ {
			// Sort left
			if left[i] > left[j] {
				temp := left[i]
				left[i] = left[j]
				left[j] = temp
			}
			// Sort right
			if right[i] > right[j] {
				temp := right[i]
				right[i] = right[j]
				right[j] = temp
			}
		}
	}

	for i := 0; i < len(left); i++ {
		totalDifference += abs(left[i] - right[i])
	}

	log.Println("a", totalDifference)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
