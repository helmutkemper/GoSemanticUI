package html

import (
  "def"
  "fmt"
)

func ExampleTag_Get() {
  t := Tag{
    Name:        "div",
    Value:       def.String{ "&nbsp;" },
    Attributes:  Attributes{
      Attribute{
        Name: "class",
        Values: def.String{ "ui", "input" },
      },
      Attribute{
        Name: "id",
        Values: def.String{ "user" },
      },
    },
  }

  fmt.Print( t.GetClosed() )

  // Output:
  // <div class="ui input" id="user">&nbsp;</div>
}