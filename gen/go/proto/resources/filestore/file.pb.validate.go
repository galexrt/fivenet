// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/filestore/file.proto

package filestore

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

// Validate checks the field values on File with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *File) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on File with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FileMultiError, or nil if none found.
func (m *File) ValidateAll() error {
	return m.validate(true)
}

func (m *File) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetData()) > 2097152 {
		err := FileValidationError{
			field:  "Data",
			reason: "value length must be at most 2097152 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Url != nil {

		if utf8.RuneCountInString(m.GetUrl()) > 128 {
			err := FileValidationError{
				field:  "Url",
				reason: "value length must be at most 128 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Delete != nil {
		// no validation rules for Delete
	}

	if m.ContentType != nil {
		// no validation rules for ContentType
	}

	if m.Extension != nil {
		// no validation rules for Extension
	}

	if len(errors) > 0 {
		return FileMultiError(errors)
	}

	return nil
}

// FileMultiError is an error wrapping multiple validation errors returned by
// File.ValidateAll() if the designated constraints aren't met.
type FileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileMultiError) AllErrors() []error { return m }

// FileValidationError is the validation error returned by File.Validate if the
// designated constraints aren't met.
type FileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileValidationError) ErrorName() string { return "FileValidationError" }

// Error satisfies the builtin error interface
func (e FileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileValidationError{}

// Validate checks the field values on FileInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FileInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FileInfoMultiError, or nil
// if none found.
func (m *FileInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *FileInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Size

	// no validation rules for ContentType

	if m.LastModified != nil {

		if all {
			switch v := interface{}(m.GetLastModified()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FileInfoValidationError{
						field:  "LastModified",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FileInfoValidationError{
						field:  "LastModified",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLastModified()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FileInfoValidationError{
					field:  "LastModified",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FileInfoMultiError(errors)
	}

	return nil
}

// FileInfoMultiError is an error wrapping multiple validation errors returned
// by FileInfo.ValidateAll() if the designated constraints aren't met.
type FileInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileInfoMultiError) AllErrors() []error { return m }

// FileInfoValidationError is the validation error returned by
// FileInfo.Validate if the designated constraints aren't met.
type FileInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileInfoValidationError) ErrorName() string { return "FileInfoValidationError" }

// Error satisfies the builtin error interface
func (e FileInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileInfoValidationError{}