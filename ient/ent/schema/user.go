package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"i-go/ient/ent/schema/property"
)

// User holds the schema definition for the Tag entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		idBuilder,
		field.String("name").
			Comment("名称").
			StorageKey("name"),
		field.Int("age").
			Comment("年龄").
			StorageKey("age"),
		field.Enum("type").
			StructTag(`validate:"required"`).
			Comment("类型").
			GoType(property.UserType("")),
		field.String("address").
			MaxLen(64).
			NotEmpty().
			Comment("地址").
			StorageKey("address"),
		field.String("phone").
			Comment("手机").
			StorageKey("phone"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//KeystoneMixin{},
		//AuditMixin{},
	}
}

// TODO(user): add indexes here.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
