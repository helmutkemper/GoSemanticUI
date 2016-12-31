{{define "autoComplete"}}<div class="ui search">
    <div class="ui {{if .Icon.Left}}left{{- end}} icon input">
        <input class="prompt" type="text" placeholder="{{.Placeholder}}">
        <i class="{{.Icon.Name}} icon"></i>
    </div>
</div>{{end}}