package fsm

import "fmt"

type Event struct {
  beforeFn           func(stateType, ...any) bool
  beforeFnRegistered bool
  tns                map[stateType]stateType
  afterFn            func(stateType, ...any)
  afterFnRegistered  bool
  fsm                *FSM
}

func (e *Event) fire(eventArgs ...any) (stateType, error) {
  if e.beforeFnRegistered && !e.beforeFn(e.fsm.state, eventArgs...) {
    return e.fsm.state, fmt.Errorf("before function prevented transition")
  }

  if _, transitionPresent := e.tns[e.fsm.state]; !transitionPresent {
    return e.fsm.state, fmt.Errorf("fired event does not trigger a transition")
  }

  if e.fsm.state == e.tns[e.fsm.state] {
    return e.fsm.state, fmt.Errorf("the transition didnt change the state")
  }

  e.fsm.state = e.tns[e.fsm.state]

  if e.afterFnRegistered {
    e.afterFn(e.fsm.state, eventArgs...)
  }

  return e.fsm.state, nil
}

func (e *Event) BeforeTn(beforeFn func(state stateType, eventArgs ...any) bool) {
  e.beforeFn = beforeFn
  e.beforeFnRegistered = true
}

func (e *Event) AfterTn(afterFn func(state stateType, eventArgs ...any)) {
  e.afterFn = afterFn
  e.afterFnRegistered = true
}

func (e *Event) Tn(fromState stateType, toState stateType) {
  e.tns[fromState] = toState
}

func newEvent(fsm *FSM) *Event {
  return &Event{
    fsm: fsm,
    tns: map[stateType]stateType{},
  }
}
