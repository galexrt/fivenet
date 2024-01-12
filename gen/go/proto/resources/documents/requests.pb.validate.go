// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/documents/requests.proto

package documents

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

// Validate checks the field values on DocRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DocRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DocRequestMultiError, or
// nil if none found.
func (m *DocRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DocRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocRequestValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocRequestValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocRequestValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocRequestValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocRequestValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocRequestValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DocumentId

	if _, ok := _DocRequest_RequestType_InLookup[m.GetRequestType()]; !ok {
		err := DocRequestValidationError{
			field:  "RequestType",
			reason: "value must be in list [DOC_ACTIVITY_TYPE_REQUESTED_ACCESS DOC_ACTIVITY_TYPE_REQUESTED_CLOSURE DOC_ACTIVITY_TYPE_REQUESTED_OPENING DOC_ACTIVITY_TYPE_REQUESTED_UPDATE DOC_ACTIVITY_TYPE_REQUESTED_OWNER_CHANGE DOC_ACTIVITY_TYPE_REQUESTED_DELETION]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCreatorJob()) > 20 {
		err := DocRequestValidationError{
			field:  "CreatorJob",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.CreatorId != nil {
		// no validation rules for CreatorId
	}

	if m.Creator != nil {

		if all {
			switch v := interface{}(m.GetCreator()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocRequestValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocRequestValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocRequestValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.CreatorJobLabel != nil {

		if utf8.RuneCountInString(m.GetCreatorJobLabel()) > 50 {
			err := DocRequestValidationError{
				field:  "CreatorJobLabel",
				reason: "value length must be at most 50 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Reason != nil {

		if utf8.RuneCountInString(m.GetReason()) > 255 {
			err := DocRequestValidationError{
				field:  "Reason",
				reason: "value length must be at most 255 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Accepted != nil {
		// no validation rules for Accepted
	}

	if len(errors) > 0 {
		return DocRequestMultiError(errors)
	}

	return nil
}

// DocRequestMultiError is an error wrapping multiple validation errors
// returned by DocRequest.ValidateAll() if the designated constraints aren't met.
type DocRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocRequestMultiError) AllErrors() []error { return m }

// DocRequestValidationError is the validation error returned by
// DocRequest.Validate if the designated constraints aren't met.
type DocRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocRequestValidationError) ErrorName() string { return "DocRequestValidationError" }

// Error satisfies the builtin error interface
func (e DocRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocRequestValidationError{}

var _DocRequest_RequestType_InLookup = map[DocActivityType]struct{}{
	13: {},
	14: {},
	15: {},
	16: {},
	17: {},
	18: {},
}