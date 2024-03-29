package user

import (
	"time"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
)

type User struct {
	ID              string                `db:"id" json:"id"`
	ProfileID       int                   `db:"profile_id" json:"profile_id"`
	Email           string                `db:"email" json:"email"`
	Password        string                `db:"password" json:"-"`
	Role            types.Role            `db:"role" json:"role"`
	AccountProvider types.AccountProvider `db:"account_provider" json:"account_provider"`
	CreatedAt       time.Time             `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time             `db:"updated_at" json:"updated_at"`
}
