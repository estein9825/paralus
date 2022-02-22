
// Code generated by go generate; DO NOT EDIT.
package sentry

import (
	driver "database/sql/driver"
	bytes "bytes"
)

// Scan converts database string to BootstrapAgentState
func (e *BootstrapAgentState) Scan(value interface{}) error {
	s := value.([]byte)
	*e = BootstrapAgentState(BootstrapAgentState_value[string(s)])
	return nil
}

// Value converts BootstrapAgentState into database string
func (e BootstrapAgentState) Value() (driver.Value, error) {
	return BootstrapAgentState_name[int32(e)], nil
}

// MarshalJSON converts BootstrapAgentState to JSON
func (e BootstrapAgentState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("\"")
	buffer.WriteString(e.String())
	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

// UnmarshalJSON converts BootstrapAgentState from JSON
func (e *BootstrapAgentState) UnmarshalJSON(b []byte) error {
	if b != nil {
		*e = BootstrapAgentState(BootstrapAgentState_value[string(b[1:len(b)-1])])
	}
	return nil
}

// implement proto enum interface
func (e BootstrapAgentState) IsEnum()  {
}

