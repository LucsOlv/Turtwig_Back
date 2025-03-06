package schema

import (
	"time"
	"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// CodeEmbedding holds the schema definition for the CodeEmbedding entity.
type CodeEmbedding struct {
	ent.Schema
}

// Fields of the CodeEmbedding.
func (CodeEmbedding) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("file_name").
			NotEmpty(),
		field.Text("file_path").
			NotEmpty(),
		field.JSON("embedding", []float32{}).
			Optional(),
		field.UUID("project_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("language").
			Optional(),
		field.String("version").
			Optional(),
	}
}

// Edges of the CodeEmbedding.
func (CodeEmbedding) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("code_embeddings").
			Field("project_id").
			Unique().
			Required(),
	}
}
