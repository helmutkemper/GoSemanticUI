package smt

import (
  "reflect"
  "bytes"
)

type Content struct{
  Icon      *Icon
  RightIcon bool
  Text      string
  Visible   bool
}

func ( ContentAStt Content ) IsEmpty() bool {
  return reflect.DeepEqual( ContentAStt, Content{} )
}

func ( ContentAStt Content ) Get () string {
  var buffer bytes.Buffer
  var icon Icon = ContentAStt.Icon

  buffer.WriteString( icon.Get() )

  buffer.WriteString( `<div class="{{if .Visible}}visible{{else}}hidden{{end}} content"><i class="right arrow icon"></i></div>` )

  if ContentAStt.RightIcon == true {
    buffer.WriteString( ContentAStt.Icon.Get() )
  }

  if ContentAStt.Text != "" {
    buffer.WriteString( ContentAStt.Text )
  }

  if ContentAStt.RightIcon == false {
    buffer.WriteString( ContentAStt.Icon.Get() )
  }

  buffer.WriteString( `</div>` )

  return buffer.String()
}
