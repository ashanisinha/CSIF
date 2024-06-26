package nfa
import (
//	"fmt"
	"runtime"
)
// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	if len(input) == 0 {
		return start == final
	}

	ch := make (chan bool, 10)

	for _, nextState := range transitions(start, input[0]) {
	
			//if Reachable(transitions, nextState, final, input[1:]) { // put reachable in another goroutine 
				ch <- Reachable(transitions, nextState, final, input[1:])
			//}
			runtime.NumGoroutine()
	}
	close(ch)

	for x := range ch {
		if x {
			return true
		}
	}
	return false
}

