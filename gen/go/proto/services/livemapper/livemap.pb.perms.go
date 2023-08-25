// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/livemapper/livemap.proto

package livemapper

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	LivemapperServicePerm perms.Category = "LivemapperService"

	LivemapperServiceCreateOrUpdateMarkerPerm            perms.Name = "CreateOrUpdateMarker"
	LivemapperServiceCreateOrUpdateMarkerAccessPermField perms.Key  = "Access"
	LivemapperServiceDeleteMarkerPerm                    perms.Name = "DeleteMarker"
	LivemapperServiceDeleteMarkerAccessPermField         perms.Key  = "Access"
	LivemapperServiceStreamPerm                          perms.Name = "Stream"
	LivemapperServiceStreamMarkersPermField              perms.Key  = "Markers"
	LivemapperServiceStreamPlayersPermField              perms.Key  = "Players"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: LivemapperService
		{
			Category: LivemapperServicePerm,
			Name:     LivemapperServiceCreateOrUpdateMarkerPerm,
			Attrs: []perms.Attr{
				{
					Key:           LivemapperServiceCreateOrUpdateMarkerAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "Lower_Rank", "Same_Rank"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: LivemapperServicePerm,
			Name:     LivemapperServiceDeleteMarkerPerm,
			Attrs: []perms.Attr{
				{
					Key:           LivemapperServiceDeleteMarkerAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "Lower_Rank", "Same_Rank"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: LivemapperServicePerm,
			Name:     LivemapperServiceStreamPerm,
			Attrs: []perms.Attr{
				{
					Key:         LivemapperServiceStreamMarkersPermField,
					Type:        permissions.JobListAttributeType,
					ValidValues: "config.Game.Livemap.Jobs",
				},
				{
					Key:  LivemapperServiceStreamPlayersPermField,
					Type: permissions.JobGradeListAttributeType,
				},
			},
		},
	})
}
