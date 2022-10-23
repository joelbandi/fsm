# fsm

Finite State Machine implementation in golang inspired by [aasm/aasm](https://github.com/aasm/aasm)

### Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/joelbandi/fsm
```

### Usage (commented)

```go
package main

import (
  "fmt"

  "github.com/joelbandi/fsm"
)

func main() {
  // Define state entities
  const (
    on = "on"
    off = "off"
  )
  // Define event entities
  const (
    toggle = iota
  )

  // Instantiate a new fsm struct with int type entities.
  myfsm := fsm.New()

  // The first argument is always designated as the default state.
  myfsm.DefineStates(on, off)

  // Get current state
  fmt.Printf("intial state is %s\n", myfsm.State())

  myfsm.On(toggle, func(e *fsm.Event) {
    e.Tn(off, on) // Define a potential transition.
    e.Tn(on, off) // Define a potential transition.

    // If BeforeTn hook returns false then the transition will be skipped. Use this for conditional transitions.
    e.BeforeTn(func(beforeState string, eventArgs ...any) bool {
      arg := eventArgs[0].(string) // Assert the type.

      if arg == "skip" {
        fmt.Printf("skipped!!! - state remains %s\n", beforeState)
        return false // skip here
      } else {
        return true // dont skip
      }
    })

    // Use AfterTn hook for logging, writing to your persistence layer etc. This will not run if the Tn did not happen.
    e.AfterTn(func(afterState string, _ ...any) {
      fmt.Printf("toggled!!! - new State is %s\n", afterState)
    })
  })

  myfsm.Fire(toggle, "dont skip")
  myfsm.Fire(toggle, "never skip")
  myfsm.Fire(toggle, "skip") // skips
  myfsm.Fire(toggle, "skip") // skips
  myfsm.Fire(toggle, "DO NOT skip")

  // force set current state
  myfsm.Hydrate(on)
  fmt.Printf("new state after hydrating is %s\n", myfsm.State())
}
```

The output of this go file is
```
intial state is on
toggled!!! - new State is off
toggled!!! - new State is on
skipped!!! - state remains on
skipped!!! - state remains on
toggled!!! - new State is off
new state after hydrating is on
```
