{{define "modalProfile"}}<div class="ui modal">
    <i class="close icon"></i>
    <div class="header">
        {{.Title}}
    </div>
    <div class="image content">
        {{if .Image -}}
        <div class="ui {{.Image.Size}} image">
        {{- else}}
        <div class="ui medium image">
        {{- end}}
            <img src="{{.Image.Url}}">
        </div>
        <div class="description">
            <div class="ui header">{{.Header}}</div>
            <p>{{.Text}}</p>
        </div>
    </div>
    <div class="actions">
        {{if .Cancel.Color -}}
        <div class="ui {{.Cancel.Color}} deny right labeled icon button">
            {{.Cancel.Text}}
            {{if .Cancel.Icon -}}
            <i class="{{.Cancel.Icon}} icon"></i>
            {{- end}}
        </div>
        {{- end}}
        {{if .Ok.Color -}}
        <div class="ui positive {{.Ok.Color}} right labeled icon button">
            {{.Ok.Text}}
            {{if .Ok.Icon -}}
            <i class="{{.Ok.Icon}} icon"></i>
            {{- end}}
        </div>
        {{- end}}
    </div>
</div>{{end}}