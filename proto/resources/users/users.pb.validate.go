// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/users/users.proto

package users

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

// Validate checks the field values on UserShort with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserShort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserShort with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserShortMultiError, or nil
// if none found.
func (m *UserShort) ValidateAll() error {
	return m.validate(true)
}

func (m *UserShort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := UserShortValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetIdentifier()) != 46 {
		err := UserShortValidationError{
			field:  "Identifier",
			reason: "value length must be 46 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetJob()) > 50 {
		err := UserShortValidationError{
			field:  "Job",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJobLabel()) > 50 {
		err := UserShortValidationError{
			field:  "JobLabel",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetJobGrade() <= 0 {
		err := UserShortValidationError{
			field:  "JobGrade",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJobGradeLabel()) > 50 {
		err := UserShortValidationError{
			field:  "JobGradeLabel",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetFirstname()); l < 1 || l > 50 {
		err := UserShortValidationError{
			field:  "Firstname",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetLastname()); l < 1 || l > 50 {
		err := UserShortValidationError{
			field:  "Lastname",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserShortMultiError(errors)
	}

	return nil
}

// UserShortMultiError is an error wrapping multiple validation errors returned
// by UserShort.ValidateAll() if the designated constraints aren't met.
type UserShortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserShortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserShortMultiError) AllErrors() []error { return m }

// UserShortValidationError is the validation error returned by
// UserShort.Validate if the designated constraints aren't met.
type UserShortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserShortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserShortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserShortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserShortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserShortValidationError) ErrorName() string { return "UserShortValidationError" }

// Error satisfies the builtin error interface
func (e UserShortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserShort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserShortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserShortValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := UserValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetIdentifier()) != 46 {
		err := UserValidationError{
			field:  "Identifier",
			reason: "value length must be 46 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetJob()) > 50 {
		err := UserValidationError{
			field:  "Job",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJobLabel()) > 50 {
		err := UserValidationError{
			field:  "JobLabel",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetJobGrade() <= 0 {
		err := UserValidationError{
			field:  "JobGrade",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetJobGradeLabel()) > 50 {
		err := UserValidationError{
			field:  "JobGradeLabel",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetFirstname()); l < 1 || l > 50 {
		err := UserValidationError{
			field:  "Firstname",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetLastname()); l < 1 || l > 50 {
		err := UserValidationError{
			field:  "Lastname",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDateofbirth()) != 10 {
		err := UserValidationError{
			field:  "Dateofbirth",
			reason: "value length must be 10 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if l := utf8.RuneCountInString(m.GetSex()); l < 1 || l > 2 {
		err := UserValidationError{
			field:  "Sex",
			reason: "value length must be between 1 and 2 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Height

	if utf8.RuneCountInString(m.GetPhoneNumber()) > 20 {
		err := UserValidationError{
			field:  "PhoneNumber",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetProps()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "Props",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "Props",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProps()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "Props",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetLicenses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserValidationError{
						field:  fmt.Sprintf("Licenses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserValidationError{
						field:  fmt.Sprintf("Licenses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  fmt.Sprintf("Licenses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Visum != nil {

		if m.GetVisum() < 0 {
			err := UserValidationError{
				field:  "Visum",
				reason: "value must be greater than or equal to 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Playtime != nil {

		if m.GetPlaytime() < 0 {
			err := UserValidationError{
				field:  "Playtime",
				reason: "value must be greater than or equal to 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on License with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *License) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on License with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LicenseMultiError, or nil if none found.
func (m *License) ValidateAll() error {
	return m.validate(true)
}

func (m *License) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetType()); l < 3 || l > 60 {
		err := LicenseValidationError{
			field:  "Type",
			reason: "value length must be between 3 and 60 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Label

	if len(errors) > 0 {
		return LicenseMultiError(errors)
	}

	return nil
}

// LicenseMultiError is an error wrapping multiple validation errors returned
// by License.ValidateAll() if the designated constraints aren't met.
type LicenseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LicenseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LicenseMultiError) AllErrors() []error { return m }

// LicenseValidationError is the validation error returned by License.Validate
// if the designated constraints aren't met.
type LicenseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseValidationError) ErrorName() string { return "LicenseValidationError" }

// Error satisfies the builtin error interface
func (e LicenseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicense.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseValidationError{}

// Validate checks the field values on UserProps with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserProps) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserProps with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserPropsMultiError, or nil
// if none found.
func (m *UserProps) ValidateAll() error {
	return m.validate(true)
}

func (m *UserProps) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := UserPropsValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Wanted != nil {
		// no validation rules for Wanted
	}

	if m.Job != nil {
		// no validation rules for Job
	}

	if len(errors) > 0 {
		return UserPropsMultiError(errors)
	}

	return nil
}

// UserPropsMultiError is an error wrapping multiple validation errors returned
// by UserProps.ValidateAll() if the designated constraints aren't met.
type UserPropsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserPropsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserPropsMultiError) AllErrors() []error { return m }

// UserPropsValidationError is the validation error returned by
// UserProps.Validate if the designated constraints aren't met.
type UserPropsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserPropsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserPropsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserPropsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserPropsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserPropsValidationError) ErrorName() string { return "UserPropsValidationError" }

// Error satisfies the builtin error interface
func (e UserPropsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserProps.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserPropsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserPropsValidationError{}

// Validate checks the field values on UserActivity with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserActivity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserActivity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserActivityMultiError, or
// nil if none found.
func (m *UserActivity) ValidateAll() error {
	return m.validate(true)
}

func (m *UserActivity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if _, ok := USER_ACTIVITY_TYPE_name[int32(m.GetType())]; !ok {
		err := UserActivityValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
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
				errors = append(errors, UserActivityValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserActivityValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserActivityValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSourceUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserActivityValidationError{
					field:  "SourceUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserActivityValidationError{
					field:  "SourceUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSourceUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserActivityValidationError{
				field:  "SourceUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTargetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserActivityValidationError{
					field:  "TargetUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserActivityValidationError{
					field:  "TargetUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTargetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserActivityValidationError{
				field:  "TargetUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetKey()); l < 1 || l > 64 {
		err := UserActivityValidationError{
			field:  "Key",
			reason: "value length must be between 1 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetOldValue()) > 256 {
		err := UserActivityValidationError{
			field:  "OldValue",
			reason: "value length must be at most 256 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetNewValue()) > 256 {
		err := UserActivityValidationError{
			field:  "NewValue",
			reason: "value length must be at most 256 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserActivityMultiError(errors)
	}

	return nil
}

// UserActivityMultiError is an error wrapping multiple validation errors
// returned by UserActivity.ValidateAll() if the designated constraints aren't met.
type UserActivityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserActivityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserActivityMultiError) AllErrors() []error { return m }

// UserActivityValidationError is the validation error returned by
// UserActivity.Validate if the designated constraints aren't met.
type UserActivityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserActivityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserActivityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserActivityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserActivityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserActivityValidationError) ErrorName() string { return "UserActivityValidationError" }

// Error satisfies the builtin error interface
func (e UserActivityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserActivity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserActivityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserActivityValidationError{}
