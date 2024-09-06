// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CronJobsColumns holds the columns for the "cron_jobs" table.
	CronJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "cron", Type: field.TypeString},
		{Name: "entry_id", Type: field.TypeInt},
		{Name: "prev", Type: field.TypeInt64},
		{Name: "next", Type: field.TypeInt64},
	}
	// CronJobsTable holds the schema information for the "cron_jobs" table.
	CronJobsTable = &schema.Table{
		Name:       "cron_jobs",
		Columns:    CronJobsColumns,
		PrimaryKey: []*schema.Column{CronJobsColumns[0]},
	}
	// SecondariesColumns holds the columns for the "secondaries" table.
	SecondariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sid", Type: field.TypeString},
		{Name: "steam_id", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "port", Type: field.TypeInt},
		{Name: "owner_id", Type: field.TypeInt},
	}
	// SecondariesTable holds the schema information for the "secondaries" table.
	SecondariesTable = &schema.Table{
		Name:       "secondaries",
		Columns:    SecondariesColumns,
		PrimaryKey: []*schema.Column{SecondariesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "secondaries_servers_secondaries",
				Columns:    []*schema.Column{SecondariesColumns[5]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ServersColumns holds the columns for the "servers" table.
	ServersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "guid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "row_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(40)"}},
		{Name: "steam_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "steam_clan_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(25)"}},
		{Name: "owner_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(25)"}},
		{Name: "steam_room", Type: field.TypeString},
		{Name: "session", Type: field.TypeString},
		{Name: "address", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(25)"}},
		{Name: "port", Type: field.TypeInt},
		{Name: "host", Type: field.TypeString},
		{Name: "platform", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "clan_only", Type: field.TypeBool},
		{Name: "lan_only", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "game_mode", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(40)"}},
		{Name: "intent", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(40)"}},
		{Name: "season", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(40)"}},
		{Name: "version", Type: field.TypeInt},
		{Name: "max_online", Type: field.TypeInt},
		{Name: "online", Type: field.TypeInt},
		{Name: "level", Type: field.TypeInt},
		{Name: "mod", Type: field.TypeBool},
		{Name: "pvp", Type: field.TypeBool},
		{Name: "password", Type: field.TypeBool},
		{Name: "dedicated", Type: field.TypeBool},
		{Name: "client_hosted", Type: field.TypeBool},
		{Name: "allow_new_players", Type: field.TypeBool},
		{Name: "server_paused", Type: field.TypeBool},
		{Name: "friend_only", Type: field.TypeBool},
		{Name: "query_version", Type: field.TypeInt64},
		{Name: "country", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "continent", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "country_code", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "city", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
		{Name: "region", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(50)"}},
	}
	// ServersTable holds the schema information for the "servers" table.
	ServersTable = &schema.Table{
		Name:       "servers",
		Columns:    ServersColumns,
		PrimaryKey: []*schema.Column{ServersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "server_query_version",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[30]},
			},
			{
				Name:    "server_version",
				Unique:  false,
				Columns: []*schema.Column{ServersColumns[18]},
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_servers_tags",
				Columns:    []*schema.Column{TagsColumns[2]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "user info table",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CronJobsTable,
		SecondariesTable,
		ServersTable,
		TagsTable,
		UsersTable,
	}
)

func init() {
	SecondariesTable.ForeignKeys[0].RefTable = ServersTable
	TagsTable.ForeignKeys[0].RefTable = ServersTable
	UsersTable.Annotation = &entsql.Annotation{}
}
