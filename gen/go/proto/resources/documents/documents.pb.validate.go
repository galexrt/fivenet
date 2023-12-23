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

	// no validation rules for Id

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

	if len(m.GetTitle()) > 21845 {
		err := DocumentValidationError{
			field:  "Title",
			reason: "value length must be at most 21845 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DocContentType_name[int32(m.GetContentType())]; !ok {
		err := DocumentValidationError{
			field:  "ContentType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 20 {
		err := DocumentValidationError{
			field:  "Content",
			reason: "value length must be at least 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetContent()) > 1750000 {
		err := DocumentValidationError{
			field:  "Content",
			reason: "value length must be at most 1750000 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CreatorJob

	if utf8.RuneCountInString(m.GetState()) > 32 {
		err := DocumentValidationError{
			field:  "State",
			reason: "value length must be at most 32 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Closed

	// no validation rules for Public

	if m.UpdatedAt != nil {

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

	}

	if m.DeletedAt != nil {

		if all {
			switch v := interface{}(m.GetDeletedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.CategoryId != nil {
		// no validation rules for CategoryId
	}

	if m.Category != nil {

		if all {
			switch v := interface{}(m.GetCategory()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentValidationError{
						field:  "Category",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentValidationError{
						field:  "Category",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Data != nil {

		if len(m.GetData()) > 1000000 {
			err := DocumentValidationError{
				field:  "Data",
				reason: "value length must be at most 1000000 bytes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
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

// Validate checks the field values on DocumentShort with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DocumentShort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentShort with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DocumentShortMultiError, or
// nil if none found.
func (m *DocumentShort) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentShort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DocumentShortValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DocumentShortValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DocumentShortValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetTitle()) < 3 {
		err := DocumentShortValidationError{
			field:  "Title",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DocContentType_name[int32(m.GetContentType())]; !ok {
		err := DocumentShortValidationError{
			field:  "ContentType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetContent()) > 1024 {
		err := DocumentShortValidationError{
			field:  "Content",
			reason: "value length must be at most 1024 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CreatorJob

	if utf8.RuneCountInString(m.GetState()) > 32 {
		err := DocumentShortValidationError{
			field:  "State",
			reason: "value length must be at most 32 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Closed

	// no validation rules for Public

	if m.UpdatedAt != nil {

		if all {
			switch v := interface{}(m.GetUpdatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "UpdatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentShortValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.DeletedAt != nil {

		if all {
			switch v := interface{}(m.GetDeletedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "DeletedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentShortValidationError{
					field:  "DeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.CategoryId != nil {
		// no validation rules for CategoryId
	}

	if m.Category != nil {

		if all {
			switch v := interface{}(m.GetCategory()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "Category",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "Category",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentShortValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				}
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
					errors = append(errors, DocumentShortValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentShortValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentShortValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DocumentShortMultiError(errors)
	}

	return nil
}

// DocumentShortMultiError is an error wrapping multiple validation errors
// returned by DocumentShort.ValidateAll() if the designated constraints
// aren't met.
type DocumentShortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentShortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentShortMultiError) AllErrors() []error { return m }

// DocumentShortValidationError is the validation error returned by
// DocumentShort.Validate if the designated constraints aren't met.
type DocumentShortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentShortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentShortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentShortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentShortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentShortValidationError) ErrorName() string { return "DocumentShortValidationError" }

// Error satisfies the builtin error interface
func (e DocumentShortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentShort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentShortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentShortValidationError{}

// Validate checks the field values on DocumentReference with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DocumentReference) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DocumentReference with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DocumentReferenceMultiError, or nil if none found.
func (m *DocumentReference) ValidateAll() error {
	return m.validate(true)
}

func (m *DocumentReference) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SourceDocumentId

	if _, ok := DocReference_name[int32(m.GetReference())]; !ok {
		err := DocumentReferenceValidationError{
			field:  "Reference",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for TargetDocumentId

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.CreatedAt != nil {

		if all {
			switch v := interface{}(m.GetCreatedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "CreatedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentReferenceValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.SourceDocument != nil {

		if all {
			switch v := interface{}(m.GetSourceDocument()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "SourceDocument",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "SourceDocument",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSourceDocument()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentReferenceValidationError{
					field:  "SourceDocument",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.TargetDocument != nil {

		if all {
			switch v := interface{}(m.GetTargetDocument()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "TargetDocument",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "TargetDocument",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTargetDocument()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentReferenceValidationError{
					field:  "TargetDocument",
					reason: "embedded message failed validation",
					cause:  err,
				}
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
					errors = append(errors, DocumentReferenceValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentReferenceValidationError{
						field:  "Creator",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreator()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentReferenceValidationError{
					field:  "Creator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DocumentReferenceMultiError(errors)
	}

	return nil
}

// DocumentReferenceMultiError is an error wrapping multiple validation errors
// returned by DocumentReference.ValidateAll() if the designated constraints
// aren't met.
type DocumentReferenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DocumentReferenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DocumentReferenceMultiError) AllErrors() []error { return m }

// DocumentReferenceValidationError is the validation error returned by
// DocumentReference.Validate if the designated constraints aren't met.
type DocumentReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocumentReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocumentReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocumentReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocumentReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocumentReferenceValidationError) ErrorName() string {
	return "DocumentReferenceValidationError"
}

// Error satisfies the builtin error interface
func (e DocumentReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocumentReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocumentReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocumentReferenceValidationError{}

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

	// no validation rules for DocumentId

	if m.GetSourceUserId() <= 0 {
		err := DocumentRelationValidationError{
			field:  "SourceUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := DocRelation_name[int32(m.GetRelation())]; !ok {
		err := DocumentRelationValidationError{
			field:  "Relation",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTargetUserId() <= 0 {
		err := DocumentRelationValidationError{
			field:  "TargetUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.CreatedAt != nil {

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

	}

	if m.Document != nil {

		if all {
			switch v := interface{}(m.GetDocument()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentRelationValidationError{
						field:  "Document",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentRelationValidationError{
						field:  "Document",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDocument()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentRelationValidationError{
					field:  "Document",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.SourceUser != nil {

		if all {
			switch v := interface{}(m.GetSourceUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentRelationValidationError{
						field:  "SourceUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentRelationValidationError{
						field:  "SourceUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSourceUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentRelationValidationError{
					field:  "SourceUser",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.TargetUser != nil {

		if all {
			switch v := interface{}(m.GetTargetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DocumentRelationValidationError{
						field:  "TargetUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DocumentRelationValidationError{
						field:  "TargetUser",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTargetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DocumentRelationValidationError{
					field:  "TargetUser",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

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
