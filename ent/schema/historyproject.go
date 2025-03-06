package schema

import (
	"time"
	"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// HistoryProject holds the schema definition for the HistoryProject entity.
type HistoryProject struct {
	ent.Schema
}

// Fields of the HistoryProject.
func (HistoryProject) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("project_id", uuid.UUID{}),
		field.UUID("history_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.JSON("changes", map[string]interface{}{}).
			Default(map[string]interface{}{}),
		field.Enum("status").
			Values("active", "reverted").
			Default("active"),
	}
}

// Edges of the HistoryProject.
func (HistoryProject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("history_projects").
			Field("project_id").
			Unique().
			Required(),
		edge.From("history", History.Type).
			Ref("history_projects").
			Field("history_id").
			Unique().
			Required(),
	}
}

// Indexes of the HistoryProject.
func (HistoryProject) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("project_id", "history_id").
			Unique(),
	}
}
