{{- define "internal-service.name" -}}
{{- .Chart.Name }}
{{- end }}

{{- define "internal-service.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "internal-service.labels" -}}
app.kubernetes.io/name: {{ include "internal-service.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version }}
{{- end }}

{{- define "internal-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "internal-service.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
