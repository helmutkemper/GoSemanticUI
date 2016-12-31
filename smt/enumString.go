package smt

var enums []string

type V int
type A []V

func (e V) String() string {
  return enums[int(e)]
}

func Ciota(s string) V {
  enums = append(enums, s)
  return V(len(enums) - 1)
}