// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/livemap/livemap.proto

package livemap

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

// Validate checks the field values on DispatchMarker with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DispatchMarker) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DispatchMarker with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DispatchMarkerMultiError,
// or nil if none found.
func (m *DispatchMarker) ValidateAll() error {
	return m.validate(true)
}

func (m *DispatchMarker) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for X

	// no validation rules for Y

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DispatchMarkerValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DispatchMarkerValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DispatchMarkerValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetId() <= 0 {
		err := DispatchMarkerValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJob()) > 50 {
		err := DispatchMarkerValidationError{
			field:  "Job",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJobLabel()) > 50 {
		err := DispatchMarkerValidationError{
			field:  "JobLabel",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Icon

	// no validation rules for IconColor

	// no validation rules for Popup

	if len(errors) > 0 {
		return DispatchMarkerMultiError(errors)
	}

	return nil
}

// DispatchMarkerMultiError is an error wrapping multiple validation errors
// returned by DispatchMarker.ValidateAll() if the designated constraints
// aren't met.
type DispatchMarkerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchMarkerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchMarkerMultiError) AllErrors() []error { return m }

// DispatchMarkerValidationError is the validation error returned by
// DispatchMarker.Validate if the designated constraints aren't met.
type DispatchMarkerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchMarkerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchMarkerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchMarkerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchMarkerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchMarkerValidationError) ErrorName() string { return "DispatchMarkerValidationError" }

// Error satisfies the builtin error interface
func (e DispatchMarkerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatchMarker.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchMarkerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchMarkerValidationError{}

// Validate checks the field values on UserMarker with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserMarker) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserMarker with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserMarkerMultiError, or
// nil if none found.
func (m *UserMarker) ValidateAll() error {
	return m.validate(true)
}

func (m *UserMarker) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for X

	// no validation rules for Y

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserMarkerValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserMarkerValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserMarkerValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetId() <= 0 {
		err := UserMarkerValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Icon

	// no validation rules for IconColor

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserMarkerValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserMarkerValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserMarkerValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMarkerMultiError(errors)
	}

	return nil
}

// UserMarkerMultiError is an error wrapping multiple validation errors
// returned by UserMarker.ValidateAll() if the designated constraints aren't met.
type UserMarkerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMarkerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMarkerMultiError) AllErrors() []error { return m }

// UserMarkerValidationError is the validation error returned by
// UserMarker.Validate if the designated constraints aren't met.
type UserMarkerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserMarkerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserMarkerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserMarkerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserMarkerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserMarkerValidationError) ErrorName() string { return "UserMarkerValidationError" }

// Error satisfies the builtin error interface
func (e UserMarkerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserMarker.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserMarkerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserMarkerValidationError{}
