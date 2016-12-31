package smt

type Variation int

const (
  ICON_VARIATION_FITTED Variation = iota
  ICON_VARIATION_LINK
  ICON_VARIATION_FLIPPED
  ICON_VARIATION_ROTATED
  ICON_VARIATION_CIRCULAR
  ICON_VARIATION_BORDERED
  ICON_VARIATION_INVERTED
)

var variations = [...]string {
  "fitted",
  "link",
  "flipped",
  "rotated",
  "circular",
  "bordered",
  "inverted",
}

func (m Variation) String() string {
  return variations[ m ]
}
