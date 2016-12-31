package smt

import (
  "reflect"
  "bytes"
)

type Button struct{
  Label           string
  Primary         bool
  Secondary       bool
  Animated        bool
  ContentHidden   Content
  ContentVisible  Content
  Icon            *Icon
  Modifiers       []string

  Variations      []string
  Size            string
  Color           string
}

func ( ButtonAStt Button ) IsEmpty() bool {
  return reflect.DeepEqual( ButtonAStt, Button{} )
}

func ( ButtonAStt Button ) Get () string {
  var buffer bytes.Buffer

  if ButtonAStt.Primary == true {
    buffer.WriteString( "primary " )
  } else if ButtonAStt.Secondary == true {
    buffer.WriteString( "secondary " )
  }

  return `{{define "button"}}<button class="ui {{- range $k, $v := .Variations }}{{ $v }}
  {{- end }}{{.Size}} {{.Color}} {{.Name}} {{- range $k, $v := .Modifiers }} {{ $v }}
  {{- end }}button">{{.Label}}</button>{{end}}`
}
