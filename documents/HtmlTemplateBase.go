package documents

import (
	"html/template"
	"strings"
)

type HtmlTemplateBase struct {
	templatePath string
	template     *template.Template
	model        interface{}
}

func (h *HtmlTemplateBase) Name() string {
	return h.templatePath
}

func (h *HtmlTemplateBase) ToString() string {

	pageContent := &strings.Builder{}
	h.template.Execute(pageContent, h.model)

	return pageContent.String()
}

func NewHtmlTemplateBase(templatePath string, content string, model interface{}) (*HtmlTemplateBase, error) {
	templateBase := &HtmlTemplateBase{
		templatePath: templatePath,
		model:        model,
	}

	pageTemplate, err := template.New(templateBase.templatePath).Parse(content)

	if err != nil {
		return templateBase, err
	}

	templateBase.template = pageTemplate

	return templateBase, nil
}
