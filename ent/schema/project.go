package schema

import (
	"time"
	"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("name").
			NotEmpty(),
		field.Text("description").
			Optional(),
		field.String("git_repository_url").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("status").
			Values("active", "archived", "deleted").
			Default("active"),
		field.UUID("owner_id", uuid.UUID{}),
		field.JSON("settings", map[string]interface{}{}).
			Default(map[string]interface{}{}),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("owned_projects").
			Field("owner_id").
			Unique().
			Required(),
		edge.To("members", ProjectMember.Type),
		edge.To("code_embeddings", CodeEmbedding.Type),
		edge.To("histories", History.Type),
		edge.To("history_projects", HistoryProject.Type),
	}
}
