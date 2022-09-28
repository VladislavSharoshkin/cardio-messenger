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

var Messages = newMessagesTable("public", "messages", "")

type messagesTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestampz
	UpdatedAt postgres.ColumnTimestampz
	DeletedAt postgres.ColumnTimestampz
	SenderID  postgres.ColumnInteger
	ChatID    postgres.ColumnInteger
	Text      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type MessagesTable struct {
	messagesTable

	EXCLUDED messagesTable
}

// AS creates new MessagesTable with assigned alias
func (a MessagesTable) AS(alias string) *MessagesTable {
	return newMessagesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MessagesTable with assigned schema name
func (a MessagesTable) FromSchema(schemaName string) *MessagesTable {
	return newMessagesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MessagesTable with assigned table prefix
func (a MessagesTable) WithPrefix(prefix string) *MessagesTable {
	return newMessagesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MessagesTable with assigned table suffix
func (a MessagesTable) WithSuffix(suffix string) *MessagesTable {
	return newMessagesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMessagesTable(schemaName, tableName, alias string) *MessagesTable {
	return &MessagesTable{
		messagesTable: newMessagesTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newMessagesTableImpl("", "excluded", ""),
	}
}

func newMessagesTableImpl(schemaName, tableName, alias string) messagesTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn = postgres.TimestampzColumn("updated_at")
		DeletedAtColumn = postgres.TimestampzColumn("deleted_at")
		SenderIDColumn  = postgres.IntegerColumn("sender_id")
		ChatIDColumn    = postgres.IntegerColumn("chat_id")
		TextColumn      = postgres.StringColumn("text")
		allColumns      = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, SenderIDColumn, ChatIDColumn, TextColumn}
		mutableColumns  = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, SenderIDColumn, ChatIDColumn, TextColumn}
	)

	return messagesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,
		SenderID:  SenderIDColumn,
		ChatID:    ChatIDColumn,
		Text:      TextColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
