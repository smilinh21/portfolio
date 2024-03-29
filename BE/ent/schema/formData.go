package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type FormData struct {
	ent.Schema
}

func (FormData) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("message").Optional(),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional()}
}
func (FormData) Edges() []ent.Edge {
	return nil
}
func (FormData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Form Data"},
	}
}
