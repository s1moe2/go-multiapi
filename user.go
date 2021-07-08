package domain

import "context"

func (p *User) validate() []string {
	var errs []string

	if len(p.Name) < 3 {
		errs = append(errs, "name: invalid length")
	}

	if len(p.Email) < 5 {
		errs = append(errs, "email: invalid length")
	}

	return errs
}

// User represents a user, which in the context of
// this application is always an administrator.
type User struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
}

// UserRepository represents the interface of a user management service.
type UserRepository interface {
	GetAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, ID string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindOrCreateUser(ctx context.Context, userData *User) (*User, bool, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, ID string) (bool, error)
}