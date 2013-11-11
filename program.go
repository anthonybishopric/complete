package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
)

/**
* Each turn, a creature executes their decision program. Decision
* programs are immutable and are copied between generations of creatures.
* Between generations, random mutations may occur. Reproduction is itself
* an operation that is defined by the decision program. All steps in the program
* are probabilistic, ie each step in the program (if it is possible to execute) will
* execute with that probability. Only one step may execute in total.
*
* Example decision program:
*
* Step1: Consume smaller competitor, 0.5
* Step2: Reproduce, 0.3
* Step3: Rest, 1
*
* This program executes as follows: in step 1, with 50% likelihood, decide to eat a smaller
* competitor. If that doesn't occur, attempt to reproduce with 30% likelihood. If not, simply
* rest.
*
* TBD - make programatically defined steps (instead of provided functions)
 */
type DecisionProgram struct {
	steps []*ProgramStep
}

func (d *DecisionProgram) Execute(creature *Creature, eco *Ecosystem) {
	totalActions := creature.size
	for i := 0; i < len(d.steps); i++ {
		if d.steps[i].Execute(creature, eco) {
			totalActions--
			if totalActions == 0 {
				return
			}
		}
	}
}

func (d *DecisionProgram) JSON() string {
	programStepStrings := make([]string, len(d.steps))
	for i := 0; i < len(d.steps); i++ {
		programStepStrings[i] = d.steps[i].JSON()
	}

	return fmt.Sprintf("{ %s,\n }", strings.Join(programStepStrings, ",\n"))
}

type ProgramStep struct {
	logic       StepFn
	probability float64
}

func (p *ProgramStep) Execute(creature *Creature, eco *Ecosystem) bool {
	if p.probability > rand.Float64() {
		return p.logic(creature, eco)
	} else {
		return false
	}
}

func (p *ProgramStep) JSON() string {
	stepName := runtime.FuncForPC(reflect.ValueOf(p.logic).Pointer()).Name()
	return fmt.Sprintf("{'name': %s,\t'probability': %d}", stepName, p.probability)
}

/**
* Create a new program given a number of mutations on the given program.
* TBD - how should the volume of mutations be determined? Right now just a
* simple number.
 */
func (self *DecisionProgram) Mutate(mutations int) *DecisionProgram {
	newProgram := DecisionProgram{
		self.steps,
	}
	mutationFuncs := AllMutations()
	for i := 0; i < mutations; i++ {
		mutation := mutationFuncs[rand.Intn(len(mutationFuncs))]
		newProgram = mutation(newProgram)
	}
	return &newProgram
}
