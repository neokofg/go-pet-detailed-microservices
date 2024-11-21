package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News.
func (News) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("title").
			NotEmpty().
			MaxLen(255),
		field.Text("content").
			NotEmpty(),
		field.String("image_url").
			Optional().
			MaxLen(2048),
		field.UUID("user_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (News) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("created_at"),
	}
}

// Edges of the News.
func (News) Edges() []ent.Edge {
	return nil
}
