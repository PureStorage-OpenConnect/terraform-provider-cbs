package models

import "fmt"

func (m *Error) String() string {
	if buf, err := m.MarshalBinary(); err != nil {
		return fmt.Sprintf("Error marshalling error msg: %+v", err)
	} else {
		return string(buf)
	}
}