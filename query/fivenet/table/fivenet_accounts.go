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

var FivenetAccounts = newFivenetAccountsTable("", "fivenet_accounts", "")

type fivenetAccountsTable struct {
	mysql.Table

	// Columns
	ID               mysql.ColumnInteger
	CreatedAt        mysql.ColumnTimestamp
	UpdatedAt        mysql.ColumnTimestamp
	Enabled          mysql.ColumnBool
	Username         mysql.ColumnString
	Password         mysql.ColumnString
	License          mysql.ColumnString
	RegToken         mysql.ColumnString
	OverrideJob      mysql.ColumnString
	OverrideJobGrade mysql.ColumnInteger
	Superuser        mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetAccountsTable struct {
	fivenetAccountsTable

	NEW fivenetAccountsTable
}

// AS creates new FivenetAccountsTable with assigned alias
func (a FivenetAccountsTable) AS(alias string) *FivenetAccountsTable {
	return newFivenetAccountsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetAccountsTable with assigned schema name
func (a FivenetAccountsTable) FromSchema(schemaName string) *FivenetAccountsTable {
	return newFivenetAccountsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetAccountsTable with assigned table prefix
func (a FivenetAccountsTable) WithPrefix(prefix string) *FivenetAccountsTable {
	return newFivenetAccountsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetAccountsTable with assigned table suffix
func (a FivenetAccountsTable) WithSuffix(suffix string) *FivenetAccountsTable {
	return newFivenetAccountsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetAccountsTable(schemaName, tableName, alias string) *FivenetAccountsTable {
	return &FivenetAccountsTable{
		fivenetAccountsTable: newFivenetAccountsTableImpl(schemaName, tableName, alias),
		NEW:                  newFivenetAccountsTableImpl("", "new", ""),
	}
}

func newFivenetAccountsTableImpl(schemaName, tableName, alias string) fivenetAccountsTable {
	var (
		IDColumn               = mysql.IntegerColumn("id")
		CreatedAtColumn        = mysql.TimestampColumn("created_at")
		UpdatedAtColumn        = mysql.TimestampColumn("updated_at")
		EnabledColumn          = mysql.BoolColumn("enabled")
		UsernameColumn         = mysql.StringColumn("username")
		PasswordColumn         = mysql.StringColumn("password")
		LicenseColumn          = mysql.StringColumn("license")
		RegTokenColumn         = mysql.StringColumn("reg_token")
		OverrideJobColumn      = mysql.StringColumn("override_job")
		OverrideJobGradeColumn = mysql.IntegerColumn("override_job_grade")
		SuperuserColumn        = mysql.BoolColumn("superuser")
		allColumns             = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, EnabledColumn, UsernameColumn, PasswordColumn, LicenseColumn, RegTokenColumn, OverrideJobColumn, OverrideJobGradeColumn, SuperuserColumn}
		mutableColumns         = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, EnabledColumn, UsernameColumn, PasswordColumn, LicenseColumn, RegTokenColumn, OverrideJobColumn, OverrideJobGradeColumn, SuperuserColumn}
	)

	return fivenetAccountsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		CreatedAt:        CreatedAtColumn,
		UpdatedAt:        UpdatedAtColumn,
		Enabled:          EnabledColumn,
		Username:         UsernameColumn,
		Password:         PasswordColumn,
		License:          LicenseColumn,
		RegToken:         RegTokenColumn,
		OverrideJob:      OverrideJobColumn,
		OverrideJobGrade: OverrideJobGradeColumn,
		Superuser:        SuperuserColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
