// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/jobs/qualifications.proto

package jobs

import (
	"github.com/galexrt/fivenet/pkg/htmlsanitizer"
)

func (m *Qualification) Sanitize() error {

	m.Abbreviation = htmlsanitizer.StripTags(m.Abbreviation)

	m.Description = htmlsanitizer.Sanitize(m.Description)

	m.Title = htmlsanitizer.Sanitize(m.Title)

	return nil
}

func (m *QualificationResult) Sanitize() error {

	m.Summary = htmlsanitizer.StripTags(m.Summary)

	return nil
}
