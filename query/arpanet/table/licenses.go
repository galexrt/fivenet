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

var Licenses = newLicensesTable("", "licenses", "")

type licensesTable struct {
	mysql.Table

	//Columns
	Type  mysql.ColumnString
	Label mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type LicensesTable struct {
	licensesTable

	NEW licensesTable
}

// AS creates new LicensesTable with assigned alias
func (a LicensesTable) AS(alias string) *LicensesTable {
	return newLicensesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LicensesTable with assigned schema name
func (a LicensesTable) FromSchema(schemaName string) *LicensesTable {
	return newLicensesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LicensesTable with assigned table prefix
func (a LicensesTable) WithPrefix(prefix string) *LicensesTable {
	return newLicensesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LicensesTable with assigned table suffix
func (a LicensesTable) WithSuffix(suffix string) *LicensesTable {
	return newLicensesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLicensesTable(schemaName, tableName, alias string) *LicensesTable {
	return &LicensesTable{
		licensesTable: newLicensesTableImpl(schemaName, tableName, alias),
		NEW:           newLicensesTableImpl("", "new", ""),
	}
}

func newLicensesTableImpl(schemaName, tableName, alias string) licensesTable {
	var (
		TypeColumn     = mysql.StringColumn("type")
		LabelColumn    = mysql.StringColumn("label")
		allColumns     = mysql.ColumnList{TypeColumn, LabelColumn}
		mutableColumns = mysql.ColumnList{LabelColumn}
	)

	return licensesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Type:  TypeColumn,
		Label: LabelColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
