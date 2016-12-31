package smt

type ButtonModifier int

const(
  BUTTON_BASIC ButtonModifier = iota
  BUTTON_INVERTED
)

var buttonModifiers = [...]string{
  "basic",
  "inverted",
}

func (m ButtonModifier) String() string {
  return buttonModifiers[ m ]
}