package main

import (
	"math/rand"
)

type StepFn func(*Creature, *Ecosystem) bool

func AllSteps() []StepFn {
	return []StepFn{
		Photosynthesis,
		ConsumeCompetitor,
		Reproduce,
	}
}

/**
* Photosynthesis gives you +2 energy
 */
func Photosynthesis(creature *Creature, eco *Ecosystem) bool {
	creature.energy += 1
	return true
}

/**
* Search the ecosystem for the next competitor you can eat. Only
* competitors that have your size or smaller are edible. You gain +1 size
* and competitor's energy - 1 by eating them.
 */
func ConsumeCompetitor(creature *Creature, eco *Ecosystem) bool {
	for i, competitor := range eco.creatures {
		// selection strategy is sort of desperate. No appreciation for their food!
		if competitor != nil && competitor.size <= creature.size && competitor.name != creature.name {

			if creature.size != competitor.size || rand.Float64() < 0.5 {
				creature.size++
				creature.energy += competitor.energy + 1
				eco.KillCreature(i)
				return true
			} else {
				break // sorry, you lose
			}
		}
	}
	return false
}

/**
* Simple reproduction.
*
* Phase 1: partner selection. the creature will pick the
* partner from the ecosystem with the same or lesser size
* with the greatest amount of energy.
*
* Phase 2: if the ecosystem has space, create a creature with
* the amount of energy
 */
func Reproduce(creature *Creature, eco *Ecosystem) bool {
	retVal := false
	for i := 0; i < creature.energy; i++ {
		if eco.HasSpace() {
			offspring := &Creature{creature.name, 1, DEFAULT_ENERGY, 0, creature.program.Mutate(1)}
			eco.AddCreature(offspring)
			retVal = true
		}
	}
	return retVal
}
