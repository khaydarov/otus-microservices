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

- name: CASHBOOK_ID
  value: a8444ad0-590e-44e5-8b3b-cc5e10feb361

- name: REVENUE_ID
  value: b8444ad0-590e-44e5-8b3b-cc5e10feb362
{{ end }}