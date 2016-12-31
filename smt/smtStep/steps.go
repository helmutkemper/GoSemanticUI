package smtStep

import (
  "def"
  "smt/html"
)

type Steps struct {
  Steps      def.String
  Size       def.V
  Variations def.A
}

func ( StepsAStt *Steps ) Get() string {

  classSteps := html.Attribute{
    Name: "class",
    Values: def.String{ "ui" },
  }
  classSteps.Values.Add( StepsAStt.Variations, StepsAStt.Size, "steps" )

  divSteps := html.Tag{
    Name: "div",
    Attributes: html.Attributes{ classSteps },
  }
  divSteps.Value.Add( StepsAStt.Steps.Get() )

  return divSteps.GetClosed()
}