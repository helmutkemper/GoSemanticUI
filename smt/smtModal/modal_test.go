package smtModal

import (
  "smt/smtSize"
  "smt/smtColor"
  "fmt"
)

func ExampleModalProfile_Get() {
  m := ModalProfile{
    Title: "Olá Mundo!",
    Description: Description{
      Header: "Temos um problema.",
      Text: "Há uma extenção em massa em andamento!",
    },
    Image: Image{
      Size: smtSize.SMALL,
      Url: "http://semantic-ui.com/images/avatar2/large/rachel.png",
    },
    ButtonOk: Button{
      Text: "Ok",
      Color: smtColor.GREEN,
    },
    ButtonCancel: Button{
      Text: "Cancel",
      Color: smtColor.RED,
    },
  }
  fmt.Print( m.Get() )

  // Output:
  // .
}