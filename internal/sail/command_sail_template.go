// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"bytes"
	"text/template"
)

// TemplateMapInterface holds templates for generating interface layer files.
var TemplateMapInterface = map[string]string{
	"internal/{{ .folder }}/interface/web/controller/{{ .entity }}_controller.go":      `package controller`,
	"internal/{{ .folder }}/interface/web/controller/{{ .entity }}_controller_test.go": `package controller`,

	"internal/{{ .folder }}/interface/web/mapper/{{ .entity }}_mapper.go":      `package mapper`,
	"internal/{{ .folder }}/interface/web/mapper/{{ .entity }}_mapper_test.go": `package mapper`,

	"internal/{{ .folder }}/interface/web/response/{{ .entity }}_response.go":      `package response`,
	"internal/{{ .folder }}/interface/web/response/{{ .entity }}_response_test.go": `package response`,
}

// RenderTemplatePath processa a chave do map usando os dados fornecidos.
func RenderTemplatePath(tmpl string, data map[string]string) (string, error) {
	t, err := template.New("path").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
