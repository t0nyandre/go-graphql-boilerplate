package mutation

import (
	"context"

	"github.com/t0nyandre/go-graphql-boilerplate/api/types"
	"github.com/t0nyandre/go-graphql-boilerplate/internal/user"
)

func (*Mutation) Register(ctx context.Context, args *struct {
	Username string
	Email    string
	Password string
}) (*types.UserType, error) {
	input := &user.User{
		Username: args.Username,
		Email:    args.Email,
		Password: args.Password,
	}

	user, err := ctx.Value("userService").(*user.Service).Create(input)
	if err != nil {
		return nil, err
	}

	return &types.UserType{User: user}, nil
}
