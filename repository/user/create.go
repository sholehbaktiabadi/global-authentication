package user

import (
	"context"
)

func (r userReciever) Create(ctx context.Context, obj User) (User, error) {
	collection := r.db.Collection(r.collectionName)
	_, err := collection.InsertOne(ctx, obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
