// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/compression/gzip/compressor/v3/gzip.proto

package envoy_extensions_compression_gzip_compressor_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Gzip with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Gzip) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetMemoryLevel(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 9 {
			return GzipValidationError{
				field:  "MemoryLevel",
				reason: "value must be inside range [1, 9]",
			}
		}

	}

	if _, ok := Gzip_CompressionLevel_name[int32(m.GetCompressionLevel())]; !ok {
		return GzipValidationError{
			field:  "CompressionLevel",
			reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := Gzip_CompressionStrategy_name[int32(m.GetCompressionStrategy())]; !ok {
		return GzipValidationError{
			field:  "CompressionStrategy",
			reason: "value must be one of the defined enum values",
		}
	}

	if wrapper := m.GetWindowBits(); wrapper != nil {

		if val := wrapper.GetValue(); val < 9 || val > 15 {
			return GzipValidationError{
				field:  "WindowBits",
				reason: "value must be inside range [9, 15]",
			}
		}

	}

	if wrapper := m.GetChunkSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 4096 || val > 65536 {
			return GzipValidationError{
				field:  "ChunkSize",
				reason: "value must be inside range [4096, 65536]",
			}
		}

	}

	return nil
}

// GzipValidationError is the validation error returned by Gzip.Validate if the
// designated constraints aren't met.
type GzipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GzipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GzipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GzipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GzipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GzipValidationError) ErrorName() string { return "GzipValidationError" }

// Error satisfies the builtin error interface
func (e GzipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGzip.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GzipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GzipValidationError{}
