//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type FivenetUserActivity struct {
	ID           uint64     `sql:"primary_key" json:"id"`
	CreatedAt    *time.Time `json:"created_at"`
	SourceUserID *int32     `json:"source_user_id"`
	TargetUserID int32      `json:"target_user_id"`
	Type         int16      `json:"type"`
	Key          string     `json:"key"`
	OldValue     *string    `json:"old_value"`
	NewValue     *string    `json:"new_value"`
	Reason       *string    `json:"reason"`
}
