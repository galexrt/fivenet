// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/dmv/vehicles.proto

package dmv

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

// Validate checks the field values on ListVehiclesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListVehiclesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListVehiclesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListVehiclesRequestMultiError, or nil if none found.
func (m *ListVehiclesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListVehiclesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPagination() == nil {
		err := ListVehiclesRequestValidationError{
			field:  "Pagination",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListVehiclesRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListVehiclesRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListVehiclesRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetOrderBy()) > 3 {
		err := ListVehiclesRequestValidationError{
			field:  "OrderBy",
			reason: "value must contain no more than 3 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetOrderBy() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListVehiclesRequestValidationError{
						field:  fmt.Sprintf("OrderBy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListVehiclesRequestValidationError{
						field:  fmt.Sprintf("OrderBy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListVehiclesRequestValidationError{
					field:  fmt.Sprintf("OrderBy[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.LicensePlate != nil {

		if utf8.RuneCountInString(m.GetLicensePlate()) > 32 {
			err := ListVehiclesRequestValidationError{
				field:  "LicensePlate",
				reason: "value length must be at most 32 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Model != nil {

		if utf8.RuneCountInString(m.GetModel()) > 32 {
			err := ListVehiclesRequestValidationError{
				field:  "Model",
				reason: "value length must be at most 32 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if len(errors) > 0 {
		return ListVehiclesRequestMultiError(errors)
	}

	return nil
}

// ListVehiclesRequestMultiError is an error wrapping multiple validation
// errors returned by ListVehiclesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListVehiclesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListVehiclesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListVehiclesRequestMultiError) AllErrors() []error { return m }

// ListVehiclesRequestValidationError is the validation error returned by
// ListVehiclesRequest.Validate if the designated constraints aren't met.
type ListVehiclesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListVehiclesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListVehiclesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListVehiclesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListVehiclesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListVehiclesRequestValidationError) ErrorName() string {
	return "ListVehiclesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListVehiclesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListVehiclesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListVehiclesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListVehiclesRequestValidationError{}

// Validate checks the field values on ListVehiclesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListVehiclesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListVehiclesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListVehiclesResponseMultiError, or nil if none found.
func (m *ListVehiclesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListVehiclesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListVehiclesResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListVehiclesResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListVehiclesResponseValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetVehicles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListVehiclesResponseValidationError{
						field:  fmt.Sprintf("Vehicles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListVehiclesResponseValidationError{
						field:  fmt.Sprintf("Vehicles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListVehiclesResponseValidationError{
					field:  fmt.Sprintf("Vehicles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListVehiclesResponseMultiError(errors)
	}

	return nil
}

// ListVehiclesResponseMultiError is an error wrapping multiple validation
// errors returned by ListVehiclesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListVehiclesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListVehiclesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListVehiclesResponseMultiError) AllErrors() []error { return m }

// ListVehiclesResponseValidationError is the validation error returned by
// ListVehiclesResponse.Validate if the designated constraints aren't met.
type ListVehiclesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListVehiclesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListVehiclesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListVehiclesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListVehiclesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListVehiclesResponseValidationError) ErrorName() string {
	return "ListVehiclesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListVehiclesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListVehiclesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListVehiclesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListVehiclesResponseValidationError{}
