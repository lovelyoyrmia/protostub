package protostub

import (
	"bytes"
	_ "embed"
	text_template "text/template"
)

var (
	//go:embed resources/service.tmpl
	serviceTmpl []byte

	//go:embed resources/client.tmpl
	clientTmpl []byte
)

func RenderTemplate(kind ProtoStubType, serviceStub *ServiceStub) ([]byte, error) {
	templString, err := kind.renderer()
	if err != nil {
		return nil, err
	}

	tmpl, err := text_template.New("Text Template").Parse(templString)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, serviceStub); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
