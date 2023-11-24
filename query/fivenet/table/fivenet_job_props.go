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

var FivenetJobProps = newFivenetJobPropsTable("", "fivenet_job_props", "")

type fivenetJobPropsTable struct {
	mysql.Table

	// Columns
	Job                 mysql.ColumnString
	UpdatedAt           mysql.ColumnTimestamp
	Theme               mysql.ColumnString
	LivemapMarkerColor  mysql.ColumnString
	QuickButtons        mysql.ColumnString
	DiscordGuildID      mysql.ColumnInteger
	DiscordLastSync     mysql.ColumnTimestamp
	DiscordSyncSettings mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetJobPropsTable struct {
	fivenetJobPropsTable

	NEW fivenetJobPropsTable
}

// AS creates new FivenetJobPropsTable with assigned alias
func (a FivenetJobPropsTable) AS(alias string) *FivenetJobPropsTable {
	return newFivenetJobPropsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetJobPropsTable with assigned schema name
func (a FivenetJobPropsTable) FromSchema(schemaName string) *FivenetJobPropsTable {
	return newFivenetJobPropsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetJobPropsTable with assigned table prefix
func (a FivenetJobPropsTable) WithPrefix(prefix string) *FivenetJobPropsTable {
	return newFivenetJobPropsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetJobPropsTable with assigned table suffix
func (a FivenetJobPropsTable) WithSuffix(suffix string) *FivenetJobPropsTable {
	return newFivenetJobPropsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetJobPropsTable(schemaName, tableName, alias string) *FivenetJobPropsTable {
	return &FivenetJobPropsTable{
		fivenetJobPropsTable: newFivenetJobPropsTableImpl(schemaName, tableName, alias),
		NEW:                  newFivenetJobPropsTableImpl("", "new", ""),
	}
}

func newFivenetJobPropsTableImpl(schemaName, tableName, alias string) fivenetJobPropsTable {
	var (
		JobColumn                 = mysql.StringColumn("job")
		UpdatedAtColumn           = mysql.TimestampColumn("updated_at")
		ThemeColumn               = mysql.StringColumn("theme")
		LivemapMarkerColorColumn  = mysql.StringColumn("livemap_marker_color")
		QuickButtonsColumn        = mysql.StringColumn("quick_buttons")
		DiscordGuildIDColumn      = mysql.IntegerColumn("discord_guild_id")
		DiscordLastSyncColumn     = mysql.TimestampColumn("discord_last_sync")
		DiscordSyncSettingsColumn = mysql.StringColumn("discord_sync_settings")
		allColumns                = mysql.ColumnList{JobColumn, UpdatedAtColumn, ThemeColumn, LivemapMarkerColorColumn, QuickButtonsColumn, DiscordGuildIDColumn, DiscordLastSyncColumn, DiscordSyncSettingsColumn}
		mutableColumns            = mysql.ColumnList{JobColumn, UpdatedAtColumn, ThemeColumn, LivemapMarkerColorColumn, QuickButtonsColumn, DiscordGuildIDColumn, DiscordLastSyncColumn, DiscordSyncSettingsColumn}
	)

	return fivenetJobPropsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Job:                 JobColumn,
		UpdatedAt:           UpdatedAtColumn,
		Theme:               ThemeColumn,
		LivemapMarkerColor:  LivemapMarkerColorColumn,
		QuickButtons:        QuickButtonsColumn,
		DiscordGuildID:      DiscordGuildIDColumn,
		DiscordLastSync:     DiscordLastSyncColumn,
		DiscordSyncSettings: DiscordSyncSettingsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}