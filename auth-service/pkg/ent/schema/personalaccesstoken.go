package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PersonalAccessToken holds the schema definition for the PersonalAccessToken entity.
type PersonalAccessToken struct {
	ent.Schema
}

// Fields of the PersonalAccessToken.
func (PersonalAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token_hash").Unique(),
		field.String("name"),
		field.JSON("abilities", []string{}),
		field.Time("last_used").Optional(),
		field.Time("expires_at"),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the PersonalAccessToken.
func (PersonalAccessToken) Edges() []ent.Edge {
	return nil
}
