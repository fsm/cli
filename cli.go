package cli

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/fsm/fsm"
	targetutil "github.com/fsm/target-util"
)

// Start begins the CLI target for a fsm.StateMachine
func Start(stateMachine fsm.StateMachine, store fsm.Store) {
	uuid := uuid()
	stateMap := targetutil.GetStateMap(stateMachine)
	reader := bufio.NewReader(os.Stdin)
	emitter := &emitter{}

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		// Step
		targetutil.Step(uuid, text, store, emitter, stateMap)
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
