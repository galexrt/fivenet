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

var ArpanetUserLocations = newArpanetUserLocationsTable("arpanet", "arpanet_user_locations", "")

type arpanetUserLocationsTable struct {
	mysql.Table

	//Columns
	UserID    mysql.ColumnInteger
	Job       mysql.ColumnString
	X         mysql.ColumnFloat
	Y         mysql.ColumnFloat
	Hidden    mysql.ColumnBool
	UpdatedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetUserLocationsTable struct {
	arpanetUserLocationsTable

	NEW arpanetUserLocationsTable
}

// AS creates new ArpanetUserLocationsTable with assigned alias
func (a ArpanetUserLocationsTable) AS(alias string) *ArpanetUserLocationsTable {
	return newArpanetUserLocationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetUserLocationsTable with assigned schema name
func (a ArpanetUserLocationsTable) FromSchema(schemaName string) *ArpanetUserLocationsTable {
	return newArpanetUserLocationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetUserLocationsTable with assigned table prefix
func (a ArpanetUserLocationsTable) WithPrefix(prefix string) *ArpanetUserLocationsTable {
	return newArpanetUserLocationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetUserLocationsTable with assigned table suffix
func (a ArpanetUserLocationsTable) WithSuffix(suffix string) *ArpanetUserLocationsTable {
	return newArpanetUserLocationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetUserLocationsTable(schemaName, tableName, alias string) *ArpanetUserLocationsTable {
	return &ArpanetUserLocationsTable{
		arpanetUserLocationsTable: newArpanetUserLocationsTableImpl(schemaName, tableName, alias),
		NEW:                       newArpanetUserLocationsTableImpl("", "new", ""),
	}
}

func newArpanetUserLocationsTableImpl(schemaName, tableName, alias string) arpanetUserLocationsTable {
	var (
		UserIDColumn    = mysql.IntegerColumn("user_id")
		JobColumn       = mysql.StringColumn("job")
		XColumn         = mysql.FloatColumn("x")
		YColumn         = mysql.FloatColumn("y")
		HiddenColumn    = mysql.BoolColumn("hidden")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		allColumns      = mysql.ColumnList{UserIDColumn, JobColumn, XColumn, YColumn, HiddenColumn, UpdatedAtColumn}
		mutableColumns  = mysql.ColumnList{JobColumn, XColumn, YColumn, HiddenColumn, UpdatedAtColumn}
	)

	return arpanetUserLocationsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:    UserIDColumn,
		Job:       JobColumn,
		X:         XColumn,
		Y:         YColumn,
		Hidden:    HiddenColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
