<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>
<p align="center">
  <a href="https://github.com/fsm/cli/releases"><img src="https://img.shields.io/github/tag/fsm/cli.svg" alt="Version"></img></a>
  <a href="https://github.com/fsm/cli/blob/master/LICENSE.md"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT License"></img></a>
  <a href="https://goreportcard.com/report/github.com/fsm/cli"><img src="https://goreportcard.com/badge/github.com/fsm/cli" alt="Go Report Card"></img></a>
  <a href="https://spectrum.chat/fsm"><img alt="Join the community on Spectrum" src="https://withspectrum.github.io/badge/badge.svg"/></a>
</p>

# CLI

CLI is a command line target for [FSM](https://github.com/fsm/fsm).

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
	// ...
}

func getStore() fsm.Store {
	// ...
}
```
