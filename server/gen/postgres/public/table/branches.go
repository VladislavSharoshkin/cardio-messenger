//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Branches = newBranchesTable("public", "branches", "")

type branchesTable struct {
	postgres.Table

	//Columns
	ID           postgres.ColumnInteger
	Name         postgres.ColumnString
	SyncID       postgres.ColumnInteger
	SyncRemoteID postgres.ColumnString
	CreatedAt    postgres.ColumnTimestampz
	UpdatedAt    postgres.ColumnTimestampz
	DeletedAt    postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BranchesTable struct {
	branchesTable

	EXCLUDED branchesTable
}

// AS creates new BranchesTable with assigned alias
func (a BranchesTable) AS(alias string) *BranchesTable {
	return newBranchesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BranchesTable with assigned schema name
func (a BranchesTable) FromSchema(schemaName string) *BranchesTable {
	return newBranchesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BranchesTable with assigned table prefix
func (a BranchesTable) WithPrefix(prefix string) *BranchesTable {
	return newBranchesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BranchesTable with assigned table suffix
func (a BranchesTable) WithSuffix(suffix string) *BranchesTable {
	return newBranchesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBranchesTable(schemaName, tableName, alias string) *BranchesTable {
	return &BranchesTable{
		branchesTable: newBranchesTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newBranchesTableImpl("", "excluded", ""),
	}
}

func newBranchesTableImpl(schemaName, tableName, alias string) branchesTable {
	var (
		IDColumn           = postgres.IntegerColumn("id")
		NameColumn         = postgres.StringColumn("name")
		SyncIDColumn       = postgres.IntegerColumn("sync_id")
		SyncRemoteIDColumn = postgres.StringColumn("sync_remote_id")
		CreatedAtColumn    = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn    = postgres.TimestampzColumn("updated_at")
		DeletedAtColumn    = postgres.TimestampzColumn("deleted_at")
		allColumns         = postgres.ColumnList{IDColumn, NameColumn, SyncIDColumn, SyncRemoteIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns     = postgres.ColumnList{NameColumn, SyncIDColumn, SyncRemoteIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return branchesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		Name:         NameColumn,
		SyncID:       SyncIDColumn,
		SyncRemoteID: SyncRemoteIDColumn,
		CreatedAt:    CreatedAtColumn,
		UpdatedAt:    UpdatedAtColumn,
		DeletedAt:    DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
