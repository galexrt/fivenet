//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FivenetJobPermissions struct {
	Job          string `sql:"primary_key" json:"job"`
	PermissionID uint64 `sql:"primary_key" json:"permission_id"`
	Val          bool   `json:"val"`
}