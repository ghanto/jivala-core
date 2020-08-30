package user

import (
	"context"

	sqb "github.com/Masterminds/squirrel"
	"github.com/ghanto/jivala-core/pkg/user/domain/user"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(ctx context.Context, dto *user.User) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}

type userRepo struct {
	db *sqlx.DB
	qb *sqb.StatementBuilderType
}

func NewUserRepository(db *sqlx.DB, qb *sqb.StatementBuilderType) Repository {
	return &userRepo{
		db: db,
		qb: qb,
	}
}

func (r *userRepo) Create(ctx context.Context, dto *user.User) (*user.User, error) {
	q, args, err := r.qb.Insert(UserTable).
		Columns("password", "email").
		Values(dto.Password, dto.Email).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return &user.User{}, err
	}

	userDB := UserDB{}
	if err := r.db.QueryRowx(q, args...).StructScan(&userDB); err != nil {
		return &user.User{}, err
	}

	return userDB.ToModel(), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	sQuery := r.qb.Select("id", "email", "created_at").From(UserTable).Where(sqb.Eq{"email": email})
	sql, args, err := sQuery.ToSql()

	if err != nil {
		return nil, err
	}

	var usr UserDB
	err = r.db.QueryRowx(sql, args...).StructScan(&usr)

	if err != nil {
		return nil, err
	}

	return usr.ToModel(), nil
}
