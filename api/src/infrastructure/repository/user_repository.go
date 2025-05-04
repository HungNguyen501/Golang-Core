package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"golang-core/api/src/common/orm"
	"golang-core/api/src/infrastructure/database"
	"golang-core/api/src/infrastructure/repository/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	FindById(ctx context.Context, id string) (*entity.User, error)
	ListByIds(ctx context.Context, limit int, offset int) (orm.PaginationData[entity.User], error)
}

type DefaultUserRepository struct {
	db *database.Db
}

func NewUserRepository(db *database.Db) UserRepository {
	return &DefaultUserRepository{
		db: db,
	}
}

// Insert new user record
func (r *DefaultUserRepository) Insert(ctx context.Context, user entity.User) (res *entity.User, err error) {
	err = r.db.Primary().RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(&user).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return res, err
	}
	return &user, nil
}

// Update user record
func (r *DefaultUserRepository) Update(ctx context.Context, user *entity.User) (res *entity.User, err error) {
	err = r.db.Primary().RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return res, err
	}
	return user, nil
}

// Delete a user by id
func (r *DefaultUserRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	err := r.db.Primary().RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().Model((*entity.User)(nil)).Where("id = ?", userID).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// Find user by id
func (r *DefaultUserRepository) FindById(ctx context.Context, id string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Replica().NewSelect().Model(user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, SkipNotFound(err)
	}
	return user, nil
}

// List all users by paging
func (r *DefaultUserRepository) ListByIds(ctx context.Context, limit int, offset int) (res orm.PaginationData[entity.User], err error) {
	query := r.db.Replica().NewSelect().
		Model(&res.Data).
		Limit(limit).
		Offset(offset).
		OrderExpr("? ?", bun.Ident("id"), bun.Safe("ASC"))
	res.Total, err = query.ScanAndCount(ctx)
	res.CurrentOffset = offset
	return res, err
}
