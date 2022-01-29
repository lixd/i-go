package main

import (
	"context"
	"fmt"
	"log"

	"i-go/17x/ient/ent"
	"i-go/17x/ient/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	createUser, err := CreateUser(context.Background(), client)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("createUser:", createUser)
	queryUser, err := QueryUser(context.Background(), client)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("queryUser:", queryUser)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` 在 找不到用户 或 找到多于一个用户 时报错,
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
