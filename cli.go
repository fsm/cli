package cli

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/fsm/fsm"
)

// Start begins the CLI target for a fsm.StateMachine
func Start(stateMachine fsm.StateMachine, store fsm.Store) {
	// Create Emitter
	emitter := &emitter{}

	// Get Traverser
	uuid := uuid()
	newTraverser := false
	traverser, err := store.FetchTraverser(uuid)
	if err != nil {
		traverser, _ = store.CreateTraverser(uuid)
		traverser.SetCurrentState("start")
		newTraverser = true
	}

	// Get Start State
	currentState := stateMachine[traverser.CurrentState()](emitter, traverser)
	if newTraverser {
		currentState.EntryAction()
	}

	// Prep Reader
	reader := bufio.NewReader(os.Stdin)
	for currentState != nil {
		// Read Input from User
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		// Pass Input to State
		currentState = stateMachine[traverser.CurrentState()](emitter, traverser)
		newState := currentState.Transition(text)
		newState.EntryAction()
		currentState = newState
		traverser.SetCurrentState(currentState.Slug)
	}
}

// uuid generates a UUID that isn't particularly compliant
// to any spec but is more than good enough for our use case
func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	str := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return str
}
