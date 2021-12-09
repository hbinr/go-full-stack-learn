package functionalprogram

import (
	"context"
	"log"
	"testing"
)

type User struct {
	id   int
	name string
}

//  TestUpdate 函数式编程-更新用户demo
func TestUpdate(t *testing.T) {
	UpdateHandle(context.Background())
}

func UpdateHandle(ctx context.Context) (err error) {
	return updateUser(ctx,
		func(ctx context.Context, user *User) (*User, error) {
			user.id = 2
			user.name = "test2"
			return user, nil
		},
	)
}

func updateUser(
	ctx context.Context,
	updateFn func(ctx context.Context, user *User) (*User, error),
) error {
	var user = User{1, "test"}
	res, err := updateFn(ctx, &user)

	if err != nil {
		return err
	}

	log.Printf("%+v", res)
	return nil
}
