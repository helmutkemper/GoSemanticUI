package html

import (
  "bytes"
  "def"
)

// comment
type Tag struct {
  Name       string
  Value      def.String
  Attributes []Attribute
}

func ( TagAStt *Tag ) GetClosed () string {
  var returnLBfr bytes.Buffer
  length := len( TagAStt.Attributes ) - 1

  returnLBfr.WriteString( `<` + TagAStt.Name + ` ` )

  for k, v := range TagAStt.Attributes {
    returnLBfr.WriteString( v.Get() )

    if k != length {
      returnLBfr.WriteString( ` ` )
    }
  }

  returnLBfr.WriteString( `>` + TagAStt.Value.Get( " " ) + `</` + TagAStt.Name + `>` )

  return returnLBfr.String()
}

func ( TagAStt *Tag ) GetOpen () string {
  var returnLBfr bytes.Buffer
  length := len( TagAStt.Attributes ) - 1

  returnLBfr.WriteString( `<` + TagAStt.Name + ` ` )

  for k, v := range TagAStt.Attributes {
    returnLBfr.WriteString( v.Get() )

    if k != length {
      returnLBfr.WriteString( ` ` )
    }
  }

  returnLBfr.WriteString( `>` )

  return returnLBfr.String()
}
