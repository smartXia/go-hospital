package global

{{- if .HasGlobal }}

import "devops-manage/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}