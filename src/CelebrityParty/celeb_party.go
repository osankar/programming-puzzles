package main

import (
	"fmt"
	"sort"
)

func main() {
	scheds := [][]float32{{6, 8}, {6, 12}, {6, 7}, {7, 8}, {7, 10}, {8, 9}, {8, 10}, {9, 12},
		{9, 10}, {10, 11}, {10, 12}, {11, 12}}
	scheds2 := [][]float32{{6.0, 8.0}, {6.5, 12.0}, {6.5, 7.0}, {7.0, 8.0}, {7.5, 10.0}, {8.0, 9.0},
		{8.0, 10.0}, {9.0, 12.0}, {9.5, 10.0}, {10.0, 11.0}, {10.0, 12.0}, {11.0, 12.0}}
	scheds3 := [][]float32{{6, 7}, {7, 9}, {10, 11}, {10, 12}, {8, 10}, {9, 11}, {6, 8},
		{9, 10}, {11, 12}, {11, 13}, {11, 14}}
	//scheds:= [][]float32{{6.5, 8.0}, {5.0, 9.3}}
	optimalTime, noOfCelebs := findOptimalPartyTime(scheds)
	fmt.Printf("You can party with %d celebs at %f \n", noOfCelebs, optimalTime)
	optimalTime, noOfCelebs = findOptimalPartyTime(scheds2)
	fmt.Printf("You can party with %d celebs at %f \n", noOfCelebs, optimalTime)
	optimalTime, noOfCelebs = findOptimalPartyTime(scheds3)
	fmt.Printf("You can party with %d celebs at %f \n", noOfCelebs, optimalTime)

}

func findOptimalPartyTime(scheds [][]float32) (float32, int) {
	movements := buildMovements(scheds)
	sort.Slice(movements, compareMovements(movements))
	fmt.Println(movements)
	maxCelebs := 0
	var optimalTime float32
	currentCelebs := 0
	for _, movement := range movements {
		if movement.typ == "enter" {
			currentCelebs++
			if maxCelebs < currentCelebs {
				maxCelebs = currentCelebs
				optimalTime = movement.time
			}
		} else {
			currentCelebs--
		}
	}
	return optimalTime, maxCelebs
}

func compareMovements(movements []move) func(i int, j int) bool {
	return func(i, j int) bool {
		//if both celeb movements are happening at the same time exit gets more priority so that the
		//exists at that point of time comes before the entrees so that the count reflects the exact no of celebs inside
		if movements[i].time == movements[j].time {
			return movements[i].typ > movements[j].typ
		}
		return movements[i].time < movements[j].time
	}
}

func buildMovements(scheds [][]float32) []move {
	movements := make([]move, 0)
	for _, sched := range scheds {
		movements = append(movements, move{time: sched[0], typ: "enter"})
		movements = append(movements, move{time: sched[1], typ: "exit"})
	}
	return movements
}

type move struct {
	time float32
	typ  string
}
