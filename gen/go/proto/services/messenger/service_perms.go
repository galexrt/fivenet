// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/messenger/messenger.proto

package messenger

import (
	permkeys "github.com/fivenet-app/fivenet/gen/go/proto/services/messenger/perms"
	"github.com/fivenet-app/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: MessengerService
	"MessengerService/DeleteMessage":      "SuperUser",
	"MessengerService/GetThreadMessages":  "MessengerService/ListThreads",
	"MessengerService/SetThreadUserState": "MessengerService/ListThreads",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{

		// Service: MessengerService
		{
			Category: permkeys.MessengerServicePerm,
			Name:     permkeys.MessengerServiceCreateOrUpdateThreadPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.MessengerServicePerm,
			Name:     permkeys.MessengerServiceDeleteThreadPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.MessengerServicePerm,
			Name:     permkeys.MessengerServiceListThreadsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.MessengerServicePerm,
			Name:     permkeys.MessengerServicePostMessagePerm,
			Attrs:    []perms.Attr{},
		},
	})
}
