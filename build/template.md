# My LeetCode Solutions
Total: {{ .Total }}

# Usage
```shell script
make update
```

# Already Submit

| ID | Title | Difficulty |
| :----:| :----: | :----: |
{{- with .Details}}
{{- range .}}
| {{.Id}} | [{{.Title}}]({{.Path}}) | {{.Difficulty}} |
{{- end}}
{{- end}}
