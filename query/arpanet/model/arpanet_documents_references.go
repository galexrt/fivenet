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

type ArpanetDocumentsReferences struct {
	ID               uint64     `sql:"primary_key" json:"id"`
	CreatedAt        *time.Time `json:"created_at"`
	SourceDocumentID uint64     `json:"source_document_id"`
	Reference        int16      `json:"reference"`
	TargetDocumentID uint64     `json:"target_document_id"`
	CreatorID        int32      `json:"creator_id"`
}
