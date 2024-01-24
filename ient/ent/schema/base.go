package schema

import (
	"context"
	"fmt"
	"i-go/ient/ent/schema/mtime"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// AuditMixin 实现了 ent.Mixin，
// 用于 schema 包内通用的审计日志功能。
type AuditMixin struct {
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		// 创建
		field.Time("created_at").
			GoType(mtime.Time{}).
			Immutable().
			Default(mtime.Now),
		field.String("created_by").
			Nillable().
			Optional(),
		// 更新
		field.Time("updated_at").
			GoType(mtime.Time{}).
			Optional().
			Nillable(),
		field.String("updated_by").
			Nillable().
			Optional(),
		//// 删除
		//field.Time("deleted_at").
		//	GoType(mtime.Time{}).
		//	Optional().
		//	Nillable(),
		//field.String("deleted_by").
		//	Optional().
		//	Nillable(),
	}
}

// Hooks AuditMixin 的钩子
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

// AuditHook 是审计日志钩子的示例
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger 声明了所有嵌入 AuditLog mixin 的
	// Schema 更变所共有的方法。 若变量 "existence "为真，
	// 则该字段存在于此更变中 (例如被其它的钩子设置过)。
	type AuditLogger interface {
		SetCreatedAt(mtime.Time)
		CreatedAt() (value mtime.Time, exists bool)
		SetCreatedBy(string)
		CreatedBy() (id string, exists bool)
		SetUpdatedAt(mtime.Time)
		UpdatedAt() (value mtime.Time, exists bool)
		SetUpdatedBy(string)
		UpdatedBy() (id string, exists bool)
		//SetDeletedAt(value mtime.Time)
		//DeletedAt() (r mtime.Time, exists bool)
		//SetDeletedBy(string)
		//DeletedBy() (id string, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		// usr, err := viewer.UserFromContext(ctx)
		// if err != nil {
		//	return nil, err
		// }
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(mtime.Now())
			// if _, exists := ml.CreatedBy(); !exists {
			//	ml.SetCreatedBy(usr.ID)
			// }
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(mtime.Now())
			// if _, exists := ml.UpdatedBy(); !exists {
			//	ml.SetUpdatedBy(usr.ID)
			// }
			//case op.Is(ent.OpDeleteOne | ent.OpDelete):
			//	ml.SetDeletedAt(mtime.Now())
			//if _, exists := ml.DeletedBy(); !exists {
			//	ml.SetDeletedBy(usr.ID)
			//}
		}
		return next.Mutate(ctx, m)
	})
}

const (
	len32   = 32
	len64   = 64
	len128  = 128
	len1024 = 1024
)
