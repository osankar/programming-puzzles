package main

import (
	"fmt"
)

func main() {
	caps := []string{"B","F", "B", "B", "F", "F", "F", "F", "B", "F", "B", "B", "F", "B", "B", "B", "F", "F", "F", "F","B"}
	giveCommands(caps)
}

func giveCommands(caps []string) []*Range {
	caps = append(caps, caps[0])
	ranges := make([]*Range, 0)
	for index := 1; index < len(caps); index++ {
		if caps[index] == caps[index-1] {
			continue
		}
		if caps[index] != caps[0] {
			fmt.Printf("People with tokens %d ", index)
		} else {
			fmt.Printf(" to %d, flip your caps ", index-1)
			fmt.Println()
		}
	}
	return ranges
}

type Range struct {
	start int
	end   int
}

func (r *Range) String() string {
	return fmt.Sprintf("[%d-%d]", r.start, r.end)
}