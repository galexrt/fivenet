// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/centrum/centrum.proto

package centrum

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	CentrumServicePerm perms.Category = "CentrumService"
	UnitServicePerm    perms.Category = "UnitService"

	CentrumServiceCreateDispatchPerm     perms.Name = "CreateDispatch"
	CentrumServiceStreamPerm             perms.Name = "Stream"
	UnitServiceAssignUnitPerm            perms.Name = "AssignUnit"
	UnitServiceAssignUnitAccessPermField perms.Key  = "Access"
	UnitServiceCreateUnitPerm            perms.Name = "CreateUnit"
	UnitServiceDeleteUnitPerm            perms.Name = "DeleteUnit"
	UnitServiceListUnitsPerm             perms.Name = "ListUnits"
	UnitServiceStreamUnitsPerm           perms.Name = "StreamUnits"
)

var PermsRemap = map[string]string{
	// Service: UnitService
	"UnitService/UpdateUnit": "UnitService/CreateUnit",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: CentrumService
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceCreateDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceStreamPerm,
			Attrs:    []perms.Attr{},
		}, // Service: UnitService
		{
			Category: UnitServicePerm,
			Name:     UnitServiceAssignUnitPerm,
			Attrs: []perms.Attr{
				{
					Key:           UnitServiceAssignUnitAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "Lower_Rank", "Same_Rank"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: UnitServicePerm,
			Name:     UnitServiceCreateUnitPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: UnitServicePerm,
			Name:     UnitServiceDeleteUnitPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: UnitServicePerm,
			Name:     UnitServiceListUnitsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: UnitServicePerm,
			Name:     UnitServiceStreamUnitsPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
