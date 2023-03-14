//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var ArpanetDocumentsReferences = newArpanetDocumentsReferencesTable("", "arpanet_documents_references", "")

type arpanetDocumentsReferencesTable struct {
	mysql.Table

	//Columns
	ID               mysql.ColumnInteger
	CreatedAt        mysql.ColumnTimestamp
	SourceDocumentID mysql.ColumnInteger
	Reference        mysql.ColumnInteger
	TargetDocumentID mysql.ColumnInteger
	CreatorID        mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetDocumentsReferencesTable struct {
	arpanetDocumentsReferencesTable

	NEW arpanetDocumentsReferencesTable
}

// AS creates new ArpanetDocumentsReferencesTable with assigned alias
func (a ArpanetDocumentsReferencesTable) AS(alias string) *ArpanetDocumentsReferencesTable {
	return newArpanetDocumentsReferencesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetDocumentsReferencesTable with assigned schema name
func (a ArpanetDocumentsReferencesTable) FromSchema(schemaName string) *ArpanetDocumentsReferencesTable {
	return newArpanetDocumentsReferencesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetDocumentsReferencesTable with assigned table prefix
func (a ArpanetDocumentsReferencesTable) WithPrefix(prefix string) *ArpanetDocumentsReferencesTable {
	return newArpanetDocumentsReferencesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetDocumentsReferencesTable with assigned table suffix
func (a ArpanetDocumentsReferencesTable) WithSuffix(suffix string) *ArpanetDocumentsReferencesTable {
	return newArpanetDocumentsReferencesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetDocumentsReferencesTable(schemaName, tableName, alias string) *ArpanetDocumentsReferencesTable {
	return &ArpanetDocumentsReferencesTable{
		arpanetDocumentsReferencesTable: newArpanetDocumentsReferencesTableImpl(schemaName, tableName, alias),
		NEW:                             newArpanetDocumentsReferencesTableImpl("", "new", ""),
	}
}

func newArpanetDocumentsReferencesTableImpl(schemaName, tableName, alias string) arpanetDocumentsReferencesTable {
	var (
		IDColumn               = mysql.IntegerColumn("id")
		CreatedAtColumn        = mysql.TimestampColumn("created_at")
		SourceDocumentIDColumn = mysql.IntegerColumn("source_document_id")
		ReferenceColumn        = mysql.IntegerColumn("reference")
		TargetDocumentIDColumn = mysql.IntegerColumn("target_document_id")
		CreatorIDColumn        = mysql.IntegerColumn("creator_id")
		allColumns             = mysql.ColumnList{IDColumn, CreatedAtColumn, SourceDocumentIDColumn, ReferenceColumn, TargetDocumentIDColumn, CreatorIDColumn}
		mutableColumns         = mysql.ColumnList{CreatedAtColumn, SourceDocumentIDColumn, ReferenceColumn, TargetDocumentIDColumn, CreatorIDColumn}
	)

	return arpanetDocumentsReferencesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		CreatedAt:        CreatedAtColumn,
		SourceDocumentID: SourceDocumentIDColumn,
		Reference:        ReferenceColumn,
		TargetDocumentID: TargetDocumentIDColumn,
		CreatorID:        CreatorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
