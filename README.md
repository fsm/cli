# CLI

CLI is a command line target for [fsm](https://github.com/fsm/fsm).

This is an ideal target for conversational interface development.

## Getting Started

```go
package main

import (
	"github.com/fsm/cli"
	"github.com/fsm/fsm"
)

func main() {
	cli.Start(getStateMachine(), getStore())
}

func getStateMachine() fsm.StateMachine {
	// TODO
}

func getStore() fsm.Store {
	// TODO
}
```

## License

[MIT](LICENSE.md)
