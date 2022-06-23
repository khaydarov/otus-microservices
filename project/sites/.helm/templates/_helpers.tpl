{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define ".helm.fullname" -}}
{{ .Chart.Name }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define ".helm.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define ".helm.labels" -}}
app: {{ .Chart.Name }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define ".helm.selectorLabels" -}}
app: {{ .Chart.Name }}
{{- end }}