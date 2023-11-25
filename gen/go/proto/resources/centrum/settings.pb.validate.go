// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/centrum/settings.proto

package centrum

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Settings with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Settings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Settings with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SettingsMultiError, or nil
// if none found.
func (m *Settings) ValidateAll() error {
	return m.validate(true)
}

func (m *Settings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := SettingsValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Enabled

	if _, ok := CentrumMode_name[int32(m.GetMode())]; !ok {
		err := SettingsValidationError{
			field:  "Mode",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := CentrumMode_name[int32(m.GetFallbackMode())]; !ok {
		err := SettingsValidationError{
			field:  "FallbackMode",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SettingsMultiError(errors)
	}

	return nil
}

// SettingsMultiError is an error wrapping multiple validation errors returned
// by Settings.ValidateAll() if the designated constraints aren't met.
type SettingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SettingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SettingsMultiError) AllErrors() []error { return m }

// SettingsValidationError is the validation error returned by
// Settings.Validate if the designated constraints aren't met.
type SettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SettingsValidationError) ErrorName() string { return "SettingsValidationError" }

// Error satisfies the builtin error interface
func (e SettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SettingsValidationError{}
