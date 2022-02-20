package admin

import (
	"context"
)

func (r adminReciever) Create(ctx context.Context, obj Admin) (Admin, error) {
	collection := r.db.Collection(r.collectionName)
	_, err := collection.InsertOne(ctx, obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
