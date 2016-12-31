package smt

import (
  "reflect"
)

type Icon struct{
  Name       IconName
  Variations []Variation
  Size       Size
  Color      Color
  Modifiers  []IconModifier
}

func ( IconAStt Icon ) IsEmpty() bool {
  return reflect.DeepEqual( IconAStt, Icon{} )
}

func ( IconAStt Icon ) Get () string {
  return `{{define "icon"}}<i class="{{- range $k, $v := .Variations }}{{ $v }}
  {{- end }} {{.Size}} {{.Color}} {{.Name}} {{- range $k, $v := .Modifiers }} {{ $v }}
  {{- end }} icon"></i>{{end}}`
}
