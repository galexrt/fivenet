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

var ArpanetAccounts = newArpanetAccountsTable("", "arpanet_accounts", "")

type arpanetAccountsTable struct {
	mysql.Table

	//Columns
	ID        mysql.ColumnInteger
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp
	Enabled   mysql.ColumnBool
	Username  mysql.ColumnString
	Password  mysql.ColumnString
	License   mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetAccountsTable struct {
	arpanetAccountsTable

	NEW arpanetAccountsTable
}

// AS creates new ArpanetAccountsTable with assigned alias
func (a ArpanetAccountsTable) AS(alias string) *ArpanetAccountsTable {
	return newArpanetAccountsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetAccountsTable with assigned schema name
func (a ArpanetAccountsTable) FromSchema(schemaName string) *ArpanetAccountsTable {
	return newArpanetAccountsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetAccountsTable with assigned table prefix
func (a ArpanetAccountsTable) WithPrefix(prefix string) *ArpanetAccountsTable {
	return newArpanetAccountsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetAccountsTable with assigned table suffix
func (a ArpanetAccountsTable) WithSuffix(suffix string) *ArpanetAccountsTable {
	return newArpanetAccountsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetAccountsTable(schemaName, tableName, alias string) *ArpanetAccountsTable {
	return &ArpanetAccountsTable{
		arpanetAccountsTable: newArpanetAccountsTableImpl(schemaName, tableName, alias),
		NEW:                  newArpanetAccountsTableImpl("", "new", ""),
	}
}

func newArpanetAccountsTableImpl(schemaName, tableName, alias string) arpanetAccountsTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		EnabledColumn   = mysql.BoolColumn("enabled")
		UsernameColumn  = mysql.StringColumn("username")
		PasswordColumn  = mysql.StringColumn("password")
		LicenseColumn   = mysql.StringColumn("license")
		allColumns      = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, EnabledColumn, UsernameColumn, PasswordColumn, LicenseColumn}
		mutableColumns  = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, EnabledColumn, UsernameColumn, PasswordColumn, LicenseColumn}
	)

	return arpanetAccountsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		Enabled:   EnabledColumn,
		Username:  UsernameColumn,
		Password:  PasswordColumn,
		License:   LicenseColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
