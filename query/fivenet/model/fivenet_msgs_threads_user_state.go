//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FivenetMsgsThreadsUserState struct {
	ThreadID  uint64  `sql:"primary_key" json:"thread_id"`
	UserID    int32   `sql:"primary_key" json:"user_id"`
	LastRead  *uint64 `json:"last_read"`
	Important *bool   `json:"important"`
	Favorite  *bool   `json:"favorite"`
	Muted     *bool   `json:"muted"`
}