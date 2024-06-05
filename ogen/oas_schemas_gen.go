// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// CreatePetsCreated is response for CreatePets operation.
type CreatePetsCreated struct{}

// Ref: #/components/schemas/Error
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Pet
type Pet struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	Tag  OptString `json:"tag"`
}

// GetID returns the value of ID.
func (s *Pet) GetID() int64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Pet) GetName() string {
	return s.Name
}

// GetTag returns the value of Tag.
func (s *Pet) GetTag() OptString {
	return s.Tag
}

// SetID sets the value of ID.
func (s *Pet) SetID(val int64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Pet) SetName(val string) {
	s.Name = val
}

// SetTag sets the value of Tag.
func (s *Pet) SetTag(val OptString) {
	s.Tag = val
}

type Pets []Pet

// PetsHeaders wraps Pets with response headers.
type PetsHeaders struct {
	XNext    OptString
	Response Pets
}

// GetXNext returns the value of XNext.
func (s *PetsHeaders) GetXNext() OptString {
	return s.XNext
}

// GetResponse returns the value of Response.
func (s *PetsHeaders) GetResponse() Pets {
	return s.Response
}

// SetXNext sets the value of XNext.
func (s *PetsHeaders) SetXNext(val OptString) {
	s.XNext = val
}

// SetResponse sets the value of Response.
func (s *PetsHeaders) SetResponse(val Pets) {
	s.Response = val
}
