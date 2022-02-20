package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, obj User) (User, error)
	FinOne(ctx context.Context, obj User) (User, error)
}

type userReciever struct {
	db             mongo.Database
	collectionName string
}

func NewUserRepository(db mongo.Database) UserRepository {
	collectionName := "company"
	return userReciever{
		db:             db,
		collectionName: collectionName,
	}
}
