package smgg

import (
	"bytes"
	"text/template"
)

type genSourcer struct {
	source string
}

type GenSourcer interface {
	CreateSource(sourcePath string) (GenSourcer, error)
	GetSource() string
	ParseAndUpdateSource(tmpl GenSetterTemplate, data any) error
}

func NewTemplateParser() GenSourcer {

	return &genSourcer{source: ""}
}

func (tp *genSourcer) GetSource() string {
	return tp.source
}

func (tp *genSourcer) ParseAndUpdateSource(tmpl GenSetterTemplate, data any) error {
	// テンプレートをパース
	t, err := template.New("code").Parse(string(tmpl))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	defer buf.Reset()
	if err := t.Execute(&buf, data); err != nil {
		return err
	}

	tp.source += buf.String()

	return nil
}
