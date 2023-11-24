package livemap

import (
	"database/sql/driver"

	"google.golang.org/protobuf/proto"
)

func (x *MarkerData) Scan(value any) error {
	switch t := value.(type) {
	case string:
		return proto.Unmarshal([]byte(t), x)
	case []byte:
		return proto.Unmarshal(t, x)
	}
	return nil
}

// Scan implements driver.Valuer for protobuf Data.
func (x *MarkerData) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}

	out, err := proto.Marshal(x)
	return out, err
}