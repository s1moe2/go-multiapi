package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	domain "multiapi"
)

// UserRepository is a service for user management.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository returns a new instance of UserRepository.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll fetches all users, returns an empty slice if no user exists
func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	users := []*domain.User{}
	err := r.db.SelectContext(ctx, &users, "SELECT id, name, email, picture FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID finds a user by ID, returns nil if not found
func (r *UserRepository) FindByID(ctx context.Context, ID string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.GetContext(ctx, user, "SELECT id, name, email, picture FROM users WHERE id = $1", ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// FindByEmail finds a user by email, returns nil if not found
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	stmt := "SELECT id, name, email, picture FROM users WHERE email = $1"
	err := r.db.GetContext(ctx, user, stmt, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// FindOrCreateUser finds a user by ID and creates it if not found
// returns a boolean indicating if the user was created
// TODO deal with passing txn around
func (r *UserRepository) FindOrCreateUser(ctx context.Context, userData *domain.User) (*domain.User, bool, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, false, err
	}
	defer tx.Rollback()

	user := &domain.User{}
	selectStmt := "SELECT id, name, email FROM users WHERE id = $1"
	err = tx.GetContext(ctx, user, selectStmt, userData.ID)
	if err == nil {
		return user, false, nil
	}
	if err != sql.ErrNoRows {
		return nil, false, parseError(err)
	}

	insertStmt := "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)"
	res, err := tx.ExecContext(ctx, insertStmt, userData.ID, userData.Name, userData.Email)
	if err != nil {
		return nil, false, parseError(err)
	}

	if rows, err := res.RowsAffected(); err != nil {
		if rows == 0 {
			return nil, false, errors.New("could not create user")
		}
		return nil, false, parseError(err)
	}

	err = tx.GetContext(ctx, user, selectStmt, userData.ID)
	if err != nil {
		return nil, false, parseError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, false, parseError(err)
	}

	return user, true, nil
}

// Create creates a new user, returning the full model
func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	stmt := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	row := r.db.QueryRowxContext(ctx, stmt, user.Name, user.Email)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, parseError(err)
	}
	return user, nil
}

// Update updates a user, returning the updated model or nil if no rows were affected
func (r *UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	stmt := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	res, err := r.db.ExecContext(ctx, stmt, user.Name, user.Email, user.ID)
	if err != nil {
		return nil, parseError(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, nil
	}
	return user, nil
}

// Delete deletes a user, only returns error if action fails
func (r *UserRepository) Delete(ctx context.Context, ID string) (bool, error) {
	stmt := "DELETE FROM users WHERE id = $1 RETURNING id"
	res, err := r.db.ExecContext(ctx, stmt, ID)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rows > 0, nil
}
