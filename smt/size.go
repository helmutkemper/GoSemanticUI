package smt

type Size int

const (
  SIZE_MEDIUM Size = iota
  SIZE_MINI
  SIZE_TINY
  SIZE_SMALL
  SIZE_NORMAL
  SIZE_LARGE
  SIZE_BIG
  SIZE_HUGE
  SIZE_MASSIVE
)

var sizes = [...]string{
  "medium",
  "mini",
  "tiny",
  "small",
  "",
  "large",
  "big",
  "huge",
  "massive",
}


func (m Size) String() string {
  return sizes[ m ]
}