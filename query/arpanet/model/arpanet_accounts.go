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

type ArpanetAccounts struct {
	ID        uint64     `sql:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Enabled   *bool      `json:"enabled"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	License   string     `json:"license"`
}
