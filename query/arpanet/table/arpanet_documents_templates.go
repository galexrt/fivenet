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

var ArpanetDocumentsTemplates = newArpanetDocumentsTemplatesTable("", "arpanet_documents_templates", "")

type arpanetDocumentsTemplatesTable struct {
	mysql.Table

	//Columns
	ID             mysql.ColumnInteger
	CreatedAt      mysql.ColumnTimestamp
	UpdatedAt      mysql.ColumnTimestamp
	Job            mysql.ColumnString
	JobGrade       mysql.ColumnInteger
	CategoryID     mysql.ColumnInteger
	Title          mysql.ColumnString
	Description    mysql.ColumnString
	ContentTitle   mysql.ColumnString
	Content        mysql.ColumnString
	AdditionalData mysql.ColumnString
	CreatorID      mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetDocumentsTemplatesTable struct {
	arpanetDocumentsTemplatesTable

	NEW arpanetDocumentsTemplatesTable
}

// AS creates new ArpanetDocumentsTemplatesTable with assigned alias
func (a ArpanetDocumentsTemplatesTable) AS(alias string) *ArpanetDocumentsTemplatesTable {
	return newArpanetDocumentsTemplatesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetDocumentsTemplatesTable with assigned schema name
func (a ArpanetDocumentsTemplatesTable) FromSchema(schemaName string) *ArpanetDocumentsTemplatesTable {
	return newArpanetDocumentsTemplatesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetDocumentsTemplatesTable with assigned table prefix
func (a ArpanetDocumentsTemplatesTable) WithPrefix(prefix string) *ArpanetDocumentsTemplatesTable {
	return newArpanetDocumentsTemplatesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetDocumentsTemplatesTable with assigned table suffix
func (a ArpanetDocumentsTemplatesTable) WithSuffix(suffix string) *ArpanetDocumentsTemplatesTable {
	return newArpanetDocumentsTemplatesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetDocumentsTemplatesTable(schemaName, tableName, alias string) *ArpanetDocumentsTemplatesTable {
	return &ArpanetDocumentsTemplatesTable{
		arpanetDocumentsTemplatesTable: newArpanetDocumentsTemplatesTableImpl(schemaName, tableName, alias),
		NEW:                            newArpanetDocumentsTemplatesTableImpl("", "new", ""),
	}
}

func newArpanetDocumentsTemplatesTableImpl(schemaName, tableName, alias string) arpanetDocumentsTemplatesTable {
	var (
		IDColumn             = mysql.IntegerColumn("id")
		CreatedAtColumn      = mysql.TimestampColumn("created_at")
		UpdatedAtColumn      = mysql.TimestampColumn("updated_at")
		JobColumn            = mysql.StringColumn("job")
		JobGradeColumn       = mysql.IntegerColumn("job_grade")
		CategoryIDColumn     = mysql.IntegerColumn("category_id")
		TitleColumn          = mysql.StringColumn("title")
		DescriptionColumn    = mysql.StringColumn("description")
		ContentTitleColumn   = mysql.StringColumn("content_title")
		ContentColumn        = mysql.StringColumn("content")
		AdditionalDataColumn = mysql.StringColumn("additional_data")
		CreatorIDColumn      = mysql.IntegerColumn("creator_id")
		allColumns           = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, JobColumn, JobGradeColumn, CategoryIDColumn, TitleColumn, DescriptionColumn, ContentTitleColumn, ContentColumn, AdditionalDataColumn, CreatorIDColumn}
		mutableColumns       = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, JobColumn, JobGradeColumn, CategoryIDColumn, TitleColumn, DescriptionColumn, ContentTitleColumn, ContentColumn, AdditionalDataColumn, CreatorIDColumn}
	)

	return arpanetDocumentsTemplatesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Job:            JobColumn,
		JobGrade:       JobGradeColumn,
		CategoryID:     CategoryIDColumn,
		Title:          TitleColumn,
		Description:    DescriptionColumn,
		ContentTitle:   ContentTitleColumn,
		Content:        ContentColumn,
		AdditionalData: AdditionalDataColumn,
		CreatorID:      CreatorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
