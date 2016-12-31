package smt

type ButtonState int

const (
  ACTIVE ButtonState = iota
  DISABLED
  LOADING
)

var buttonStates = [...]string{
  "active",
  "disabled",
  "loading",
}

func (m ButtonState) String() string {
  return buttonStates[ m ]
}
