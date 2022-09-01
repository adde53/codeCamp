{{/*
Create a fully qualified codecamp22 name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "sinch.cc2022.codecamp22.fullname" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "sinch.cc2022.codecamp22.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{/*
Labels.
*/}}
{{- define "sinch.cc2022.codecamp22.labels" -}}
app.kubernetes.io/name: {{ template "sinch.cc2022.codecamp22.fullname" . }}
app.kubernetes.io/version: {{ .Chart.Version }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ template "sinch.cc2022.codecamp22.chart" . }}
{{- end -}}

{/*
Selector labels
*/}}
{{- define "sinch.cc2022.codecamp22.selectorLabels" -}}
app.kubernetes.io/name: {{ include "sinch.cc2022.codecamp22.fullname" . }}
{{- end -}}


{{/*
Create image registry url
*/}}
{{- define "sinch.cc2022.codecamp22.registryUrl" -}}
{{- .Values.imageCredentials.registry.url -}}
{{- end -}}