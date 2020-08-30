package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/ghanto/jivala-core/pkg/user/domain/user"
	"github.com/ghanto/jivala-core/pkg/util"
)

func TestUserRepo_Create(t *testing.T) {
	ctx := context.Background()

	db := util.SetupDb(ctx)
	defer db.Close()

	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	repo := NewUserRepository(db, &qb)
	util.EmptyTable(ctx, db, UserTable)

	recordsToCreate := []*user.User{
		{Email: "chuck@norris.com", Password: "secret_password"},
		{Email: "sean@doe.com", Password: "secret_password"},
		{Email: "jack@sparrow.com", Password: "secret_password"},
	}

	for _, tc := range recordsToCreate {
		_, err := repo.Create(ctx, tc)

		if err != nil {
			t.Errorf("unexpected error while creating users: %s", err)
		}
	}

	usr, err := repo.GetByEmail(ctx, "chuck@norris.com")
	if err != nil {
		t.Errorf("unexpected error while getting a user: %s", err)
	}

	if usr.Email != "chuck@norris.com" {
		t.Errorf("got another user than expected: %s", err)
	}

	dbUsers := []UserDB{}
	err = db.Select(&dbUsers, fmt.Sprintf("select id from %s", UserTable))
	if err != nil {
		t.Errorf("unexpected error while getting users: %s", err)
	}

	rowsAffected := len(dbUsers)

	if rowsAffected != 3 {
		t.Errorf(fmt.Sprintf("expected exactly 3 users, got %v", rowsAffected))
	}
}
