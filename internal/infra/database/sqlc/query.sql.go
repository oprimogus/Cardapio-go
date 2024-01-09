// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProfile = `-- name: CreateProfile :exec
INSERT INTO profile (name, last_name, cpf, phone, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
`

type CreateProfileParams struct {
	Name     string `db:"name" json:"name"`
	LastName string `db:"last_name" json:"last_name"`
	Cpf      string `db:"cpf" json:"cpf"`
	Phone    string `db:"phone" json:"phone"`
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) error {
	_, err := q.db.Exec(ctx, createProfile,
		arg.Name,
		arg.LastName,
		arg.Cpf,
		arg.Phone,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (email, password, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
`

type CreateUserParams struct {
	Email           string          `db:"email" json:"email"`
	Password        pgtype.Text     `db:"password" json:"password"`
	Role            UserRole        `db:"role" json:"role"`
	AccountProvider AccountProvider `db:"account_provider" json:"account_provider"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.AccountProvider,
	)
	return err
}

const getProfile = `-- name: GetProfile :one
SELECT id, name, last_name, cpf, phone, created_at, updated_at FROM profile
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProfile(ctx context.Context, id int32) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfile, id)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LastName,
		&i.Cpf,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :many
SELECT id, profile_id, email, role, created_at, updated_at FROM users
ORDER BY created_at desc
LIMIT $1 OFFSET $2
`

type GetUserParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

type GetUserRow struct {
	ID        pgtype.UUID        `db:"id" json:"id"`
	ProfileID pgtype.Int4        `db:"profile_id" json:"profile_id"`
	Email     string             `db:"email" json:"email"`
	Role      UserRole           `db:"role" json:"role"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) ([]GetUserRow, error) {
	rows, err := q.db.Query(ctx, getUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserRow
	for rows.Next() {
		var i GetUserRow
		if err := rows.Scan(
			&i.ID,
			&i.ProfileID,
			&i.Email,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, profile_id, email, password, role, created_at, updated_at FROM users
WHERE id = $1
LIMIT 1
`

type GetUserByIdRow struct {
	ID        pgtype.UUID        `db:"id" json:"id"`
	ProfileID pgtype.Int4        `db:"profile_id" json:"profile_id"`
	Email     string             `db:"email" json:"email"`
	Password  pgtype.Text        `db:"password" json:"password"`
	Role      UserRole           `db:"role" json:"role"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (GetUserByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.ProfileID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProfile = `-- name: UpdateProfile :exec
UPDATE profile
SET 
    name = $2,
    last_name = $3,
    phone = $4,
    updated_at = NOW()
WHERE id = $1
`

type UpdateProfileParams struct {
	ID       int32  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	LastName string `db:"last_name" json:"last_name"`
	Phone    string `db:"phone" json:"phone"`
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) error {
	_, err := q.db.Exec(ctx, updateProfile,
		arg.ID,
		arg.Name,
		arg.LastName,
		arg.Phone,
	)
	return err
}

const updateProfileCpf = `-- name: UpdateProfileCpf :exec
UPDATE profile
SET 
    cpf = $2,
    updated_at = NOW()
WHERE id = $1
`

type UpdateProfileCpfParams struct {
	ID  int32  `db:"id" json:"id"`
	Cpf string `db:"cpf" json:"cpf"`
}

func (q *Queries) UpdateProfileCpf(ctx context.Context, arg UpdateProfileCpfParams) error {
	_, err := q.db.Exec(ctx, updateProfileCpf, arg.ID, arg.Cpf)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET 
    email = $2,
    role = $3,
    updated_at = NOW()
WHERE id = $1
`

type UpdateUserParams struct {
	ID    pgtype.UUID `db:"id" json:"id"`
	Email string      `db:"email" json:"email"`
	Role  UserRole    `db:"role" json:"role"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.ID, arg.Email, arg.Role)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET 
    password = $2,
    updated_at = NOW()
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID       pgtype.UUID `db:"id" json:"id"`
	Password pgtype.Text `db:"password" json:"password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.ID, arg.Password)
	return err
}

const updateUserProfile = `-- name: UpdateUserProfile :exec
UPDATE users
SET 
    profile_id = $2,
    updated_at = NOW()
WHERE id = $1
`

type UpdateUserProfileParams struct {
	ID        pgtype.UUID `db:"id" json:"id"`
	ProfileID pgtype.Int4 `db:"profile_id" json:"profile_id"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error {
	_, err := q.db.Exec(ctx, updateUserProfile, arg.ID, arg.ProfileID)
	return err
}
