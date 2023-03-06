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

var ArpanetDocumentsJobAccess = newArpanetDocumentsJobAccessTable("arpanet", "arpanet_documents_job_access", "")

type arpanetDocumentsJobAccessTable struct {
	mysql.Table

	//Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	UpdatedAt    mysql.ColumnTimestamp
	DocumentID   mysql.ColumnInteger
	Name         mysql.ColumnString
	MinimumGrade mysql.ColumnInteger
	Access       mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetDocumentsJobAccessTable struct {
	arpanetDocumentsJobAccessTable

	NEW arpanetDocumentsJobAccessTable
}

// AS creates new ArpanetDocumentsJobAccessTable with assigned alias
func (a ArpanetDocumentsJobAccessTable) AS(alias string) *ArpanetDocumentsJobAccessTable {
	return newArpanetDocumentsJobAccessTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetDocumentsJobAccessTable with assigned schema name
func (a ArpanetDocumentsJobAccessTable) FromSchema(schemaName string) *ArpanetDocumentsJobAccessTable {
	return newArpanetDocumentsJobAccessTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetDocumentsJobAccessTable with assigned table prefix
func (a ArpanetDocumentsJobAccessTable) WithPrefix(prefix string) *ArpanetDocumentsJobAccessTable {
	return newArpanetDocumentsJobAccessTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetDocumentsJobAccessTable with assigned table suffix
func (a ArpanetDocumentsJobAccessTable) WithSuffix(suffix string) *ArpanetDocumentsJobAccessTable {
	return newArpanetDocumentsJobAccessTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetDocumentsJobAccessTable(schemaName, tableName, alias string) *ArpanetDocumentsJobAccessTable {
	return &ArpanetDocumentsJobAccessTable{
		arpanetDocumentsJobAccessTable: newArpanetDocumentsJobAccessTableImpl(schemaName, tableName, alias),
		NEW:                            newArpanetDocumentsJobAccessTableImpl("", "new", ""),
	}
}

func newArpanetDocumentsJobAccessTableImpl(schemaName, tableName, alias string) arpanetDocumentsJobAccessTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		UpdatedAtColumn    = mysql.TimestampColumn("updated_at")
		DocumentIDColumn   = mysql.IntegerColumn("document_id")
		NameColumn         = mysql.StringColumn("name")
		MinimumGradeColumn = mysql.IntegerColumn("minimum_grade")
		AccessColumn       = mysql.StringColumn("access")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DocumentIDColumn, NameColumn, MinimumGradeColumn, AccessColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DocumentIDColumn, NameColumn, MinimumGradeColumn, AccessColumn}
	)

	return arpanetDocumentsJobAccessTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		UpdatedAt:    UpdatedAtColumn,
		DocumentID:   DocumentIDColumn,
		Name:         NameColumn,
		MinimumGrade: MinimumGradeColumn,
		Access:       AccessColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
