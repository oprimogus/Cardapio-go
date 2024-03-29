package user

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/oprimogus/cardapiogo/internal/domain/types"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

const (
	INVALID_PASSWORD = "Invalid Password."
)

// Service struct
type Service struct {
	repository Repository
}

// NewService Service constructor
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

// CreateUser create a user in database
func (u *Service) CreateUser(ctx context.Context, newUser CreateUserParams) error {
	hashPassword, err := u.HashPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("fail in generate hash of password: %s", err)
	}
	newUser.Password = hashPassword
	newUser.AccountProvider = types.AccountProviderLocal
	return u.repository.CreateUser(ctx, newUser)
}

// CreateUserWithOAuth create a user in database
func (u *Service) CreateUserWithOAuth(ctx context.Context, newUser CreateUserWithOAuthParams) error {
	return u.repository.CreateUserWithOAuth(ctx, newUser)
}

// GetUser return a user from database by ID
func (u *Service) GetUser(ctx context.Context, id string) (User, error) {
	return u.repository.GetUserByID(ctx, id)
}

// GetUser return a user from database by email
func (u *Service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return u.repository.GetUserByEmail(ctx, email)
}

// GetUsersList return a user from database
func (u *Service) GetUsersList(ctx context.Context, items int, page int) ([]*User, error) {
	return u.repository.GetUsersList(ctx, items, page)
}

// UpdateUserPassword change the password of user
func (u *Service) UpdateUserPassword(ctx context.Context, params UpdateUserPasswordParams) error {
	user, err := u.GetUser(ctx, params.ID)
	if err != nil {
		return err
	}

	if !u.IsValidPassword(params.Password, user.Password) {
		return errors.New(http.StatusBadRequest, INVALID_PASSWORD)
	}

	params.NewPassword, err = u.HashPassword(params.NewPassword)
	if err != nil {
		return err
	}
	return u.repository.UpdateUserPassword(ctx, params)
}

// UpdateUser change the user data
func (u *Service) UpdateUser(ctx context.Context, params UpdateUserParams) error {
	user, err := u.GetUser(ctx, params.ID)
	if err != nil {
		return err
	}

	if !u.IsValidPassword(params.Password, user.Password) {
		return errors.New(http.StatusBadRequest, INVALID_PASSWORD)
	}

	return u.repository.UpdateUser(ctx, params)
}

// UpdateUserProfile set the profile of a user
func (u *Service) UpdateUserProfile(ctx context.Context, params UpdateUserProfileParams) error {
	return u.repository.UpdateUserProfile(ctx, params)
}

// HashPassword generate a hash of password for save in database
func (u *Service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// IsValidPassword verify if password hash is valid
func (u *Service) IsValidPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
