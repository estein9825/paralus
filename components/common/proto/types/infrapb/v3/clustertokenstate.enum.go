
// Code generated by go generate; DO NOT EDIT.
package infrav3

import (
	driver "database/sql/driver"
	bytes "bytes"
)

// Scan converts database string to ClusterTokenState
func (e *ClusterTokenState) Scan(value interface{}) error {
	s := value.([]byte)
	*e = ClusterTokenState(ClusterTokenState_value[string(s)])
	return nil
}

// Value converts ClusterTokenState into database string
func (e ClusterTokenState) Value() (driver.Value, error) {
	return ClusterTokenState_name[int32(e)], nil
}

// MarshalJSON converts ClusterTokenState to JSON
func (e ClusterTokenState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("\"")
	buffer.WriteString(e.String())
	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

// UnmarshalJSON converts ClusterTokenState from JSON
func (e *ClusterTokenState) UnmarshalJSON(b []byte) error {
	if b != nil {
		*e = ClusterTokenState(ClusterTokenState_value[string(b[1:len(b)-1])])
	}
	return nil
}

// implement proto enum interface
func (e ClusterTokenState) IsEnum()  {
}

