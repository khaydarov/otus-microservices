{{ define "app_envs" }}
- name: DATABASE_URI
  valueFrom:
    secretKeyRef:
      key: DATABASE_URI
      name: {{ .Chart.Name }}-secret

- name: APP_PORT
  value: {{ .Values.deployment.pod.port | quote }}

- name: JWT_SECRET
  valueFrom:
    secretKeyRef:
      key: JWT_SECRET
      name: {{ .Chart.Name }}-secret

- name: TRACER_PROVIDER
  value: {{ .Values.jaeger.dns }}

- name: BILLING_HOST
  value: {{ .Values.billing.host }}
{{ end }}