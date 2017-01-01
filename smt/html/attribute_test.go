package html

import (
  "def"
  "fmt"
)

func ExampleAttribute_Get() {
  a := Attribute{
    Name: "class",
    Values: def.String{ "ui input" },
  }
  fmt.Print( a.Get() )

  // Output:
  // class="ui input"
}
