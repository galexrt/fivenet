//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type OwnedVehicles struct {
	Owner     *string `json:"owner"`
	Plate     string  `sql:"primary_key" json:"plate"`
	Model     string  `json:"model"`
	Vehicle   *string `json:"vehicle"`
	Type      string  `json:"type"`
	Stored    bool    `json:"stored"`
	Carseller *int32  `json:"carseller"`
	Owners    *string `json:"owners"`
	Storage   *string `json:"storage"`
	Trunk     *string `json:"trunk"`
	Glovebox  *string `json:"glovebox"`
}