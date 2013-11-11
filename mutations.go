package main

import (
	"math/rand"
)

type MutationFn func(DecisionProgram) DecisionProgram

func AllMutations() []MutationFn {
	return []MutationFn{
		swapStep,
		alterStepProbability,
		addStep,
		removeStep,
	}
}

/**
* This is a mutation that swaps the places of two steps in the program
 */
func swapStep(program DecisionProgram) DecisionProgram {
	if len(program.steps) > 1 {
		numSteps := len(program.steps)
		firstStepIndex := rand.Intn(numSteps)
		secondStepIndex := rand.Intn(numSteps-firstStepIndex) + firstStepIndex

		tmpStep := program.steps[firstStepIndex]
		program.steps[firstStepIndex] = program.steps[secondStepIndex]
		program.steps[secondStepIndex] = tmpStep
	}
	return program
}

/**
* Changes the probability of a step in the program to [0,1)
 */
func alterStepProbability(program DecisionProgram) DecisionProgram {
	if len(program.steps) > 0 {
		randStep := rand.Intn(len(program.steps))

		program.steps[randStep].probability = rand.Float64()
	}
	return program
}

/**
* Appends a new step to the end of the program
 */
func addStep(program DecisionProgram) DecisionProgram {

	allSteps := AllSteps()
	stepIndex := rand.Intn(len(allSteps))
	program.steps = append(program.steps, &ProgramStep{allSteps[stepIndex], rand.Float64()})

	return program
}

func removeStep(program DecisionProgram) DecisionProgram {
	if len(program.steps) > 0 {
		deletedStep := rand.Intn(len(program.steps))
		program.steps = append(program.steps[:deletedStep], program.steps[deletedStep+1:]...)

	}
	return program
}
