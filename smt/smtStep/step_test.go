package smtStep

import (
  "smt/smtIcon"
  "smt/smtIconName"
  "fmt"
  "smt/smtStepState"
  "def"
  "smt/smtSize"
)

func ExampleStep_Get() {
  s1 := Step{
    Icon: smtIcon.Icon{
      Name: smtIconName.LIST,
    },
    States: def.A{ smtStepState.ACTIVE },
    Title: "Download",
    Description: "Select map source for download",
  }

  s2 := Step{
    Icon: smtIcon.Icon{
      Name: smtIconName.WAIT,
    },
    Title: "Process",
    Description: "Wait process end",
  }

  s := Steps{
    Steps: def.String{ s1.Get(), s2.Get() },
    Size: smtSize.TINY,
  }

  fmt.Print( s.Get() )

  // Output:
  // .
}
