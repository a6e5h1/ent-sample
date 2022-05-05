package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),
		field.UUID("uuid", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.Uint64("user_id"),
		field.String("title").
			MaxLen(255).
			NotEmpty(),
		field.String("body").
			MaxLen(10000).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Nillable().
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poster", User.Type).
			Ref("posts").
			Field("user_id").
			Required().
			Unique(),
		edge.To("comments", Comment.Type),
	}
}
