package types

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/t0nyandre/go-graphql-boilerplate/internal/user"
)

type UserType struct {
	User *user.User
}

func (t *UserType) ID() graphql.ID {
	return graphql.ID(t.User.ID)
}

func (t *UserType) Username() string {
	return t.User.Username
}

func (t *UserType) Email() string {
	return t.User.Email
}

func (t *UserType) CreatedAt() graphql.Time {
	return graphql.Time{Time: t.User.CreatedAt}
}

func (t *UserType) UpdatedAt() *graphql.Time {
	if t.User.CreatedAt != t.User.UpdatedAt {
		return &graphql.Time{Time: t.User.UpdatedAt}
	}
	return nil
}
