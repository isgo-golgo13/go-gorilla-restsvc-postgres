{{/*
Expand the name of the chart.
*/}}
{{- define "go-gorilla-restsvc-postgres.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "go-gorilla-restsvc-postgres.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "go-gorilla-restsvc-postgres.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels 
*/}}
{{- define "go-gorilla-restsvc-postgres.labels" -}}
helm.sh/chart: {{ include "go-gorilla-restsvc-postgres.chart" . }}
{{ include "go-gorilla-restsvc-postgres.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}



{{/*
Selector labels
*/}}
{{- define "go-gorilla-restsvc-postgres.selectorLabels" -}}
app.kubernetes.io/app: {{ include "go-gorilla-restsvc-postgres.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}




{{/*
Create the name of the service account to use
*/}}
{{- define "go-gorilla-restsvc-postgres.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "go-gorilla-restsvc-postgres.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}