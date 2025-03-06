package schema

import (
    "time"
    "github.com/google/uuid"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            Immutable(),
        field.String("email").
            Unique().
            NotEmpty(),
        field.String("password").
            NotEmpty().
            Sensitive(),
        field.String("name").
            NotEmpty(),
        field.Time("created_at").
            Default(time.Now).
            Immutable(),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
        field.Time("last_login").
            Optional().
            Nillable(),
        field.Enum("status").
            Values("active", "inactive", "suspended").
            Default("active"),
        field.Enum("role").
            Values("user", "admin", "manager").
            Default("user"),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("owned_projects", Project.Type),
        edge.To("project_memberships", ProjectMember.Type),
        edge.To("histories", History.Type),
    }
}
