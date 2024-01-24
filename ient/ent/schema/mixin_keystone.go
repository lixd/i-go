package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// KeystoneMixin 实现了 ent.Mixin，包括 keystone 相关的几个通用字段
type KeystoneMixin struct {
	mixin.Schema
}

// Fields of the KeystoneMixin.
func (KeystoneMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain_id").
			MaxLen(40).
			NotEmpty().
			Comment("域ID").
			StorageKey("domain_id"),
		field.String("project_id").
			MaxLen(88).
			NotEmpty().
			Comment("项目ID").
			StorageKey("project_id"),
		field.String("account_id").
			MaxLen(88).
			NotEmpty().
			Comment("账号ID").
			StorageKey("account_id"),
	}
}
