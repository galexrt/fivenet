// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/auth/auth.proto

package auth

import "github.com/galexrt/arpanet/pkg/perms"

var PermsRemap = map[string]string{
	// Service: AuthService
	"AuthService/ChooseCharacter": "AuthService/GetCharacters",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

const (
	AuthServicePermKey = "AuthService"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: AuthService
		{
			Key:  AuthServicePermKey,
			Name: "GetCharacters",
		},
		{
			Key:  AuthServicePermKey,
			Name: "GetCharacters",
		},
		{
			Key:  AuthServicePermKey,
			Name: "SetJob",
		},
	})
}
