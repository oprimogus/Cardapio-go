// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateProfile(ctx context.Context, arg CreateProfileParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetProfile(ctx context.Context, id int32) (Profile, error)
	GetUser(ctx context.Context, arg GetUserParams) ([]GetUserRow, error)
	GetUserById(ctx context.Context, id pgtype.UUID) (GetUserByIdRow, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) error
	UpdateProfileCpf(ctx context.Context, arg UpdateProfileCpfParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error
}

var _ Querier = (*Queries)(nil)