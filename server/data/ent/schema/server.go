package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/246859/schematype"
	"github.com/dstgo/lobby/server/data/ent/server"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(server.FieldQueryVersion),
		index.Fields(server.FieldVersion),
	}
}

func (Server) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		// network options
		field.String("guid").SchemaType(schematype.MySQL().VarChar(50)),
		field.String("row_id").SchemaType(schematype.MySQL().VarChar(40)),
		// only at steam platform
		field.String("steam_id").SchemaType(schematype.MySQL().VarChar(50)),
		// only for clan server
		field.String("steam_clan_id").SchemaType(schematype.MySQL().VarChar(25)),
		// only for server without password
		field.String("owner_id").SchemaType(schematype.MySQL().VarChar(25)),
		field.String("steam_room"),
		field.String("session"),
		field.String("address").SchemaType(schematype.MySQL().VarChar(25)),
		field.Int("port"),
		field.String("host"),
		field.String("platform").SchemaType(schematype.MySQL().VarChar(30)),
		// lan or clan
		field.Bool("clan_only"),
		field.Bool("lan_only"),
		// game information
		field.String("name"),
		field.String("game_mode").SchemaType(schematype.MySQL().VarChar(40)),
		field.String("intent").SchemaType(schematype.MySQL().VarChar(40)),
		field.String("season").SchemaType(schematype.MySQL().VarChar(40)),
		field.Int("version"),
		field.Int("max_online"),
		field.Int("online"),
		field.Int("level"),
		// room options
		field.Bool("mod"),
		field.Bool("pvp"),
		field.Bool("password"),
		field.Bool("dedicated"),
		field.Bool("client_hosted"),
		field.Bool("allow_new_players"),
		field.Bool("server_paused"),
		field.Bool("friend_only"),
		// query version
		field.Int64("query_version"),
		// meta info
		field.String("country").SchemaType(schematype.MySQL().VarChar(50)),
		field.String("continent").SchemaType(schematype.MySQL().VarChar(50)),
		field.String("country_code").SchemaType(schematype.MySQL().VarChar(50)),
		field.String("city").SchemaType(schematype.MySQL().VarChar(50)),
		field.String("region").SchemaType(schematype.MySQL().VarChar(50)),
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("secondaries", Secondary.Type),
	}
}
