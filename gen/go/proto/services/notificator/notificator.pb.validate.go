// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/notificator/notificator.proto

package notificator

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

// Validate checks the field values on GetNotificationsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNotificationsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNotificationsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNotificationsRequestMultiError, or nil if none found.
func (m *GetNotificationsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNotificationsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPagination() == nil {
		err := GetNotificationsRequestValidationError{
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
				errors = append(errors, GetNotificationsRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNotificationsRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNotificationsRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.IncludeRead != nil {
		// no validation rules for IncludeRead
	}

	if len(errors) > 0 {
		return GetNotificationsRequestMultiError(errors)
	}

	return nil
}

// GetNotificationsRequestMultiError is an error wrapping multiple validation
// errors returned by GetNotificationsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetNotificationsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNotificationsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNotificationsRequestMultiError) AllErrors() []error { return m }

// GetNotificationsRequestValidationError is the validation error returned by
// GetNotificationsRequest.Validate if the designated constraints aren't met.
type GetNotificationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNotificationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNotificationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNotificationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNotificationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNotificationsRequestValidationError) ErrorName() string {
	return "GetNotificationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNotificationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNotificationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNotificationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNotificationsRequestValidationError{}

// Validate checks the field values on GetNotificationsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNotificationsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNotificationsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNotificationsResponseMultiError, or nil if none found.
func (m *GetNotificationsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNotificationsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNotificationsResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNotificationsResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNotificationsResponseValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetNotificationsResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetNotificationsResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetNotificationsResponseValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetNotificationsResponseMultiError(errors)
	}

	return nil
}

// GetNotificationsResponseMultiError is an error wrapping multiple validation
// errors returned by GetNotificationsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetNotificationsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNotificationsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNotificationsResponseMultiError) AllErrors() []error { return m }

// GetNotificationsResponseValidationError is the validation error returned by
// GetNotificationsResponse.Validate if the designated constraints aren't met.
type GetNotificationsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNotificationsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNotificationsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNotificationsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNotificationsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNotificationsResponseValidationError) ErrorName() string {
	return "GetNotificationsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetNotificationsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNotificationsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNotificationsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNotificationsResponseValidationError{}

// Validate checks the field values on ReadNotificationsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadNotificationsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadNotificationsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadNotificationsRequestMultiError, or nil if none found.
func (m *ReadNotificationsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadNotificationsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := len(m.GetIds()); l < 1 || l > 20 {
		err := ReadNotificationsRequestValidationError{
			field:  "Ids",
			reason: "value must contain between 1 and 20 items, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ReadNotificationsRequestMultiError(errors)
	}

	return nil
}

// ReadNotificationsRequestMultiError is an error wrapping multiple validation
// errors returned by ReadNotificationsRequest.ValidateAll() if the designated
// constraints aren't met.
type ReadNotificationsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadNotificationsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadNotificationsRequestMultiError) AllErrors() []error { return m }

// ReadNotificationsRequestValidationError is the validation error returned by
// ReadNotificationsRequest.Validate if the designated constraints aren't met.
type ReadNotificationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadNotificationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadNotificationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadNotificationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadNotificationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadNotificationsRequestValidationError) ErrorName() string {
	return "ReadNotificationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadNotificationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadNotificationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadNotificationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadNotificationsRequestValidationError{}

// Validate checks the field values on ReadNotificationsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadNotificationsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadNotificationsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadNotificationsResponseMultiError, or nil if none found.
func (m *ReadNotificationsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadNotificationsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReadNotificationsResponseMultiError(errors)
	}

	return nil
}

// ReadNotificationsResponseMultiError is an error wrapping multiple validation
// errors returned by ReadNotificationsResponse.ValidateAll() if the
// designated constraints aren't met.
type ReadNotificationsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadNotificationsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadNotificationsResponseMultiError) AllErrors() []error { return m }

// ReadNotificationsResponseValidationError is the validation error returned by
// ReadNotificationsResponse.Validate if the designated constraints aren't met.
type ReadNotificationsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadNotificationsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadNotificationsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadNotificationsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadNotificationsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadNotificationsResponseValidationError) ErrorName() string {
	return "ReadNotificationsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadNotificationsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadNotificationsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadNotificationsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadNotificationsResponseValidationError{}

// Validate checks the field values on StreamRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamRequestMultiError, or
// nil if none found.
func (m *StreamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LastId

	if len(errors) > 0 {
		return StreamRequestMultiError(errors)
	}

	return nil
}

// StreamRequestMultiError is an error wrapping multiple validation errors
// returned by StreamRequest.ValidateAll() if the designated constraints
// aren't met.
type StreamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamRequestMultiError) AllErrors() []error { return m }

// StreamRequestValidationError is the validation error returned by
// StreamRequest.Validate if the designated constraints aren't met.
type StreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamRequestValidationError) ErrorName() string { return "StreamRequestValidationError" }

// Error satisfies the builtin error interface
func (e StreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamRequestValidationError{}

// Validate checks the field values on StreamResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamResponseMultiError,
// or nil if none found.
func (m *StreamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LastId

	// no validation rules for RestartStream

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  fmt.Sprintf("Notifications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamResponseValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Token != nil {

		if all {
			switch v := interface{}(m.GetToken()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  "Token",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamResponseValidationError{
						field:  "Token",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamResponseValidationError{
					field:  "Token",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StreamResponseMultiError(errors)
	}

	return nil
}

// StreamResponseMultiError is an error wrapping multiple validation errors
// returned by StreamResponse.ValidateAll() if the designated constraints
// aren't met.
type StreamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamResponseMultiError) AllErrors() []error { return m }

// StreamResponseValidationError is the validation error returned by
// StreamResponse.Validate if the designated constraints aren't met.
type StreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamResponseValidationError) ErrorName() string { return "StreamResponseValidationError" }

// Error satisfies the builtin error interface
func (e StreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamResponseValidationError{}

// Validate checks the field values on TokenUpdate with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TokenUpdate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TokenUpdate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TokenUpdateMultiError, or
// nil if none found.
func (m *TokenUpdate) ValidateAll() error {
	return m.validate(true)
}

func (m *TokenUpdate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetExpires()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TokenUpdateValidationError{
					field:  "Expires",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TokenUpdateValidationError{
					field:  "Expires",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpires()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TokenUpdateValidationError{
				field:  "Expires",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.NewToken != nil {
		// no validation rules for NewToken
	}

	if m.UserInfo != nil {

		if all {
			switch v := interface{}(m.GetUserInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TokenUpdateValidationError{
						field:  "UserInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TokenUpdateValidationError{
						field:  "UserInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUserInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TokenUpdateValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.JobProps != nil {

		if all {
			switch v := interface{}(m.GetJobProps()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TokenUpdateValidationError{
						field:  "JobProps",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TokenUpdateValidationError{
						field:  "JobProps",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetJobProps()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TokenUpdateValidationError{
					field:  "JobProps",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TokenUpdateMultiError(errors)
	}

	return nil
}

// TokenUpdateMultiError is an error wrapping multiple validation errors
// returned by TokenUpdate.ValidateAll() if the designated constraints aren't met.
type TokenUpdateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenUpdateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenUpdateMultiError) AllErrors() []error { return m }

// TokenUpdateValidationError is the validation error returned by
// TokenUpdate.Validate if the designated constraints aren't met.
type TokenUpdateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenUpdateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenUpdateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenUpdateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenUpdateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenUpdateValidationError) ErrorName() string { return "TokenUpdateValidationError" }

// Error satisfies the builtin error interface
func (e TokenUpdateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenUpdate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenUpdateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenUpdateValidationError{}
