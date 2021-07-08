package mock

import (
	"context"
	"multiapi"
)

type UserRepository struct {
	GetAllFn           func(ctx context.Context) ([]*domain.User, error)
	FindByIDFn         func(ctx context.Context, ID string) (*domain.User, error)
	FindByEmailFn      func(ctx context.Context, email string) (*domain.User, error)
	FindOrCreateUserFn func(ctx context.Context, userData *domain.User) (*domain.User, bool, error)
	CreateFn           func(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateFn           func(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteFn           func(ctx context.Context, ID string) (bool, error)
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	return r.GetAllFn(ctx)
}

func (r *UserRepository) FindByID(ctx context.Context, ID string) (*domain.User, error) {
	return r.FindByIDFn(ctx, ID)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.FindByEmailFn(ctx, email)
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return r.CreateFn(ctx, user)
}

func (r *UserRepository) FindOrCreateUser(ctx context.Context, userData *domain.User) (*domain.User, bool, error) {
	return r.FindOrCreateUserFn(ctx, userData)
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	return r.UpdateFn(ctx, user)
}

func (r *UserRepository) Delete(ctx context.Context, ID string) (bool, error) {
	return r.DeleteFn(ctx, ID)
}
