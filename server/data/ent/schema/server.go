package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

func (Server) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		// network options
		field.String("guid"),
		field.String("row_id"),
		// only at steam platform
		field.String("steam_id"),
		// only for clan server
		field.String("steam_clan_id"),
		// only for server without password
		field.String("owner_id"),
		field.String("steam_room"),
		field.String("session"),
		field.String("address"),
		field.Int("port"),
		field.String("host"),
		field.String("platform"),
		// lan or clan
		field.Bool("clan_only"),
		field.Bool("lan_only"),
		// game information
		field.String("name"),
		field.String("game_mode"),
		field.String("intent"),
		field.String("season"),
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
		field.String("country"),
		field.String("continent"),
		field.String("country_code"),
		field.String("city"),
		field.String("region"),
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("secondaries", Secondary.Type),
	}
}
