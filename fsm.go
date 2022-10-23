package fsm

type FSM struct {
  state   stateType
  states  []stateType
  initial stateType
  events  map[eventType]*Event
}

func (fsm *FSM) DefineStates(states ...stateType) {
  if len(states) == 0 {
    panic("need to have atleast one state")
  }
  fsm.initial = states[0]
  fsm.state = fsm.initial
  fsm.states = states
}

func (fsm *FSM) Hydrate(currentState stateType) {
  var validState bool

  for _, state := range fsm.states {
    if currentState == state {
      validState = true
      break
    }
  }

  if !validState {
    panic("cannot hydrate with an undefined state")
  }

  fsm.state = currentState
}

func (fsm *FSM) State() stateType {
  return fsm.state
}

func (fsm *FSM) On(eventName eventType, registrationFn func(event *Event)) {
  newEvent := newEvent(fsm)
  registrationFn(newEvent)
  fsm.events[eventName] = newEvent
}

func (fsm *FSM) Fire(eventName eventType, eventArgs ...any) (stateType, error) {
  return fsm.events[eventName].fire(eventArgs...)
}

func New() *FSM {
  return &FSM{
    events: map[eventType]*Event{},
  }
}
