// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/documents/templates.proto

package documents

import (
	"github.com/galexrt/fivenet/pkg/htmlsanitizer"
)

func (m *Template) Sanitize() error {

	m.Description = htmlsanitizer.Sanitize(m.Description)

	m.Title = htmlsanitizer.Sanitize(m.Title)

	return nil
}

func (m *TemplateShort) Sanitize() error {

	m.Description = htmlsanitizer.Sanitize(m.Description)

	m.Title = htmlsanitizer.Sanitize(m.Title)

	return nil
}
