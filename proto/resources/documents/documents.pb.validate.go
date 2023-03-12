// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/documents/documents.proto

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

// Validate checks the field values on Document with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Document) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Document with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DocumentMultiError, or nil
// if none found.
func (m *Document) ValidateAll() error {
	return m.validate(true)
}

func (m *Document) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 0 {
		err := DocumentValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocumentValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentValidationError{
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
				errors = append(errors, DocumentValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetTitle()) < 3 {
		err := DocumentValidationError{
			field:  "Title",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 30 {
		err := DocumentValidationError{
			field:  "Content",
			reason: "value length must be at least 30 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DOCUMENT_CONTENT_TYPE_name[int32(m.GetContentType())]; !ok {
		err := DocumentValidationError{
			field:  "ContentType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Closed

	if utf8.RuneCountInString(m.GetState()) > 24 {
		err := DocumentValidationError{
			field:  "State",
			reason: "value length must be at most 24 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCreator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocumentValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentValidationError{
				field:  "Creator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Public

	// no validation rules for CategoryID

	if m.GetTargetDocumentID() < 0 {
		err := DocumentValidationError{
			field:  "TargetDocumentID",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DocumentMultiError(errors)
	}

	return nil
}

// DocumentMultiError is an error wrapping multiple validation errors returned
// by Document.ValidateAll() if the designated constraints aren't met.
type DocumentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentMultiError) AllErrors() []error { return m }

// DocumentValidationError is the validation error returned by
// Document.Validate if the designated constraints aren't met.
type DocumentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentValidationError) ErrorName() string { return "DocumentValidationError" }

// Error satisfies the builtin error interface
func (e DocumentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocument.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentValidationError{}

// Validate checks the field values on DocumentTemplate with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DocumentTemplate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentTemplate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentTemplateMultiError, or nil if none found.
func (m *DocumentTemplate) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentTemplate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := DocumentTemplateValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetJobGrade() < 0 {
		err := DocumentTemplateValidationError{
			field:  "JobGrade",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTitle()) < 3 {
		err := DocumentTemplateValidationError{
			field:  "Title",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) > 255 {
		err := DocumentTemplateValidationError{
			field:  "Description",
			reason: "value length must be at most 255 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContentTitle()) < 3 {
		err := DocumentTemplateValidationError{
			field:  "ContentTitle",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 12 {
		err := DocumentTemplateValidationError{
			field:  "Content",
			reason: "value length must be at least 12 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for AdditionalData

	// no validation rules for CreatorID

	if len(errors) > 0 {
		return DocumentTemplateMultiError(errors)
	}

	return nil
}

// DocumentTemplateMultiError is an error wrapping multiple validation errors
// returned by DocumentTemplate.ValidateAll() if the designated constraints
// aren't met.
type DocumentTemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentTemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentTemplateMultiError) AllErrors() []error { return m }

// DocumentTemplateValidationError is the validation error returned by
// DocumentTemplate.Validate if the designated constraints aren't met.
type DocumentTemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentTemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentTemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentTemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentTemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentTemplateValidationError) ErrorName() string { return "DocumentTemplateValidationError" }

// Error satisfies the builtin error interface
func (e DocumentTemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentTemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentTemplateValidationError{}

// Validate checks the field values on DocumentTemplateShort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DocumentTemplateShort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentTemplateShort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentTemplateShortMultiError, or nil if none found.
func (m *DocumentTemplateShort) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentTemplateShort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Job

	// no validation rules for JobGrade

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for CreatorID

	if len(errors) > 0 {
		return DocumentTemplateShortMultiError(errors)
	}

	return nil
}

// DocumentTemplateShortMultiError is an error wrapping multiple validation
// errors returned by DocumentTemplateShort.ValidateAll() if the designated
// constraints aren't met.
type DocumentTemplateShortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentTemplateShortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentTemplateShortMultiError) AllErrors() []error { return m }

// DocumentTemplateShortValidationError is the validation error returned by
// DocumentTemplateShort.Validate if the designated constraints aren't met.
type DocumentTemplateShortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentTemplateShortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentTemplateShortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentTemplateShortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentTemplateShortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentTemplateShortValidationError) ErrorName() string {
	return "DocumentTemplateShortValidationError"
}

// Error satisfies the builtin error interface
func (e DocumentTemplateShortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentTemplateShort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentTemplateShortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentTemplateShortValidationError{}

// Validate checks the field values on DocumentAccess with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DocumentAccess) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentAccess with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DocumentAccessMultiError,
// or nil if none found.
func (m *DocumentAccess) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentAccess) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DocumentID

	for idx, item := range m.GetJobAccess() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentAccessValidationError{
						field:  fmt.Sprintf("JobAccess[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentAccessValidationError{
						field:  fmt.Sprintf("JobAccess[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentAccessValidationError{
					field:  fmt.Sprintf("JobAccess[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetUserAccess() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentAccessValidationError{
						field:  fmt.Sprintf("UserAccess[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentAccessValidationError{
						field:  fmt.Sprintf("UserAccess[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentAccessValidationError{
					field:  fmt.Sprintf("UserAccess[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DocumentAccessMultiError(errors)
	}

	return nil
}

// DocumentAccessMultiError is an error wrapping multiple validation errors
// returned by DocumentAccess.ValidateAll() if the designated constraints
// aren't met.
type DocumentAccessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentAccessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentAccessMultiError) AllErrors() []error { return m }

// DocumentAccessValidationError is the validation error returned by
// DocumentAccess.Validate if the designated constraints aren't met.
type DocumentAccessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentAccessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentAccessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentAccessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentAccessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentAccessValidationError) ErrorName() string { return "DocumentAccessValidationError" }

// Error satisfies the builtin error interface
func (e DocumentAccessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentAccess.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentAccessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentAccessValidationError{}

// Validate checks the field values on DocumentJobAccess with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DocumentJobAccess) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentJobAccess with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentJobAccessMultiError, or nil if none found.
func (m *DocumentJobAccess) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentJobAccess) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocumentJobAccessValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentJobAccessValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentJobAccessValidationError{
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
				errors = append(errors, DocumentJobAccessValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentJobAccessValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentJobAccessValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DocumentID

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := DocumentJobAccessValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMinimumGrade() < 0 {
		err := DocumentJobAccessValidationError{
			field:  "MinimumGrade",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DOCUMENT_ACCESS_name[int32(m.GetAccess())]; !ok {
		err := DocumentJobAccessValidationError{
			field:  "Access",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DocumentJobAccessMultiError(errors)
	}

	return nil
}

// DocumentJobAccessMultiError is an error wrapping multiple validation errors
// returned by DocumentJobAccess.ValidateAll() if the designated constraints
// aren't met.
type DocumentJobAccessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentJobAccessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentJobAccessMultiError) AllErrors() []error { return m }

// DocumentJobAccessValidationError is the validation error returned by
// DocumentJobAccess.Validate if the designated constraints aren't met.
type DocumentJobAccessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentJobAccessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentJobAccessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentJobAccessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentJobAccessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentJobAccessValidationError) ErrorName() string {
	return "DocumentJobAccessValidationError"
}

// Error satisfies the builtin error interface
func (e DocumentJobAccessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentJobAccess.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentJobAccessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentJobAccessValidationError{}

// Validate checks the field values on DocumentUserAccess with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DocumentUserAccess) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentUserAccess with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentUserAccessMultiError, or nil if none found.
func (m *DocumentUserAccess) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentUserAccess) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocumentUserAccessValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentUserAccessValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentUserAccessValidationError{
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
				errors = append(errors, DocumentUserAccessValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentUserAccessValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentUserAccessValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DocumentID

	if m.GetUserID() <= 0 {
		err := DocumentUserAccessValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DOCUMENT_ACCESS_name[int32(m.GetAccess())]; !ok {
		err := DocumentUserAccessValidationError{
			field:  "Access",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DocumentUserAccessMultiError(errors)
	}

	return nil
}

// DocumentUserAccessMultiError is an error wrapping multiple validation errors
// returned by DocumentUserAccess.ValidateAll() if the designated constraints
// aren't met.
type DocumentUserAccessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentUserAccessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentUserAccessMultiError) AllErrors() []error { return m }

// DocumentUserAccessValidationError is the validation error returned by
// DocumentUserAccess.Validate if the designated constraints aren't met.
type DocumentUserAccessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentUserAccessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentUserAccessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentUserAccessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentUserAccessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentUserAccessValidationError) ErrorName() string {
	return "DocumentUserAccessValidationError"
}

// Error satisfies the builtin error interface
func (e DocumentUserAccessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentUserAccess.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentUserAccessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentUserAccessValidationError{}

// Validate checks the field values on DocumentCategory with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DocumentCategory) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentCategory with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentCategoryMultiError, or nil if none found.
func (m *DocumentCategory) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentCategory) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 128 {
		err := DocumentCategoryValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) > 255 {
		err := DocumentCategoryValidationError{
			field:  "Description",
			reason: "value length must be at most 255 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJob()) > 20 {
		err := DocumentCategoryValidationError{
			field:  "Job",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DocumentCategoryMultiError(errors)
	}

	return nil
}

// DocumentCategoryMultiError is an error wrapping multiple validation errors
// returned by DocumentCategory.ValidateAll() if the designated constraints
// aren't met.
type DocumentCategoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentCategoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentCategoryMultiError) AllErrors() []error { return m }

// DocumentCategoryValidationError is the validation error returned by
// DocumentCategory.Validate if the designated constraints aren't met.
type DocumentCategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentCategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentCategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentCategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentCategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentCategoryValidationError) ErrorName() string { return "DocumentCategoryValidationError" }

// Error satisfies the builtin error interface
func (e DocumentCategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentCategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentCategoryValidationError{}

// Validate checks the field values on DocumentRelation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DocumentRelation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentRelation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentRelationMultiError, or nil if none found.
func (m *DocumentRelation) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentRelation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocumentRelationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentRelationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentRelationValidationError{
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
				errors = append(errors, DocumentRelationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentRelationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentRelationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DocumentID

	if m.GetTargetUserID() <= 0 {
		err := DocumentRelationValidationError{
			field:  "TargetUserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DOCUMENT_RELATION_name[int32(m.GetRelation())]; !ok {
		err := DocumentRelationValidationError{
			field:  "Relation",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCauseUserID() <= 0 {
		err := DocumentRelationValidationError{
			field:  "CauseUserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DocumentRelationMultiError(errors)
	}

	return nil
}

// DocumentRelationMultiError is an error wrapping multiple validation errors
// returned by DocumentRelation.ValidateAll() if the designated constraints
// aren't met.
type DocumentRelationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentRelationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentRelationMultiError) AllErrors() []error { return m }

// DocumentRelationValidationError is the validation error returned by
// DocumentRelation.Validate if the designated constraints aren't met.
type DocumentRelationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentRelationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentRelationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentRelationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentRelationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentRelationValidationError) ErrorName() string { return "DocumentRelationValidationError" }

// Error satisfies the builtin error interface
func (e DocumentRelationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentRelation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentRelationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentRelationValidationError{}
