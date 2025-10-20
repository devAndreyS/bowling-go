package main

import (
    "fmt"
)

func calcScore(rolls []int) int {
    total := 0
    i := 0

    for frame := 1; frame <= 10; frame++ {
    	// Strike
    	if rolls[i] == 10 {
            total += 10 + rolls[i+1] + rolls[i+2]
            i++
            continue
    	}

    	frameScore := rolls[i] + rolls[i+1]

    	// Spare
    	if frameScore == 10 {
            total += 10 + rolls[i+2]
    	} else {
            total += frameScore
    	}

    	i += 2
    }

    return total
}

func main() {
    tests := [][]int{
    	{10, 7, 3, 9, 0, 10, 10, 2, 8, 6, 4, 2, 3, 5, 5, 10, 10, 10},
    	{5, 2, 5, 3, 5, 4, 5, 5, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
    	{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
    }

    for idx, rolls := range tests {
    	fmt.Printf("Game %d score: %d\n", idx+1, calcScore(rolls))
    }
}