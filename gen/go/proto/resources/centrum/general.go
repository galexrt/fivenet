package centrum

import (
	"database/sql/driver"
	"slices"

	"github.com/galexrt/fivenet/pkg/utils/protoutils"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	UnitAttributeStatic               = "static"
	UnitAttributeNoDispatchAutoAssign = "no_dispatch_auto_assign"

	DispatchAttributeMultiple  = "multiple"
	DispatchAttributeDuplicate = "duplicate"
	DispatchAttributeTooOld    = "too_old"
)

func (x *Attributes) Scan(value any) error {
	switch t := value.(type) {
	case string:
		return protojson.Unmarshal([]byte(t), x)
	case []byte:
		return protojson.Unmarshal(t, x)
	}
	return nil
}

// Scan implements driver.Valuer for protobuf Attributes.
func (x *Attributes) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}

	out, err := protoutils.Marshal(x)
	return string(out), err
}

func (x *Attributes) Has(attribute string) bool {
	if len(x.List) == 0 {
		return false
	}

	return slices.Contains(x.List, attribute)
}

func (x *Attributes) Add(attribute string) bool {
	if x.Has(attribute) {
		return false
	}

	if x.List == nil {
		x.List = []string{attribute}
	} else {
		x.List = append(x.List, attribute)
	}

	return true
}

func (x *Attributes) Remove(attribute string) bool {
	if !x.Has(attribute) {
		return false
	}

	x.List = slices.DeleteFunc(x.List, func(item string) bool {
		return item == attribute
	})

	return true
}
