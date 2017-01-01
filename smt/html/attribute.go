package html

import (
  "def"
)

type Attribute struct {
  Name   string
  Values def.String
}

type Attributes []Attribute

func ( AttributeAStt *Attribute ) Get () string {
  return AttributeAStt.Values.Get( " ", AttributeAStt.Name + `="`, `"` )
}

