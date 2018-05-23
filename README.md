<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# CLI

[![Version](https://img.shields.io/github/tag/fsm/cli.svg)](https://github.com/fsm/cli/releases)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/fsm/cli/blob/master/LICENSE.md) 
[![Go Report Card](https://goreportcard.com/badge/github.com/fsm/cli)](https://goreportcard.com/report/github.com/fsm/cli) 
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/fsm/Lobby)

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
