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

var ArpanetUserPermissions = newArpanetUserPermissionsTable("arpanet", "arpanet_user_permissions", "")

type arpanetUserPermissionsTable struct {
	mysql.Table

	//Columns
	UserID       mysql.ColumnInteger
	PermissionID mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetUserPermissionsTable struct {
	arpanetUserPermissionsTable

	NEW arpanetUserPermissionsTable
}

// AS creates new ArpanetUserPermissionsTable with assigned alias
func (a ArpanetUserPermissionsTable) AS(alias string) *ArpanetUserPermissionsTable {
	return newArpanetUserPermissionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetUserPermissionsTable with assigned schema name
func (a ArpanetUserPermissionsTable) FromSchema(schemaName string) *ArpanetUserPermissionsTable {
	return newArpanetUserPermissionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetUserPermissionsTable with assigned table prefix
func (a ArpanetUserPermissionsTable) WithPrefix(prefix string) *ArpanetUserPermissionsTable {
	return newArpanetUserPermissionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetUserPermissionsTable with assigned table suffix
func (a ArpanetUserPermissionsTable) WithSuffix(suffix string) *ArpanetUserPermissionsTable {
	return newArpanetUserPermissionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetUserPermissionsTable(schemaName, tableName, alias string) *ArpanetUserPermissionsTable {
	return &ArpanetUserPermissionsTable{
		arpanetUserPermissionsTable: newArpanetUserPermissionsTableImpl(schemaName, tableName, alias),
		NEW:                         newArpanetUserPermissionsTableImpl("", "new", ""),
	}
}

func newArpanetUserPermissionsTableImpl(schemaName, tableName, alias string) arpanetUserPermissionsTable {
	var (
		UserIDColumn       = mysql.IntegerColumn("user_id")
		PermissionIDColumn = mysql.IntegerColumn("permission_id")
		allColumns         = mysql.ColumnList{UserIDColumn, PermissionIDColumn}
		mutableColumns     = mysql.ColumnList{}
	)

	return arpanetUserPermissionsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:       UserIDColumn,
		PermissionID: PermissionIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
