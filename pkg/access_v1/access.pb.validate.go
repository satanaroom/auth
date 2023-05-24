// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: access.proto

package access_v1

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

// Validate checks the field values on EmptyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyResponseMultiError, or
// nil if none found.
func (m *EmptyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyResponseMultiError(errors)
	}

	return nil
}

// EmptyResponseMultiError is an error wrapping multiple validation errors
// returned by EmptyResponse.ValidateAll() if the designated constraints
// aren't met.
type EmptyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyResponseMultiError) AllErrors() []error { return m }

// EmptyResponseValidationError is the validation error returned by
// EmptyResponse.Validate if the designated constraints aren't met.
type EmptyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyResponseValidationError) ErrorName() string { return "EmptyResponseValidationError" }

// Error satisfies the builtin error interface
func (e EmptyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyResponseValidationError{}

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckRequestMultiError, or
// nil if none found.
func (m *CheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EndpointAddress

	if len(errors) > 0 {
		return CheckRequestMultiError(errors)
	}

	return nil
}

// CheckRequestMultiError is an error wrapping multiple validation errors
// returned by CheckRequest.ValidateAll() if the designated constraints aren't met.
type CheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckRequestMultiError) AllErrors() []error { return m }

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}
