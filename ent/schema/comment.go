package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),
		field.UUID("uuid", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.Uint64("user_id"),
		field.Uint64("post_id"),
		field.String("content").
			MaxLen(500).
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

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("commenter", User.Type).
			Ref("comments").
			Field("user_id").
			Required().
			Unique(),
		edge.From("post", Post.Type).
			Ref("comments").
			Field("post_id").
			Required().
			Unique(),
	}
}
