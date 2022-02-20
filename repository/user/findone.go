package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r userReciever) FinOne(ctx context.Context, obj User) (User, error) {
	var user User
	collection := r.db.Collection(r.collectionName)
	err := collection.FindOne(ctx, bson.D{{Key: "email", Value: obj.Email}}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
