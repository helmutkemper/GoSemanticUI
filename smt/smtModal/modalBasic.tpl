{{define "modalBasic"}}<div class="ui basic modal">
    <i class="close icon"></i>
    <div class="header">
        {{.Title}}
    </div>
    <div class="image content">
        {{if .Icon}}
        <div class="image">
            <i class="{{.Icon.Name}} {{.Icon.Size}} icon"></i>
        </div>
        {{ end}}
        <div class="description">
            <p>{{.Description}}</p>
        </div>
    </div>
    <div class="actions">
        <div class="two fluid ui inverted buttons">
            {{if .Cancel.Color -}}
            <div class="ui cancel {{.Cancel.Color}} basic inverted button">
                <i class="remove {{.Cancel.Color}} icon"></i>
                {{.Cancel.Text}}
            </div>
            {{- end}}
            {{if .Ok.Color -}}
            <div class="ui ok {{.Ok.Color}} basic inverted button">
                <i class="checkmark {{.Ok.Color}} icon"></i>
                {{.Ok.Text}}
            </div>
            {{- end}}
        </div>
    </div>
</div>{{end}}