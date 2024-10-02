package protostub

import (
	"bytes"
	_ "embed"
	text_template "text/template"
)

//go:embed resources/service.tmpl
var serviceTmpl []byte

func RenderTemplate(serviceStub *ServiceStub) ([]byte, error) {
	tmpl, err := text_template.New("Text Template").Parse(string(serviceTmpl))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, serviceStub); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
