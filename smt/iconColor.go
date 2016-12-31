package smt

type Color int

const (
	COLOR_RED Color = iota
  COLOR_ORANGE
  COLOR_YELLOW
  COLOR_OLIVE
  COLOR_GREEN
  COLOR_TEAL
  COLOR_BLUE
  COLOR_VIOLET
  COLOR_PURPLE
  COLOR_PINK
  COLOR_BROWN
  COLOR_GREY
  COLOR_BLACK
)

var colors = [...]string{
  "red",
  "orange",
  "yellow",
  "olive",
  "green",
  "teal",
  "blue",
  "violet",
  "purple",
  "pink",
  "brown",
  "grey",
  "black",
}

func (m Color) String() string {
  return colors[ m ]
}