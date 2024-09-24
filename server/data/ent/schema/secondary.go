package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Secondary holds the schema definition for the Secondary entity.
type Secondary struct {
	ent.Schema
}

// Fields of the Secondary.
func (Secondary) Fields() []ent.Field {
	return []ent.Field{
		field.String("sid"),
		field.String("steam_id"),
		field.String("address"),
		field.Int("port"),
		field.Int("owner_id"),
		field.Int64("query_version"),
	}
}

// Edges of the Secondary.
func (Secondary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("servers", Server.Type).
			Ref("secondaries").Field("owner_id").
			Unique().Required(),
	}
}
