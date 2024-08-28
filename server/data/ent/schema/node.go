package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dstgo/lobby/pkg/ids"
	"github.com/dstgo/lobby/pkg/ts"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

func (Node) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("remote node info table"),
	}
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").DefaultFunc(ids.ULID).Unique(),
		field.String("name").Unique(),
		field.String("address").Unique(),
		field.String("note"),
		field.Int64("created_at").DefaultFunc(ts.UnixMicro),
		field.Int64("updated_at").DefaultFunc(ts.UnixMicro).UpdateDefault(ts.UnixMicro),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("containers", Container.Type),
	}
}
