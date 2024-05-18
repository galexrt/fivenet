// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: resources/messenger/mail.proto

package messenger

import (
	"github.com/fivenet-app/fivenet/pkg/htmlsanitizer"
)

func (m *Mail) Sanitize() error {

	m.Body = htmlsanitizer.Sanitize(m.Body)

	m.Subject = htmlsanitizer.Sanitize(m.Subject)

	return nil
}