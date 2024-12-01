package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func b() {
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
	totalSimilarity := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, SEPERATOR)
		_left, _ := strconv.Atoi(splitLine[0])
		_right, _ := strconv.Atoi(splitLine[1])
		left = append(left, _left)
		right = append(right, _right)
	}

	for i := 0; i < len(left); i++ {
		found := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				found++
			}
		}
		totalSimilarity += left[i] * found
	}

	log.Println("b", totalSimilarity)

}
