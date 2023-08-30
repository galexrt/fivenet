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

var FivenetDocumentsTemplates = newFivenetDocumentsTemplatesTable("", "fivenet_documents_templates", "")

type fivenetDocumentsTemplatesTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	UpdatedAt    mysql.ColumnTimestamp
	DeletedAt    mysql.ColumnTimestamp
	Weight       mysql.ColumnInteger
	CategoryID   mysql.ColumnInteger
	Title        mysql.ColumnString
	Description  mysql.ColumnString
	ContentTitle mysql.ColumnString
	Content      mysql.ColumnString
	State        mysql.ColumnString
	Access       mysql.ColumnString
	Schema       mysql.ColumnString
	CreatorID    mysql.ColumnInteger
	CreatorJob   mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetDocumentsTemplatesTable struct {
	fivenetDocumentsTemplatesTable

	NEW fivenetDocumentsTemplatesTable
}

// AS creates new FivenetDocumentsTemplatesTable with assigned alias
func (a FivenetDocumentsTemplatesTable) AS(alias string) *FivenetDocumentsTemplatesTable {
	return newFivenetDocumentsTemplatesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetDocumentsTemplatesTable with assigned schema name
func (a FivenetDocumentsTemplatesTable) FromSchema(schemaName string) *FivenetDocumentsTemplatesTable {
	return newFivenetDocumentsTemplatesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetDocumentsTemplatesTable with assigned table prefix
func (a FivenetDocumentsTemplatesTable) WithPrefix(prefix string) *FivenetDocumentsTemplatesTable {
	return newFivenetDocumentsTemplatesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetDocumentsTemplatesTable with assigned table suffix
func (a FivenetDocumentsTemplatesTable) WithSuffix(suffix string) *FivenetDocumentsTemplatesTable {
	return newFivenetDocumentsTemplatesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetDocumentsTemplatesTable(schemaName, tableName, alias string) *FivenetDocumentsTemplatesTable {
	return &FivenetDocumentsTemplatesTable{
		fivenetDocumentsTemplatesTable: newFivenetDocumentsTemplatesTableImpl(schemaName, tableName, alias),
		NEW:                            newFivenetDocumentsTemplatesTableImpl("", "new", ""),
	}
}

func newFivenetDocumentsTemplatesTableImpl(schemaName, tableName, alias string) fivenetDocumentsTemplatesTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		UpdatedAtColumn    = mysql.TimestampColumn("updated_at")
		DeletedAtColumn    = mysql.TimestampColumn("deleted_at")
		WeightColumn       = mysql.IntegerColumn("weight")
		CategoryIDColumn   = mysql.IntegerColumn("category_id")
		TitleColumn        = mysql.StringColumn("title")
		DescriptionColumn  = mysql.StringColumn("description")
		ContentTitleColumn = mysql.StringColumn("content_title")
		ContentColumn      = mysql.StringColumn("content")
		StateColumn        = mysql.StringColumn("state")
		AccessColumn       = mysql.StringColumn("access")
		SchemaColumn       = mysql.StringColumn("schema")
		CreatorIDColumn    = mysql.IntegerColumn("creator_id")
		CreatorJobColumn   = mysql.StringColumn("creator_job")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, WeightColumn, CategoryIDColumn, TitleColumn, DescriptionColumn, ContentTitleColumn, ContentColumn, StateColumn, AccessColumn, SchemaColumn, CreatorIDColumn, CreatorJobColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, WeightColumn, CategoryIDColumn, TitleColumn, DescriptionColumn, ContentTitleColumn, ContentColumn, StateColumn, AccessColumn, SchemaColumn, CreatorIDColumn, CreatorJobColumn}
	)

	return fivenetDocumentsTemplatesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		UpdatedAt:    UpdatedAtColumn,
		DeletedAt:    DeletedAtColumn,
		Weight:       WeightColumn,
		CategoryID:   CategoryIDColumn,
		Title:        TitleColumn,
		Description:  DescriptionColumn,
		ContentTitle: ContentTitleColumn,
		Content:      ContentColumn,
		State:        StateColumn,
		Access:       AccessColumn,
		Schema:       SchemaColumn,
		CreatorID:    CreatorIDColumn,
		CreatorJob:   CreatorJobColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
