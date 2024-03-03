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

var FivenetJobsUserProps = newFivenetJobsUserPropsTable("", "fivenet_jobs_user_props", "")

type fivenetJobsUserPropsTable struct {
	mysql.Table

	// Columns
	UserID       mysql.ColumnInteger
	AbsenceBegin mysql.ColumnDate
	AbsenceEnd   mysql.ColumnDate

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetJobsUserPropsTable struct {
	fivenetJobsUserPropsTable

	NEW fivenetJobsUserPropsTable
}

// AS creates new FivenetJobsUserPropsTable with assigned alias
func (a FivenetJobsUserPropsTable) AS(alias string) *FivenetJobsUserPropsTable {
	return newFivenetJobsUserPropsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetJobsUserPropsTable with assigned schema name
func (a FivenetJobsUserPropsTable) FromSchema(schemaName string) *FivenetJobsUserPropsTable {
	return newFivenetJobsUserPropsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetJobsUserPropsTable with assigned table prefix
func (a FivenetJobsUserPropsTable) WithPrefix(prefix string) *FivenetJobsUserPropsTable {
	return newFivenetJobsUserPropsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetJobsUserPropsTable with assigned table suffix
func (a FivenetJobsUserPropsTable) WithSuffix(suffix string) *FivenetJobsUserPropsTable {
	return newFivenetJobsUserPropsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetJobsUserPropsTable(schemaName, tableName, alias string) *FivenetJobsUserPropsTable {
	return &FivenetJobsUserPropsTable{
		fivenetJobsUserPropsTable: newFivenetJobsUserPropsTableImpl(schemaName, tableName, alias),
		NEW:                       newFivenetJobsUserPropsTableImpl("", "new", ""),
	}
}

func newFivenetJobsUserPropsTableImpl(schemaName, tableName, alias string) fivenetJobsUserPropsTable {
	var (
		UserIDColumn       = mysql.IntegerColumn("user_id")
		AbsenceBeginColumn = mysql.DateColumn("absence_begin")
		AbsenceEndColumn   = mysql.DateColumn("absence_end")
		allColumns         = mysql.ColumnList{UserIDColumn, AbsenceBeginColumn, AbsenceEndColumn}
		mutableColumns     = mysql.ColumnList{UserIDColumn, AbsenceBeginColumn, AbsenceEndColumn}
	)

	return fivenetJobsUserPropsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:       UserIDColumn,
		AbsenceBegin: AbsenceBeginColumn,
		AbsenceEnd:   AbsenceEndColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
