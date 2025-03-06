package schema

import (
	"time"
	"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("project_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.String("action").
			NotEmpty(),
		field.JSON("details", map[string]interface{}{}).
			Default(map[string]interface{}{}),
		field.String("ip_address").
			Optional(),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("histories").
			Field("user_id").
			Unique().
			Required(),
		edge.From("project", Project.Type).
			Ref("histories").
			Field("project_id").
			Unique().
			Required(),
		edge.To("history_projects", HistoryProject.Type),
	}
}
