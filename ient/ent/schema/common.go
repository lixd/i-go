package schema

import (
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

var (
	idBuilder = field.String("id").
			MaxLen(64).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(func() string {
			return uuid.NewString()
		})

	nameBuilder = field.String("name").
			MaxLen(len64).
			StructTag(`validate:"required,min=1,max=64"`).
			Validate(func(s string) error {
			return nil
		})

	descriptionBuilder = field.String("description").
				MaxLen(len1024)

	domainBuilder = field.String("domain").
			MaxLen(len128).
			StructTag(`validate:"required,max=128"`).
			NotEmpty().
			Immutable()

	userBuilder = field.String("user").
			MaxLen(len128).
			StructTag(`validate:"required,max=128"`).
			NotEmpty().
			Immutable()

	workflowBuilder = field.String("workflow").
			MaxLen(len64).
			NotEmpty().
			Immutable().
			StructTag(`validate:"required,max=64"`)
)
