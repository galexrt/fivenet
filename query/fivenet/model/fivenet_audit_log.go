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

type FivenetAuditLog struct {
	ID        uint64     `sql:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UserID    int32      `json:"user_id"`
	Service   string     `json:"service"`
	Method    string     `json:"method"`
	State     int16      `json:"state"`
	Data      *string    `json:"data"`
}
