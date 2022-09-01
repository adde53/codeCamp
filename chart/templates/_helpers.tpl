{{/*
Create a fully qualified codeCamp name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "sinch.cc2022.codeCamp.fullname" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "sinch.cc2022.codeCamp.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{/*
Labels.
*/}}
{{- define "sinch.cc2022.codeCamp.labels" -}}
app.kubernetes.io/name: {{ template "sinch.cc2022.codeCamp.fullname" . }}
app.kubernetes.io/version: {{ .Chart.Version }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ template "sinch.cc2022.codeCamp.chart" . }}
{{- end -}}

{/*
Selector labels
*/}}
{{- define "sinch.cc2022.codeCamp.selectorLabels" -}}
app.kubernetes.io/name: {{ include "sinch.cc2022.codeCamp.fullname" . }}
{{- end -}}


{{/*
Create image registry url
*/}}
{{- define "sinch.cc2022.codeCamp.registryUrl" -}}
{{- .Values.imageCredentials.registry.url -}}
{{- end -}}