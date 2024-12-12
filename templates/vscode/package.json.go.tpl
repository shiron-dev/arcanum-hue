{
  "name": "{{.Name | kebabcase }}",
  "displayName": "{{.Name | kebabcase }}",
  "description": "{{.Description}}",
  "version": "{{.Version}}",
  "engines": {
    "vscode": "^1.95.0"
  },
  "categories": [
    "Themes"
  ],
  "contributes": {
    "themes": [
      {{ range $i, $v := .Themes }}
      {
        "label": "{{$v.Name}}",
        "uiTheme": "{{ $v.UITheme }}",
        "path": "./themes/{{ $v.Name | kebabcase }}-color-theme.json"
      }{{ if ne $i (sub (len $.Themes) 1) }},{{ end }}
      {{ end }}
    ]
  }
}
