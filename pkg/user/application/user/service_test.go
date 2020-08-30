package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
)

func TestService_Register(t *testing.T) {
	testCases := map[string]struct {
		given       user.User
		givenErr    error
		expectedRet *user.User
		expectedErr error
	}{
		"all good": {
			given: user.User{
				Email:    "testuser1",
				Password: "fff",
			},
			expectedRet: &user.User{
				Email:    "testuser1",
				Password: "fff",
			},
		},
		"repository errors": {
			given: user.User{
				Email:    "testuser2",
				Password: "tbd",
			},
			givenErr:    errors.New("cannot connect to database"),
			expectedErr: errors.New("cannot connect to database"),
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		repo := NewMockRepo(tc.givenErr, &tc.given)
		s := NewService(repo)

		u, err := s.Register(ctx, tc.given)
		if tc.expectedErr != nil {
			assert.Error(t, err)
			assert.Equal(t, tc.expectedRet, u)
			assert.EqualError(t, err, tc.expectedErr.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedRet, u)
		}
	}
}

func TestService_GetByEmail(t *testing.T) {
	testCases := map[string]struct {
		given       *user.User
		givenErr    error
		expectedRet *user.User
		expectedErr error
	}{
		"all good": {
			given: &user.User{
				Email:    "testuser1@test.com",
				Password: "fff",
			},
			expectedRet: &user.User{
				Email:    "testuser1@test.com",
				Password: "fff",
			},
		},
		"repository errors": {
			given: &user.User{
				Email:    "testuser2",
				Password: "tbd",
			},
			givenErr:    errors.New("cannot connect to database"),
			expectedErr: errors.New("cannot connect to database"),
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		repo := NewMockRepo(tc.givenErr, tc.given)
		s := NewService(repo)

		u, err := s.GetByEmail(ctx, "test@zupa.com")
		if tc.expectedErr != nil {
			assert.Error(t, err)
			assert.Equal(t, tc.expectedRet, u)
			assert.EqualError(t, err, tc.expectedErr.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedRet, u)
		}
	}
}
