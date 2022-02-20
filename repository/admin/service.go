package admin

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository interface {
	Create(ctx context.Context, obj Admin) (Admin, error)
	FinOne(ctx context.Context, obj Admin) (Admin, error)
}

type adminReciever struct {
	db             mongo.Database
	collectionName string
}

func NewAdminRepository(db mongo.Database) AdminRepository {
	collectionName := "admin"
	return adminReciever{
		db:             db,
		collectionName: collectionName,
	}
}
