package admin

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r adminReciever) FinOne(ctx context.Context, obj Admin) (Admin, error) {
	var admin Admin
	collection := r.db.Collection(r.collectionName)
	err := collection.FindOne(ctx, bson.D{{Key: "email", Value: obj.Email}}).Decode(&admin)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
