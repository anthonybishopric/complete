package main

import (
	"fmt"
)

const (
	DEFAULT_ENERGY = 2
)

func main() {
	eco := &Ecosystem{}
	eco.Init()
	eco.Debug()
	var cmd string
	var val int
	for true {
		fmt.Printf("'t N' run N turns/ 'x 0' terminate / 'c ID' debug creature with ID\n")
		fmt.Scanf("%s %d", &cmd, &val)

		if cmd == "t" {
			for x := 0; x < val; x++ {
				eco.ExecuteTurn()
			}
			eco.Debug()
		} else if cmd == "c" {
			if val < 0 || val >= len(eco.creatures) {
				fmt.Printf("Invalid index %d", val)
			} else {
				creature := eco.creatures[val]
				fmt.Printf("\n%s\n\n%s\n", creature.JSON(), creature.program.JSON())
			}
		} else if cmd == "x" {
			return
		} else {
			fmt.Printf("Unrecognized command %s\n", cmd)
		}
	}
}
