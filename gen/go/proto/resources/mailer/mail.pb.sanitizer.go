// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/mailer/mail.proto

package mailer

import (
	"github.com/fivenet-app/fivenet/pkg/htmlsanitizer"
)

func (m *Mail) Sanitize() error {

	m.Body = htmlsanitizer.Sanitize(m.Body)

	m.Subject = htmlsanitizer.Sanitize(m.Subject)

	return nil
}