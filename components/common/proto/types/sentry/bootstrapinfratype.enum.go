
// Code generated by go generate; DO NOT EDIT.
package sentry

import (
	driver "database/sql/driver"
	bytes "bytes"
)

// Scan converts database string to BootstrapInfraType
func (e *BootstrapInfraType) Scan(value interface{}) error {
	s := value.([]byte)
	*e = BootstrapInfraType(BootstrapInfraType_value[string(s)])
	return nil
}

// Value converts BootstrapInfraType into database string
func (e BootstrapInfraType) Value() (driver.Value, error) {
	return BootstrapInfraType_name[int32(e)], nil
}

// MarshalJSON converts BootstrapInfraType to JSON
func (e BootstrapInfraType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("\"")
	buffer.WriteString(e.String())
	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

// UnmarshalJSON converts BootstrapInfraType from JSON
func (e *BootstrapInfraType) UnmarshalJSON(b []byte) error {
	if b != nil {
		*e = BootstrapInfraType(BootstrapInfraType_value[string(b[1:len(b)-1])])
	}
	return nil
}

// implement proto enum interface
func (e BootstrapInfraType) IsEnum()  {
}

