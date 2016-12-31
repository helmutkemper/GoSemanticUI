package smt

import (
  "github.com/helmutkemper/gOsm/utilMath"
  "strconv"
)

type line struct {
  Line string
}

type Header struct {
  Header            []string
  JavaScriptGlobal  []string
  JavaScriptOnReady []string
  Initialized         bool
}

func ( HeaderAStt *Header ) init() {
  HeaderAStt.Header            = make( []string, 0 )
  HeaderAStt.JavaScriptGlobal  = make( []string, 0 )
  HeaderAStt.JavaScriptOnReady = make( []string, 0 )
  HeaderAStt.Initialized       = true;
}

func ( HeaderAStt *Header ) SetTitle ( title string ) {
  if HeaderAStt.Initialized == false {
    HeaderAStt.init()
  }

  var n int
  HeaderAStt.Header, n = ExtendLine( HeaderAStt.Header, 3 )
  HeaderAStt.Header[ n ] = `<title>` + title + `</title>`
}

func ( HeaderAStt *Header ) AddMap ( divIdAStr string, latitudeAFlt float64, longitudeAFlt float64, zoomAInt int ) {
  if HeaderAStt.Initialized == false {
    HeaderAStt.init()
  }

  var n int
  HeaderAStt.JavaScriptGlobal, n = ExtendLine( HeaderAStt.JavaScriptGlobal, 1 )
  HeaderAStt.JavaScriptGlobal[ n ] = `
  var ` + divIdAStr + `;
  `

  HeaderAStt.JavaScriptOnReady, n = ExtendLine( HeaderAStt.JavaScriptOnReady, 1 )
  HeaderAStt.JavaScriptOnReady[ n ] = `
  ` + divIdAStr + ` = new L.Map('` + divIdAStr + `');
  var osmUrl='http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
  var osmAttrib='Map data © <a href="http://openstreetmap.org">OpenStreetMap</a> contributors';
  var osm = new L.TileLayer( osmUrl, { attribution: osmAttrib } );
  ` + divIdAStr + `.setView( new L.LatLng( ` + utilMath.FloatToString( latitudeAFlt ) + `,` + utilMath.FloatToString( longitudeAFlt ) + ` ), ` + strconv.Itoa( zoomAInt ) + ` );
  ` + divIdAStr + `.addLayer( osm );
  `

  HeaderAStt.Header, n = ExtendLine( HeaderAStt.Header, 2 )
  HeaderAStt.Header[ n + 0 ] = `<link rel="stylesheet" href="./static/js/leafletjs/leaflet.css" />`
  HeaderAStt.Header[ n + 1 ] = `<script src="./static/js/leafletjs/leaflet.js"></script>`
}

func ( HeaderAStt *Header ) AddJsOnReady ( lineAStr string ) {
  var n int

  HeaderAStt.JavaScriptOnReady, n = ExtendLine( HeaderAStt.JavaScriptOnReady, 1 )
  HeaderAStt.JavaScriptOnReady[ n ] = lineAStr
}

func ( HeaderAStt *Header ) AddJsGlobal ( lineAStr string ) {
  var n int

  HeaderAStt.JavaScriptGlobal, n = ExtendLine( HeaderAStt.JavaScriptGlobal, 1 )
  HeaderAStt.JavaScriptGlobal[ n ] = lineAStr
}

func ( HeaderAStt *Header ) AddFw () {
  if HeaderAStt.Initialized == false {
    HeaderAStt.init()
  }

  var n int
  HeaderAStt.Header, n = ExtendLine( HeaderAStt.Header, 19 )
  HeaderAStt.Header[ n + 0  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/semantic.css">`
  HeaderAStt.Header[ n + 1  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/reset.css">`
  HeaderAStt.Header[ n + 2  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/site.css">`
  HeaderAStt.Header[ n + 3  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/grid.css">`
  HeaderAStt.Header[ n + 4  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/icon.css">`
  HeaderAStt.Header[ n + 5  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/input.css">`
  HeaderAStt.Header[ n + 6  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/button.css">`
  HeaderAStt.Header[ n + 7  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/divider.css">`
  HeaderAStt.Header[ n + 8  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/label.css">`
  HeaderAStt.Header[ n + 9  ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/dropdown.css">`
  HeaderAStt.Header[ n + 10 ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/transition.css">`
  HeaderAStt.Header[ n + 11 ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/components/popup.css">`
  HeaderAStt.Header[ n + 12 ] = `<link rel="stylesheet" type="text/css" href="./static/js/Semantic-UI/dist/semantic.min.css" />`
  HeaderAStt.Header[ n + 13 ] = `<script src="./static/js/jquery-3.1.1.min.js"></script>`
  HeaderAStt.Header[ n + 14 ] = `<script src="./static/js/iframe.js"></script>`
  HeaderAStt.Header[ n + 15 ] = `<script type="text/javascript" src="./static/js/Semantic-UI/dist/components/popup.js"></script>`
  HeaderAStt.Header[ n + 16 ] = `<script type="text/javascript" src="./static/js/Semantic-UI/dist/components/dropdown.js"></script>`
  HeaderAStt.Header[ n + 17 ] = `<script type="text/javascript" src="./static/js/Semantic-UI/dist/components/transition.js"></script>`
  HeaderAStt.Header[ n + 18 ] = `<script src="./static/js/Semantic-UI/dist/semantic.min.js"></script>`
}

func ( HeaderAStt *Header ) MetaAdd () {
  if HeaderAStt.Initialized == false {
    HeaderAStt.init()
  }

  var n int
  HeaderAStt.Header, n = ExtendLine( HeaderAStt.Header, 3 )
  HeaderAStt.Header[ n + 0 ] = `<meta charset="utf-8" />`
  HeaderAStt.Header[ n + 1 ] = `<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />`
  HeaderAStt.Header[ n + 2 ] = `<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=5.0">`
}

func ( HeaderAStt *Header ) GetLinksTemplateName () string {
  return `smtLinks`
}

func ( HeaderAStt *Header ) GetLinksTemplate () string {
  return `{{define "smtLinks"}}
  {{- range $k, $v := .Header }}
    {{ htmlSafe $v  }}
  {{- end }}
{{end}}`
}

func ( HeaderAStt *Header ) GetJavaScriptTemplate () string {
  return `{{define "smtJavaScript"}}
  {{- range $k, $v := .JavaScriptGlobal }}
    {{ htmlSafe $v  }}
  {{- end }}
  $( document ).ready(
    function(){
      {{- range $k, $v := .JavaScriptOnReady }}
        {{ htmlSafe $v  }}
      {{- end}}
  });
{{end}}`
}

func ( HeaderAStt *Header ) GetJavaScriptTemplateName () string {
  return `smtJavaScript`
}

func ExtendLine( slice []string, addSpace int ) ( []string, int ) {
  n := len( slice )
  if n == cap( slice ) {
    newHeader := make( []string, len( slice ), len( slice ) + addSpace )
    copy( newHeader, slice )
    slice = newHeader
  }

  slice = slice[ 0 : n+addSpace ]
  return slice, n
}