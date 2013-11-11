package main

import (
	"errors"
	"fmt"
	"sort"
)

type Ecosystem struct {
	creatures       [100]*Creature
	spacesLeft      int
	emittedEnergy   int // TBD - unused but could provide a cap on life created
	turn            int
	nameChannel     chan string
	termNameChannel chan bool
}

// implement sort.Interface methods

func (e *Ecosystem) Len() int {
	return len(e.creatures)
}

func (e *Ecosystem) Less(i, j int) bool {
	if e.creatures[i] == nil {
		return false
	}
	return e.creatures[i].size > e.creatures[j].size
}

func (e *Ecosystem) Swap(i, j int) {
	tmpCreature := e.creatures[i]
	e.creatures[i] = e.creatures[j]
	e.creatures[j] = tmpCreature
}

func (e *Ecosystem) KillCreature(ordinal int) {
	e.creatures[ordinal] = nil
	e.spacesLeft++
}

func (e *Ecosystem) HasSpace() bool {
	return e.spacesLeft > 0
}

func (e *Ecosystem) AddCreature(creature *Creature) (int, error) {
	if !e.HasSpace() {
		return -1, errors.New("Cannot add new creature; no space left")
	}
	for i := 0; i < len(e.creatures); i++ {
		if e.creatures[i] == nil {
			e.creatures[i] = creature
			e.spacesLeft--
			return i, nil
		}
	}
	return -1, errors.New("Eco thought it had space but no creature was nil!")
}

func (e *Ecosystem) Init() {
	e.nameChannel = make(chan string)
	e.termNameChannel = make(chan bool)
	go CreatureNames(e.nameChannel, e.termNameChannel)
	e.FillEcosystem()
}

func (e *Ecosystem) ExecuteTurn() {
	sort.Sort(e) // bigger creatures get to go first
	for i, creature := range e.creatures {
		if creature != nil {
			creature.program.Execute(creature, e)
			creature.age++
			creature.energy = creature.energy - creature.size
			if creature.energy <= 0 {
				e.KillCreature(i)
			}
		}
	}
	e.FillEcosystem()
	e.turn++
}

func (e *Ecosystem) Debug() {
	fmt.Printf("on turn %d...\n", e.turn)
	for i, creature := range e.creatures {
		fmt.Printf("%d: %s\n", i, creature.JSON())
	}
}

/**
* Fill the ecoystem with new creatures. Returns the
* number of creatures successfully created
 */
func (e *Ecosystem) FillEcosystem() int {
	addedCreatures := 0
	for i := 0; i < len(e.creatures); i++ {
		if e.creatures[i] == nil {
			program := &DecisionProgram{}
			creature := Creature{<-e.nameChannel, 1, DEFAULT_ENERGY, 0, program.Mutate(5)}
			e.creatures[i] = &creature
			addedCreatures++
		}
	}
	e.spacesLeft = 0
	fmt.Printf("Filled void with %d creatures\n", addedCreatures)
	return addedCreatures
}
