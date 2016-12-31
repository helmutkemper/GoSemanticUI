package smt

type IconModifier int

const (
	ICON_LEFT IconModifier = iota
  ICON_FOCUS
  ICON_DISABLED
  ICON_LOADING
  ICON_ERROR
)

var iconModifiers = [...]string{
  "left",
  "focus",
  "disabled",
  "loading",
  "error",
}

func (m IconModifier) String() string {
  return iconModifiers[ m ]
}
