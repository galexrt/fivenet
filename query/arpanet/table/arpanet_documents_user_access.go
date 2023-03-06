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

var ArpanetDocumentsUserAccess = newArpanetDocumentsUserAccessTable("arpanet", "arpanet_documents_user_access", "")

type arpanetDocumentsUserAccessTable struct {
	mysql.Table

	//Columns
	ID         mysql.ColumnInteger
	CreatedAt  mysql.ColumnTimestamp
	UpdatedAt  mysql.ColumnTimestamp
	DocumentID mysql.ColumnInteger
	Identifier mysql.ColumnString
	Access     mysql.ColumnString
	UserID     mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetDocumentsUserAccessTable struct {
	arpanetDocumentsUserAccessTable

	NEW arpanetDocumentsUserAccessTable
}

// AS creates new ArpanetDocumentsUserAccessTable with assigned alias
func (a ArpanetDocumentsUserAccessTable) AS(alias string) *ArpanetDocumentsUserAccessTable {
	return newArpanetDocumentsUserAccessTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetDocumentsUserAccessTable with assigned schema name
func (a ArpanetDocumentsUserAccessTable) FromSchema(schemaName string) *ArpanetDocumentsUserAccessTable {
	return newArpanetDocumentsUserAccessTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetDocumentsUserAccessTable with assigned table prefix
func (a ArpanetDocumentsUserAccessTable) WithPrefix(prefix string) *ArpanetDocumentsUserAccessTable {
	return newArpanetDocumentsUserAccessTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetDocumentsUserAccessTable with assigned table suffix
func (a ArpanetDocumentsUserAccessTable) WithSuffix(suffix string) *ArpanetDocumentsUserAccessTable {
	return newArpanetDocumentsUserAccessTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetDocumentsUserAccessTable(schemaName, tableName, alias string) *ArpanetDocumentsUserAccessTable {
	return &ArpanetDocumentsUserAccessTable{
		arpanetDocumentsUserAccessTable: newArpanetDocumentsUserAccessTableImpl(schemaName, tableName, alias),
		NEW:                             newArpanetDocumentsUserAccessTableImpl("", "new", ""),
	}
}

func newArpanetDocumentsUserAccessTableImpl(schemaName, tableName, alias string) arpanetDocumentsUserAccessTable {
	var (
		IDColumn         = mysql.IntegerColumn("id")
		CreatedAtColumn  = mysql.TimestampColumn("created_at")
		UpdatedAtColumn  = mysql.TimestampColumn("updated_at")
		DocumentIDColumn = mysql.IntegerColumn("document_id")
		IdentifierColumn = mysql.StringColumn("identifier")
		AccessColumn     = mysql.StringColumn("access")
		UserIDColumn     = mysql.IntegerColumn("user_id")
		allColumns       = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DocumentIDColumn, IdentifierColumn, AccessColumn, UserIDColumn}
		mutableColumns   = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DocumentIDColumn, IdentifierColumn, AccessColumn, UserIDColumn}
	)

	return arpanetDocumentsUserAccessTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		DocumentID: DocumentIDColumn,
		Identifier: IdentifierColumn,
		Access:     AccessColumn,
		UserID:     UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
