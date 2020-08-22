package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	sqb "github.com/Masterminds/squirrel"
)

type Repository interface {
	Create(ctx context.Context, dto *CreateUserDb) (*UserDB, error)
	GetByEmail(ctx context.Context, email string) (*UserDB, error)
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

func (r *userRepo) Create(ctx context.Context, dto *CreateUserDb) (*UserDB, error) {
	q, args, err := r.qb.Insert(UserTable).
		Columns("login", "password", "email").
		Values(dto.Login, dto.Password, dto.Email).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return &UserDB{}, err
	}

	userDB := UserDB{}
	if err := r.db.QueryRowx(q, args...).StructScan(&userDB); err != nil {
		return &UserDB{}, err
	}

	return &userDB, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*UserDB, error) {
	sQuery := r.qb.Select("id", "email").From(UserTable).Where(sqb.Eq{"email": email})
	sql, args, err := sQuery.ToSql()

	if err != nil {
		return nil, err
	}

	var usr UserDB
	err = r.db.QueryRowx(sql, args...).StructScan(&usr)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}
