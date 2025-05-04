package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Email     *string    `bun:"email,unique"`
	Name      *string    `bun:"name"`
	Password  *string    `bun:"password"`
	CreatedAt *time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at,notnull"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
}
